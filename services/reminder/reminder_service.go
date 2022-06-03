package reminder

import (
	"context"
	"fmt"
	"io/ioutil"
	"kost/entities"
	"net/http"
	"time"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type ServiceReminder struct {
	gapi *oauth2.Config
}

func NewReminderServices(config *oauth2.Config) *ServiceReminder {
	return &ServiceReminder{
		gapi: config,
	}
}

func (sr *ServiceReminder) GetLoginUrl(state string) string {
	// oauthState := gcal.GenerateStateOauthCookie(w)
	authUrl := sr.gapi.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return authUrl
}

func (sr *ServiceReminder) GetUserDataFromGoogle(w http.ResponseWriter, code string) ([]byte, error) {
	token, err := sr.gapi.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	cookieUrl := http.Cookie{Name: "oauthcode", Value: token.AccessToken, Expires: time.Now().Add(365 * 24 * time.Hour)}
	http.SetCookie(w, &cookieUrl)
	fmt.Println(token.AccessToken)

	const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}

func (sr *ServiceReminder) CreateEvent(code string, sreminder entities.Reminder) (*calendar.Event, error) {

	var reminder entities.Reminder
	copier.Copy(&reminder, &sreminder)
	token, err := sr.gapi.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	cookieUrl := http.Cookie{Name: "oauthcode", Value: token.AccessToken, Expires: time.Now().Add(365 * 24 * time.Hour)}
	echo.New().AcquireContext().SetCookie(&cookieUrl)
	fmt.Println(token.AccessToken)

	client := sr.gapi.Client(context.Background(), token)
	srv, err := calendar.NewService(context.TODO(), option.WithHTTPClient(client))
	if err != nil {
		log.Warn(err)
	}

	start_datatime := reminder.Year + "-" + reminder.Month + "-" + reminder.Day + "T" + reminder.Start + ":00+09:00"
	end_datatime := reminder.Year + "-" + reminder.Month + "-" + reminder.Day + "T" + reminder.End + ":00+09:00"

	event := &calendar.Event{
		Summary:     reminder.Title,
		Description: reminder.Description,
		Start: &calendar.EventDateTime{
			DateTime: start_datatime,
			TimeZone: "Asia/Jakarta",
		},
		End: &calendar.EventDateTime{
			DateTime: end_datatime,
			TimeZone: "Asia/Jakarta",
		},
		Attendees: []*calendar.EventAttendee{
			{Email: "gadipuranto@gmail.com"},
		},
	}

	remind, err := srv.Events.Insert("primary", event).Do()
	if err != nil {
		log.Warn(err)
	}

	return remind, nil
}

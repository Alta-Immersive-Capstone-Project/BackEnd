package reminder

import (
	"context"
	"fmt"
	"io/ioutil"
	"kost/entities"
	"net/http"

	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
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
	authUrl := sr.gapi.AuthCodeURL(state, oauth2.AccessTypeOnline)
	return authUrl
}

func (sr *ServiceReminder) GetUserDataFromGoogle(code string) ([]byte, error) {
	token, err := sr.gapi.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
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

func (sr *ServiceReminder) CreateEvent(reminder *entities.Reminder) *calendar.Event {

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

	return event
}

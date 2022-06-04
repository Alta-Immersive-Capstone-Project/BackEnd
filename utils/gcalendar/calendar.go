package gcalendar

import (
	// "fmt"

	"context"
	"fmt"
	"kost/configs"
	"kost/entities"
	"time"

	// "log"
	// "net/http"
	"github.com/labstack/gommon/log"
	"golang.org/x/oauth2"
	calendar "google.golang.org/api/calendar/v3"
)

type AuthConfig struct {
	authconfig *oauth2.Config
}

func NewAuthConfig(config *configs.AppConfig) *AuthConfig {
	var googleOauthConfig = &oauth2.Config{
		ClientID:     config.GCalendar.ClientID,
		ClientSecret: config.GCalendar.ClientSecret,
		RedirectURL:  config.GCalendar.RedirectUrl,
		Scopes:       []string{calendar.CalendarScope, calendar.CalendarEventsScope},
		Endpoint:     oauth2.Endpoint{AuthURL: "https://accounts.google.com/o/oauth2/auth", TokenURL: "https://oauth2.googleapis.com/token"},
	}

	return &AuthConfig{
		authconfig: googleOauthConfig,
	}
}

func (a *AuthConfig) Login(state string) string {
	authUrl := a.authconfig.AuthCodeURL(state, oauth2.AccessTypeOnline)
	return authUrl
}

func (a *AuthConfig) CreateReminder(code string, data entities.DataReminder) (calendar.Event, error) {

	token, err := a.authconfig.Exchange(context.Background(), code)
	if err != nil {
		log.Warn(err)
		return calendar.Event{}, err

	}
	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))

	calendarService, err := calendar.New(client)
	if err != nil {
		log.Warn(err)
		return calendar.Event{}, err
	}
	var now = time.Now()
	newEvent := &calendar.Event{
		Summary:     fmt.Sprintf("Reminder Pembayaran Sewa Kost %s", data.Title),
		Description: fmt.Sprintf("Harap Segera Membayar Tagihan dengan kode Booking ID %s, Total Tagihan %d, Silahkan Klik Link Diberikut untuk Melakukan Pembayaran %s", data.BookingID, data.Price, data.RedirectURL),
		Start:       &calendar.EventDateTime{DateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 2, 0, time.UTC).Format(time.RFC3339)},
		End:         &calendar.EventDateTime{DateTime: time.Date(now.Year(), now.Month(), now.Day()+1, now.Hour(), 0, 2, 0, time.UTC).Format(time.RFC3339)},
		Attendees: []*calendar.EventAttendee{
			{Email: data.Email},
		},
	}

	reminder, err := calendarService.Events.Insert("primary", newEvent).Do()
	if err != nil {
		log.Warn(err)
		return calendar.Event{}, err
	}
	return *reminder, nil
}

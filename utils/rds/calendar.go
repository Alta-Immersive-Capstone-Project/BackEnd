package utils

import (
	"kost/configs"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewAuthClient(config *configs.AppConfig) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.GCalendar.ClientID,
		ClientSecret: config.GCalendar.ClientSecret,
		RedirectURL:  config.GCalendar.RedirectUri,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/calendar"},
		Endpoint: google.Endpoint,
	}
}

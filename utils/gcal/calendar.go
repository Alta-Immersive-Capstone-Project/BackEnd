package gcal

import (
	// "fmt"

	"kost/configs"

	// "log"
	// "net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewAuthClient(config *configs.AppConfig) *oauth2.Config {
	var googleOauthConfig = &oauth2.Config{
		ClientID:     config.GCalendar.ClientID,
		ClientSecret: config.GCalendar.ClientSecret,
		RedirectURL:  config.GCalendar.RedirectUri,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/calendar"},
		Endpoint: google.Endpoint,
	}

	return googleOauthConfig
}

// func GenerateStateOauthCookie(w http.ResponseWriter) string {
// 	var expiration = time.Now().Add(365 * 24 * time.Hour)
// 	state := "token-state"
// 	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
// 	http.SetCookie(w, &cookie)

// 	return state
// }

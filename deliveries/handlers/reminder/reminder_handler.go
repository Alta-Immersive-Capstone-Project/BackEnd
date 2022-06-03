package reminder

import (
	"context"
	"fmt"
	"kost/deliveries/helpers"
	"kost/entities"
	"kost/services/reminder"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type HandlerReminder struct {
	auth reminder.ReminderService
}

var (
	config *oauth2.Config
)

func NewHandlersReminder(auth reminder.ReminderService) *HandlerReminder {
	return &HandlerReminder{
		auth: auth,
	}
}

func (h *HandlerReminder) OauthLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie := http.Cookie{Name: "oauthstate", Value: "token-state", Expires: time.Now().Add(365 * 24 * time.Hour)}
		c.SetCookie(&cookie)
		authUrl := h.auth.GetLoginUrl(cookie.Value)

		return c.Redirect(http.StatusTemporaryRedirect, authUrl)
	}
}

func (h *HandlerReminder) OauthCallback() echo.HandlerFunc {
	return func(c echo.Context) error {
		oauthState, _ := c.Request().Cookie("oauthstate")
		if c.QueryParam("state") != oauthState.Value {
			log.Warn("invalid oauth google state")
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}

		data, err := h.auth.GetUserDataFromGoogle(c.QueryParam("code"))
		if err != nil {
			log.Warn(err)
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}

		fmt.Fprintf(c.Response().Writer, "UserInfo: %s\n", data)
		return c.String(http.StatusOK, "success")
	}
}

func (h *HandlerReminder) CreateReminder() echo.HandlerFunc {
	return func(c echo.Context) error {
		oauthState, _ := c.Request().Cookie("oauthstate")
		if c.QueryParam("state") != oauthState.Value {
			log.Warn("invalid oauth google state")
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}

		ctx := context.Background()
		code := c.QueryParam("code")
		token, err := config.Exchange(ctx, code)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}

		oauthCode := new(http.Cookie)
		oauthCode.Name = "oauthtoken"
		oauthCode.Value = code
		c.SetCookie(oauthCode)

		client := config.Client(ctx, token)
		service, err := calendar.NewService(ctx, option.WithHTTPClient(client))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}

		request := entities.Reminder{
			Title:       "Coba Create Event",
			Description: "test desc create event",
			Year:        "2022",
			Month:       "06",
			Day:         "03",
			Start:       "11:00",
			End:         "17:00",
			Attendees:   "gadipuranto@gmail.com",
		}
		event := h.auth.CreateEvent(&request)
		remind, err := service.Events.Insert("primary", event).Do()
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}

		fmt.Printf("Event Created: %s\n", remind.HtmlLink)
		return c.JSON(http.StatusOK, helpers.StatusOK("OK", remind.HtmlLink))

	}
}

package reminder

import (
	"fmt"
	"kost/deliveries/helpers"
	"kost/entities"
	"kost/services/reminder"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type HandlerReminder struct {
	auth reminder.ReminderService
}

// var (
// 	config *oauth2.Config
// )

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
		oauthState, err := c.Request().Cookie("oauthstate")
		if c.QueryParam("state") != oauthState.Value {
			log.Warn("invalid oauth google state")
			return err
		}

		data, err := h.auth.GetUserDataFromGoogle(c.Response().Writer, c.QueryParam("code"))
		if err != nil {
			log.Warn(err)
			fmt.Println(data)
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

		code := c.QueryParams().Get("code")
		oauthCode := http.Cookie{Name: "oauthcode", Value: code, Expires: time.Now().Add(2 * time.Hour)}
		c.SetCookie(&oauthCode)

		var request entities.Reminder
		err := c.Bind(&request)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
		}

		remind, err := h.auth.CreateEvent(c.QueryParam("code"), request)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}

		fmt.Printf("Event Created: %s\n", remind.HtmlLink)
		return c.JSON(http.StatusOK, helpers.StatusOK("OK", remind.HtmlLink))

	}
}

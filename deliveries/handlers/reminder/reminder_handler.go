package reminder

import (
	"fmt"
	"kost/deliveries/helpers"
	"kost/deliveries/middlewares"
	"kost/entities"
	"kost/services/reminder"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type HandlerReminder struct {
	auth reminder.ReminderService
}

var codeOauth string

func NewHandlersReminder(auth reminder.ReminderService) *HandlerReminder {
	return &HandlerReminder{
		auth: auth,
	}
}

func (h *HandlerReminder) OauthLogin() echo.HandlerFunc {
	return func(c echo.Context) error {

		authUrl := h.auth.GetLoginUrl("oauthstate")
		return c.Redirect(http.StatusTemporaryRedirect, authUrl)
	}
}

func (h *HandlerReminder) OauthCallback() echo.HandlerFunc {
	return func(c echo.Context) error {
		code := c.FormValue("code")
		fmt.Println(code)
		if c.QueryParam("state") != "oauthstate" {
			log.Warn("invalid oauth google state")
			return c.JSON(http.StatusBadRequest, "Error Oauth")
		}
		codeOauth = code

		return c.JSON(http.StatusOK, "Oke")
	}
}

func (h *HandlerReminder) CreateReminder() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := middlewares.ExtractTokenRole(c)
		if role == "customer" {
			return c.JSON(http.StatusForbidden, helpers.ErrorAuthorize())
		}
		if codeOauth == "" {
			authUrl := h.auth.GetLoginUrl("oauthstate")
			return c.Redirect(http.StatusTemporaryRedirect, authUrl)
		}
		var request entities.AddReminderPay
		err := c.Bind(&request)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
		}
		data, err := h.auth.CreateEvent(codeOauth, request.BookingID)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, "Error Create Event")
		}
		return c.JSON(http.StatusOK, helpers.StatusOK("Create Event Berhasil", data))
	}
}

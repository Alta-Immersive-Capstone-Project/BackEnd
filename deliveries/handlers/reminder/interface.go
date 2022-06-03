package reminder

import "github.com/labstack/echo/v4"

type ReminderHandler interface {
	OauthLogin() echo.HandlerFunc
	OauthCallback() echo.HandlerFunc
	CreateReminder() echo.HandlerFunc
}

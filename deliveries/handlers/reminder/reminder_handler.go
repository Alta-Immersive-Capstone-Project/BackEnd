package reminder

import "github.com/labstack/echo/v4"

type ReminderHandler interface {
	GeListEvent() echo.HandlerFunc
}

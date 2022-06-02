package handlers

import "github.com/labstack/echo/v4"

type IForgotHandler interface {
	SendEmail() echo.HandlerFunc
	ResetPassword() echo.HandlerFunc
}

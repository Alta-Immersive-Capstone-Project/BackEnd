package handlers

import "github.com/labstack/echo/v4"

type HandleUser interface {
	CreateInternal(c echo.Context) error
	CreateCustomer(c echo.Context) error
	UpdateInternal(c echo.Context) error
	UpdateCustomer(c echo.Context) error
	DeleteInternal(c echo.Context) error
	DeleteCustomer(c echo.Context) error
}

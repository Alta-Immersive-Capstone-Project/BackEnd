package handlers

import "github.com/labstack/echo/v4"

type HandleUser interface {
	CreateInternal() echo.HandlerFunc
	CreateCustomer() echo.HandlerFunc
	UpdateInternal() echo.HandlerFunc
	UpdateCustomer() echo.HandlerFunc
	DeleteInternal() echo.HandlerFunc
	DeleteCustomer() echo.HandlerFunc
}

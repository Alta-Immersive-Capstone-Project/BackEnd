package handlers

import "github.com/labstack/echo/v4"

type CityHandler interface {
	CreateCity() echo.HandlerFunc
	GetAllCity() echo.HandlerFunc
	GetIDCity() echo.HandlerFunc
	UpdateCity() echo.HandlerFunc
	DeleteCity() echo.HandlerFunc
}

package handlers

import "github.com/labstack/echo/v4"

type ReviewHandler interface {
	InsertComment(c echo.Context) error
	GetByHouseID(c echo.Context) error
}

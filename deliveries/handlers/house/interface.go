package house

import "github.com/labstack/echo/v4"

type IHouseHandler interface {
	Store() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetAllByDist() echo.HandlerFunc
	Show() echo.HandlerFunc
}

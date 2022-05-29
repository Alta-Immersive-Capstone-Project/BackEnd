package district

import "github.com/labstack/echo/v4"

type IDistrictHandler interface {
	Store() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetAllByCity() echo.HandlerFunc
	Show() echo.HandlerFunc
}

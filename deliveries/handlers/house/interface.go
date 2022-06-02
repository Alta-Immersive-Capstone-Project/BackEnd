package house

import "github.com/labstack/echo/v4"

type IHouseHandler interface {
	Store() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetAllByDist() echo.HandlerFunc
	Show() echo.HandlerFunc
	Index() echo.HandlerFunc
	SelectHouseByDistrict() echo.HandlerFunc
	SelectHouseByCities() echo.HandlerFunc
	SelectHouseByCtyAndDst() echo.HandlerFunc
	SearchByTitle() echo.HandlerFunc
	// SearchBylocation() echo.HandlerFunc
}

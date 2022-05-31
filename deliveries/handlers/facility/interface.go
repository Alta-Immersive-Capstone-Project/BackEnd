package handlers

import "github.com/labstack/echo/v4"

type FacilityHandler interface {
	CreateFacility() echo.HandlerFunc
	GetAllFacility() echo.HandlerFunc
	GetFacilityID() echo.HandlerFunc
	UpdateFacility() echo.HandlerFunc
	DeleteFacility() echo.HandlerFunc
	GetNearFacility() echo.HandlerFunc
}

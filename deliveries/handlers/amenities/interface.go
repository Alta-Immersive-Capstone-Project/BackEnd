package handlers

import "github.com/labstack/echo/v4"

type AmenitiesHandler interface {
	CreateAmenities() echo.HandlerFunc
	GetAllAmenities() echo.HandlerFunc
	GetAmenitiesID() echo.HandlerFunc
	UpdateAmenities() echo.HandlerFunc
	DeleteAmenities() echo.HandlerFunc
}

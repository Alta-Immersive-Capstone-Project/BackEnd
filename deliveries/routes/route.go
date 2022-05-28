package routes

import (
	"kost/deliveries/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Path(e *echo.Echo, f *handlers.HandlersFacility, a *handlers.HandlersAmenities) {

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	facility := e.Group("/facility")
	facility.POST("", f.CreateFacility(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("K05T")}))
	facility.GET("", f.GetAllFacility())
	facility.GET("/:id", f.GetFacilityID())
	facility.PUT("/:id", f.UpdateFacility(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("K05T")}))
	facility.DELETE("/:id", f.DeleteFacility(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("K05T")}))

	amenities := e.Group("/amenities")
	amenities.POST("", a.CreateAmenities(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("K05T")}))
	amenities.GET("", a.GetAllAmenities())
	amenities.GET("/:id", a.GetAmenitiesID())
	amenities.PUT("/:id", a.UpdateAmenities(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("K05T")}))
	amenities.DELETE("/:id", a.DeleteAmenities(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("K05T")}))
}

package routes

import (
	"kost/deliveries/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Path(e *echo.Echo, f *handlers.HandlersFacility) {

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	facility := e.Group("/facility")
	facility.POST("", f.CreateFacility(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("K05T")}))
	facility.GET("", f.GetAllFacility())
	facility.GET("/:id", f.GetFacilityID())
	facility.PUT("/:id", f.UpdateFacility(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("K05T")}))
	facility.DELETE("/:id", f.DeleteFacility(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("K05T")}))
}

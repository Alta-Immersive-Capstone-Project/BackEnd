package routes

import (
<<<<<<< HEAD
	"backend/be8/deliveries/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterInternalRoute(e *echo.Echo, i *handlers.InternalHandler) {
	group := e.Group("/internal")
	group.POST("", i.CreateInternal) // Registration customer
	// group.PUT("", , middleware.JWTMiddleware())    // Edit customer profile
	// group.DELETE("", , middleware.JWTMiddleware()) // delete customer

=======
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
>>>>>>> c4ca72a1ed7c4b21751f758877248c248310c2f9
}

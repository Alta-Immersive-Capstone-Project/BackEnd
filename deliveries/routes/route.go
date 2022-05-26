package routes

import (
	"backend/be8/deliveries/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterInternalRoute(e *echo.Echo, i *handlers.InternalHandler) {
	group := e.Group("/internal")
	group.POST("", i.CreateInternal) // Registration customer
	// group.PUT("", , middleware.JWTMiddleware())    // Edit customer profile
	// group.DELETE("", , middleware.JWTMiddleware()) // delete customer

}

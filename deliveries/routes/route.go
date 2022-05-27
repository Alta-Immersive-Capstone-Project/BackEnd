package routes

import (
	"backend/be8/deliveries/handlers"

	"github.com/labstack/echo/v4"

	"backend/be8/deliveries/middlewares"
)

func UserRoute(e *echo.Echo, u *handlers.UserHandler) {
	internalGroup := e.Group("/internal")
	internalGroup.POST("/user", u.CreateInternal, middlewares.JWTMiddleware())
	internalGroup.DELETE("/user/:id", u.DeleteInternal, middlewares.JWTMiddleware())
	internalGroup.PUT("/user/:id", u.UpdateInternal, middlewares.JWTMiddleware())

	customerGroup := e.Group("/customer")

	customerGroup.POST("", u.CreateCustomer)
	customerGroup.PUT("/:id", u.UpdateCustomer, middlewares.JWTMiddleware())
	customerGroup.DELETE("/:id", u.DeleteCustomer, middlewares.JWTMiddleware())

}

func AuthRoute(e *echo.Echo, l *handlers.AuthHandler) {
	e.POST("/login", l.Login)

}

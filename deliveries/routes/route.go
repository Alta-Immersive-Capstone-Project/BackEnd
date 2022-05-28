package routes

import (
	"kost/deliveries/handlers"
	rh "kost/deliveries/handlers/reviews"
	th "kost/deliveries/handlers/transactions"
	"kost/deliveries/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func ReviewsPath(e *echo.Echo, rh rh.ReviewHandler) {
	e.POST("/reviews", rh.InsertComment)
	e.GET("/reviews/:room_id", rh.GetByRoomID)
}

func TransactionPath(e *echo.Echo, th th.TransactionHandler) {
	e.GET("/transactions", th.GetAllTransactionbyCustomer)
	e.GET("/admin/transactions", th.GetAllTransactionbyConsultant)
	e.POST("/transactions", th.InsertTransaction)
	e.PUT("/transactions/:booking_id", th.UpdateTransaction)
}

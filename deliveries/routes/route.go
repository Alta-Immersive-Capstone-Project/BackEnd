package routes

import (
	"kost/deliveries/handlers"
	dh "kost/deliveries/handlers/district"
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

func Path(e *echo.Echo, f *handlers.HandlersFacility, a *handlers.HandlersAmenities, d dh.IDistrictHandler) {
	e.Use(middleware.CORS())

	facility := e.Group("/facilities")
	facility.POST("", f.CreateFacility(), middlewares.JWTMiddleware())
	facility.GET("", f.GetAllFacility())
	facility.GET("/:id", f.GetFacilityID())
	facility.PUT("/:id", f.UpdateFacility(), middlewares.JWTMiddleware())
	facility.DELETE("/:id", f.DeleteFacility(), middlewares.JWTMiddleware())

	amenities := e.Group("/amenities")
	amenities.POST("", a.CreateAmenities(), middlewares.JWTMiddleware())
	amenities.GET("", a.GetAllAmenities())
	amenities.GET("/:id", a.GetAmenitiesID())
	amenities.PUT("/:id", a.UpdateAmenities(), middlewares.JWTMiddleware())
	amenities.DELETE("/:id", a.DeleteAmenities(), middlewares.JWTMiddleware())

	district := e.Group("/districts")
	district.POST("", d.Store(), middlewares.JWTMiddleware())
	district.GET("/:id", d.Show())
	e.GET("/cities/:id/districts", d.GetAllByCity())
	district.PUT("/:id", d.Update(), middlewares.JWTMiddleware())
	district.DELETE("/:id", d.Delete(), middlewares.JWTMiddleware())
}

func ReviewsPath(e *echo.Echo, rh rh.ReviewHandler) {
	// Customer
	e.POST("/reviews", rh.InsertComment, middlewares.JWTMiddleware())
	e.GET("/reviews/:room_id", rh.GetByRoomID)
}

func TransactionPath(e *echo.Echo, th th.TransactionHandler) {
	jwt := e.Group("", middlewares.JWTMiddleware())

	// Customer
	jwt.POST("/transactions", th.InsertTransaction)
	jwt.GET("/transactions", th.GetAllTransactionbyCustomer)

	// Admin
	jwt.GET("/admin/transactions", th.GetAllTransactionbyConsultant)
	jwt.PUT("/admin/transactions/:booking_id", th.UpdateTransaction)
}

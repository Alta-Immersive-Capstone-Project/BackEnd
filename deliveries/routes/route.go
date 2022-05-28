package routes

import (
	"kost/deliveries/handlers"
	rh "kost/deliveries/handlers/reviews"
	th "kost/deliveries/handlers/transactions"
	"kost/deliveries/middlewares"

	"github.com/labstack/echo/v4"
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
}
func RoomPath(e *echo.Echo, r *handlers.HandlersRoom) {
	facility := e.Group("/room")
	facility.POST("", r.CreateRoom(), middlewares.JWTMiddleware())
	facility.GET("", r.GetAllRoom())
	facility.GET("/:id", r.GetIDRoom())
	facility.PUT("/:id", r.UpdateRoom(), middlewares.JWTMiddleware())
	facility.DELETE("/:id", r.DeleteRoom(), middlewares.JWTMiddleware())
}
func CityPath(e *echo.Echo, C *handlers.HandlersCity) {
	facility := e.Group("/city")
	facility.POST("", C.CreateCity(), middlewares.JWTMiddleware())
	facility.GET("", C.GetAllCity())
	facility.GET("/:id", C.GetIDCity())
	facility.PUT("/:id", C.UpdateCity(), middlewares.JWTMiddleware())
	facility.DELETE("/:id", C.DeleteCity(), middlewares.JWTMiddleware())
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

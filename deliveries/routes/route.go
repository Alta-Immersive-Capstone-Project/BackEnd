package routes

import (
	"kost/deliveries/handlers"
	amenities "kost/deliveries/handlers/amenities"
	city "kost/deliveries/handlers/city"
	district "kost/deliveries/handlers/district"
	facility "kost/deliveries/handlers/facility"
	house "kost/deliveries/handlers/house"
	review "kost/deliveries/handlers/reviews"
	room "kost/deliveries/handlers/room"
	transaction "kost/deliveries/handlers/transactions"
	user "kost/deliveries/handlers/user"
	"kost/deliveries/middlewares"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo, u user.HandleUser) {
	internalGroup := e.Group("/admin")
	internalGroup.POST("", u.CreateInternal, middlewares.JWTMiddleware())
	internalGroup.DELETE("/:id", u.DeleteInternal, middlewares.JWTMiddleware())
	internalGroup.PUT("/:id", u.UpdateInternal, middlewares.JWTMiddleware())

	customerGroup := e.Group("/customer")

	customerGroup.POST("", u.CreateCustomer)
	customerGroup.PUT("/:id", u.UpdateCustomer, middlewares.JWTMiddleware())
	customerGroup.DELETE("/:id", u.DeleteCustomer, middlewares.JWTMiddleware())
}

func AuthRoute(e *echo.Echo, l *handlers.AuthHandler) {
	e.POST("/login", l.Login)
}

func Path(e *echo.Echo, f *facility.HandlersFacility, a *amenities.HandlersAmenities, d district.IDistrictHandler, h house.IHouseHandler) {

	facility := e.Group("/facilities")
	facility.POST("", f.CreateFacility(), middlewares.JWTMiddleware())
	e.GET("/houses/:id/facilities", f.GetAllFacility())
	facility.GET("/:id", f.GetFacilityID())
	facility.PUT("/:id", f.UpdateFacility(), middlewares.JWTMiddleware())
	facility.DELETE("/:id", f.DeleteFacility(), middlewares.JWTMiddleware())

	amenities := e.Group("/room/:id/amenities")
	amenities.POST("", a.CreateAmenities(), middlewares.JWTMiddleware())
	amenities.GET("", a.GetAmenitiesID())
	amenities.PUT("", a.UpdateAmenities(), middlewares.JWTMiddleware())
	amenities.DELETE("", a.DeleteAmenities(), middlewares.JWTMiddleware())

	district := e.Group("/districts")
	district.POST("", d.Store(), middlewares.JWTMiddleware())
	district.GET("/:id", d.Show())
	e.GET("/cities/:id/districts", d.GetAllByCity())
	district.PUT("/:id", d.Update(), middlewares.JWTMiddleware())
	district.DELETE("/:id", d.Delete(), middlewares.JWTMiddleware())

	house := e.Group("/houses")
	house.POST("", h.Store(), middlewares.JWTMiddleware())
	house.GET("/:id", h.Show())
	e.GET("/districts/:id/houses", h.GetAllByDist())
	house.PUT("/:id", h.Update(), middlewares.JWTMiddleware())
	house.DELETE("/:id", h.Delete(), middlewares.JWTMiddleware())
}
func RoomPath(e *echo.Echo, r *room.HandlersRoom) {
	facility := e.Group("/room")
	facility.POST("", r.CreateRoom(), middlewares.JWTMiddleware())
	facility.GET("", r.GetAllRoom())
	facility.GET("/:id", r.GetIDRoom())
	facility.PUT("/:id", r.UpdateRoom(), middlewares.JWTMiddleware())
	facility.DELETE("/:id", r.DeleteRoom(), middlewares.JWTMiddleware())
}
func CityPath(e *echo.Echo, C *city.HandlersCity) {
	facility := e.Group("/cities")
	facility.POST("", C.CreateCity(), middlewares.JWTMiddleware())
	facility.GET("", C.GetAllCity())
	facility.GET("/:id", C.GetIDCity())
	facility.PUT("/:id", C.UpdateCity(), middlewares.JWTMiddleware())
	facility.DELETE("/:id", C.DeleteCity(), middlewares.JWTMiddleware())
}

func ReviewsPath(e *echo.Echo, review review.ReviewHandler) {
	// Customer
	e.POST("/reviews", review.InsertComment, middlewares.JWTMiddleware())
	e.GET("/reviews/:room_id", review.GetByRoomID)
}

func TransactionPath(e *echo.Echo, transaction transaction.TransactionHandler) {
	jwt := e.Group("", middlewares.JWTMiddleware())

	// Customer
	jwt.POST("/transactions", transaction.InsertTransaction)
	jwt.GET("/transactions", transaction.GetAllTransactionbyCustomer)

	// Admin
	jwt.GET("/admin/transactions", transaction.GetAllTransactionbyConsultant)
	jwt.PUT("/admin/transactions/:booking_id", transaction.UpdateTransaction)
}

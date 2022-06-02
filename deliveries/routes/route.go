package routes

import (
	"kost/deliveries/handlers"
	amenities "kost/deliveries/handlers/amenities"
	city "kost/deliveries/handlers/city"
	district "kost/deliveries/handlers/district"
	facility "kost/deliveries/handlers/facility"
	forgot "kost/deliveries/handlers/forgot"
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
	internalGroup.POST("", u.CreateInternal(), middlewares.JWTMiddleware())
	internalGroup.DELETE("/:id", u.DeleteInternal(), middlewares.JWTMiddleware())
	internalGroup.PUT("/:id", u.UpdateInternal(), middlewares.JWTMiddleware())
	internalGroup.GET("", u.GetAllMember(), middlewares.JWTMiddleware())

	customerGroup := e.Group("/customer")

	customerGroup.POST("", u.CreateCustomer())
	customerGroup.GET("/:id", u.GetCustomerByID())
	customerGroup.PUT("/:id", u.UpdateCustomer(), middlewares.JWTMiddleware())
	customerGroup.DELETE("/:id", u.DeleteCustomer(), middlewares.JWTMiddleware())
}

func AuthRoute(e *echo.Echo, l *handlers.AuthHandler, f *forgot.ForgotHandler) {
	e.POST("/login", l.Login())
	e.GET("/forgot", f.SendEmail())
	e.POST("/forgot", f.ResetPassword(), middlewares.JWTMiddleware())
}

func Path(e *echo.Echo, f *facility.HandlersFacility, a *amenities.HandlersAmenities, d district.IDistrictHandler, h house.IHouseHandler) {

	facility := e.Group("/facilities")
	facility.POST("", f.CreateFacility(), middlewares.JWTMiddleware())
	e.GET("/houses/:id/facilities", f.GetNearFacility())
	e.GET("/districts/:id/facilities", f.GetAllFacility())

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
	district.GET("", d.Index())
	district.GET("/:id", d.Show())
	e.GET("/cities/:id/districts", d.GetAllByCity())
	district.PUT("/:id", d.Update(), middlewares.JWTMiddleware())
	district.DELETE("/:id", d.Delete(), middlewares.JWTMiddleware())

	house := e.Group("/houses")
	house.POST("", h.Store(), middlewares.JWTMiddleware())
	house.GET("", h.Index())
	house.GET("/search", h.SearchByTitle())
	// house.GET("/search", h.SearchBylocation())
	house.GET("/:id", h.Show())
	house.PUT("/:id", h.Update(), middlewares.JWTMiddleware())
	house.DELETE("/:id", h.Delete(), middlewares.JWTMiddleware())
	house.GET("/district/:id", h.SelectHouseByDistrict())
	e.GET("/districts/:id/houses", h.GetAllByDist())
	e.GET("/cities/:cid/districts/houses", h.SelectHouseByCities())
	e.GET("/cities/:cid/districts/:dist_id/houses", h.SelectHouseByCtyAndDst())
}
func RoomPath(e *echo.Echo, r *room.HandlersRoom) {
	room := e.Group("/room")
	room.POST("", r.CreateRoom(), middlewares.JWTMiddleware())
	e.GET("/rooms/:id", r.GetAllRoom())
	room.GET("/:id", r.GetIDRoom())
	room.PUT("/:id", r.UpdateRoom(), middlewares.JWTMiddleware())
	room.DELETE("/:id", r.DeleteRoom(), middlewares.JWTMiddleware())
}
func CityPath(e *echo.Echo, C *city.HandlersCity) {
	city := e.Group("/cities")
	city.POST("", C.CreateCity(), middlewares.JWTMiddleware())
	city.GET("", C.GetAllCity())
	city.GET("/:id", C.GetIDCity())
	city.PUT("/:id", C.UpdateCity(), middlewares.JWTMiddleware())
	city.DELETE("/:id", C.DeleteCity(), middlewares.JWTMiddleware())
}

func ReviewsPath(e *echo.Echo, review review.ReviewHandler) {
	// Customer
	e.POST("/reviews", review.InsertComment, middlewares.JWTMiddleware())
	e.GET("/reviews/:room_id", review.GetByRoomID)
}

func TransactionPath(e *echo.Echo, transaction transaction.TransactionHandler) {
	// Customer
	e.POST("/transactions", transaction.InsertTransaction, middlewares.JWTMiddleware())
	e.GET("/transactions", transaction.GetAllTransactionbyCustomer, middlewares.JWTMiddleware())
	e.POST("/transactions/callback", transaction.UpdateStatus)

	// Admin
	admin := e.Group("/admin/transactions", middlewares.JWTMiddleware())
	admin.GET("", transaction.GetAllTransactionbyConsultant)
	admin.PUT("/:booking_id", transaction.UpdateTransaction)
	admin.GET("/kost", transaction.GetAllTransactionbyKost)
}

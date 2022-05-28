package main

import (
	"kost/configs"
	"kost/deliveries/handlers"
	"kost/deliveries/routes"
	"kost/repositories/facility"
	cFacility "kost/services/facility"

	"github.com/go-playground/validator"

	userRepository "kost/repositories/user"
	authService "kost/services/auth"

	storageProvider "kost/services/storage"

	userService "kost/services/user"

	"kost/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	// Get Config
	config := configs.Get()
	// Init DB
	DB := utils.NewMysqlGorm(config)
	// Init Facility Service
	facilityRepo := facility.NewFacilityDB(DB)
	facilityService := cFacility.NewServiceFacility(facilityRepo)
	facilityHandler := handlers.NewHandlersFacility(facilityService, validator.New())
	// Initiate Echo
	e := echo.New()
	// Connect To Route
	routes.Path(e, facilityHandler)

	e.Logger.Fatal(e.Start(":8000"))
	// utils.Migrate(DB)

	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	// 	AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	// }))

	userRepository := userRepository.NewUserRepository(DB)

	authService := authService.NewAuthService(userRepository)
	userService := userService.NewUserService(userRepository, validator.New())

	s3 := storageProvider.NewS3()

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService, s3)

	routes.AuthRoute(e, authHandler)
	routes.UserRoute(e, userHandler)

	e.Logger.Fatal(e.Start(":" + config.App.Port))
}

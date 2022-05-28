package main

import (
	"kost/configs"
	"kost/deliveries/handlers"
	"kost/deliveries/routes"
	"kost/repositories/amenities"
	"kost/repositories/facility"
	cAmenities "kost/services/amenities"
	cFacility "kost/services/facility"

	"github.com/go-playground/validator/v10"

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

	// Init Amenities Service
	amenitiesRepo := amenities.NewAmenitiesDB(DB)
	amenitiesService := cAmenities.NewServiceAmenities(amenitiesRepo)
	amenitiesHandler := handlers.NewHandlersAmenities(amenitiesService, validator.New())
	// Initiate Echo
	e := echo.New()
	// Connect To Route
	routes.Path(e, facilityHandler, amenitiesHandler)
	// utils.Migrate(DB)

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

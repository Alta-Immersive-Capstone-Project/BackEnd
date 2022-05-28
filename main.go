package main

import (
	"kost/configs"
	"kost/deliveries/handlers"
	"kost/deliveries/middlewares"
	"kost/deliveries/routes"
	"kost/deliveries/validations"
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

	districtRepo "kost/repositories/district"

	districtServices "kost/services/district"

	districtHandlers "kost/deliveries/handlers/district"

	reviewRepo "kost/repositories/reviews"
	transactionRepo "kost/repositories/transactions"

	reviewService "kost/services/reviews"
	transactionService "kost/services/transactions"

	reviewHandlers "kost/deliveries/handlers/reviews"
	transactionHandlers "kost/deliveries/handlers/transactions"
)

func main() {
	// Get Config
	config := configs.Get()

	// Init DB
	DB := utils.NewMysqlGorm(config)

	// Migrate
	utils.Migrate(DB)

	// Initiate Echo
	e := echo.New()

	// Repositories
	userRepository := userRepository.NewUserRepository(DB)
	facilityRepo := facility.NewFacilityDB(DB)
	amenitiesRepo := amenities.NewAmenitiesDB(DB)
	reviewsRepo := reviewRepo.NewReviewModel(DB)
	transactionsRepo := transactionRepo.NewTransactionModel(DB)
	districtRepo := districtRepo.NewDistrictRepo(DB)

	// Validation
	validation := validations.NewValidation(validator.New())

	// Services
	s3 := storageProvider.NewS3()
	userService := userService.NewUserService(userRepository, validator.New())
	authService := authService.NewAuthService(userRepository)
	facilityService := cFacility.NewServiceFacility(facilityRepo)
	amenitiesService := cAmenities.NewServiceAmenities(amenitiesRepo)
	reviewsService := reviewService.NewReviewService(reviewsRepo)
	transactionsService := transactionService.NewTransactionService(transactionsRepo)
	districtService := districtServices.NewDistService(districtRepo)
	// Handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService, s3)
	facilityHandler := handlers.NewHandlersFacility(facilityService, validator.New())
	amenitiesHandler := handlers.NewHandlersAmenities(amenitiesService, validator.New())
	reviewsHandler := reviewHandlers.NewReviewHandler(reviewsService, validation)
	transactionsHandler := transactionHandlers.NewTransactionHandler(transactionsService, validation)
	districtHandler := districtHandlers.NewDistrictHandler(districtService, validation)
	// Middlewares
	middlewares.General(e)

	// Routes
	routes.AuthRoute(e, authHandler)
	routes.UserRoute(e, userHandler)
	routes.Path(e, facilityHandler, amenitiesHandler, districtHandler)
	routes.ReviewsPath(e, reviewsHandler)
	routes.TransactionPath(e, transactionsHandler)

	e.Logger.Fatal(e.Start(":" + config.App.Port))
}

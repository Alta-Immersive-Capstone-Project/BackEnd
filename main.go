package main

import (
	"kost/configs"
	"kost/deliveries/handlers"
	"kost/deliveries/middlewares"
	"kost/deliveries/routes"
	"kost/deliveries/validations"
	"kost/repositories/amenities"
	"kost/repositories/city"
	"kost/repositories/facility"
	"kost/repositories/image"
	"kost/repositories/room"
	cAmenities "kost/services/amenities"
	cFacility "kost/services/facility"

	"github.com/go-playground/validator/v10"

	userRepository "kost/repositories/user"
	authService "kost/services/auth"

	storageProvider "kost/services/storage"

	citesService "kost/services/city"
	roomsService "kost/services/room"
	userService "kost/services/user"

	utils "kost/utils/rds"

	"github.com/labstack/echo/v4"

	districtRepo "kost/repositories/district"
	houseRepo "kost/repositories/house"

	districtServices "kost/services/district"
	houseServices "kost/services/house"

	districtHandlers "kost/deliveries/handlers/district"
	houseHandlers "kost/deliveries/handlers/house"

	reviewRepo "kost/repositories/reviews"
	transactionRepo "kost/repositories/transactions"

	reviewService "kost/services/reviews"
	transactionService "kost/services/transactions"

	reviewHandlers "kost/deliveries/handlers/reviews"
	transactionHandlers "kost/deliveries/handlers/transactions"

	amenitiesHandlers "kost/deliveries/handlers/amenities"
	cityHandlers "kost/deliveries/handlers/city"
	facilityHandlers "kost/deliveries/handlers/facility"
	roomHandlers "kost/deliveries/handlers/room"
	userHandlers "kost/deliveries/handlers/user"
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
	cityRepo := city.NewCityDB(DB)
	roomRepo := room.NewRoomDB(DB)
	imageRepo := image.NewImageDB(DB)
	districtRepo := districtRepo.NewDistrictRepo(DB)
	houseRepo := houseRepo.NewHouseRepo(DB)

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
	cityService := citesService.NewServiceCity(cityRepo)
	roomService := roomsService.NewServiceRoom(roomRepo, imageRepo)
	districtService := districtServices.NewDistService(districtRepo)
	houseService := houseServices.NewHouseService(houseRepo)

	// Handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := userHandlers.NewUserHandler(userService, s3, validation)
	facilityHandler := facilityHandlers.NewHandlersFacility(facilityService, validation)
	amenitiesHandler := amenitiesHandlers.NewHandlersAmenities(amenitiesService, validation)
	reviewsHandler := reviewHandlers.NewReviewHandler(reviewsService, validation)
	transactionsHandler := transactionHandlers.NewTransactionHandler(transactionsService, validation)
	cityHandler := cityHandlers.NewHandlersCity(cityService, validator.New())
	roomHandler := roomHandlers.NewHandlersRoom(roomService, validator.New())
	districtHandler := districtHandlers.NewDistrictHandler(districtService, validation)
	houseHandler := houseHandlers.NewHouseHandler(houseService, validation)

	// Middlewares
	middlewares.General(e)

	// Routes
	routes.AuthRoute(e, authHandler)
	routes.UserRoute(e, userHandler)
	routes.Path(e, facilityHandler, amenitiesHandler, districtHandler, houseHandler)
	routes.ReviewsPath(e, reviewsHandler)
	routes.TransactionPath(e, transactionsHandler)
	routes.CityPath(e, cityHandler)
	routes.RoomPath(e, roomHandler)

	// e.Logger.Fatal(e.Start(":" + config.App.Port))
	e.Logger.Fatal(e.Start(":8000"))
}

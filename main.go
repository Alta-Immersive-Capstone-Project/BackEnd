package main

import (
	"backend/be8/configs"
	"backend/be8/deliveries/handlers"
	"backend/be8/deliveries/routes"

	userRepository "backend/be8/repositories/user"
	authService "backend/be8/services/auth"

	storageProvider "backend/be8/services/storage"

	userService "backend/be8/services/user"

	"backend/be8/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	config := configs.Get()
	db := utils.NewMysqlGorm(config)
	// utils.Migrate(db)

	e := echo.New()
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	// 	AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	// }))

	userRepository := userRepository.NewUserRepository(db)

	authService := authService.NewAuthService(userRepository)
	userService := userService.NewUserService(userRepository)

	s3 := storageProvider.NewS3()

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService, s3)

	routes.AuthRoute(e, authHandler)
	routes.UserRoute(e, userHandler)

	e.Logger.Fatal(e.Start(":" + config.App.Port))
}

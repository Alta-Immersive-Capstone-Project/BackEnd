package main

import (
	"kost/configs"
	"kost/deliveries/handlers"
	"kost/deliveries/routes"
	"kost/repositories/facility"
	cFacility "kost/services/facility"
	"kost/utils/rds"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	// Get Config
	config := configs.InitConfig()
	// Init DB
	DB := rds.InitDB(config)
	// Init Facility Service
	facilityRepo := facility.NewFacilityDB(DB)
	facilityService := cFacility.NewServiceFacility(facilityRepo)
	facilityHandler := handlers.NewHandlersFacility(facilityService, validator.New())
	// Initiate Echo
	e := echo.New()
	// Connect To Route
	routes.Path(e, facilityHandler)

	e.Logger.Fatal(e.Start(":8000"))
}

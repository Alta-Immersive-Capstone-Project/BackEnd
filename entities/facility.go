package entities

import "gorm.io/gorm"

type Facility struct {
	gorm.Model
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	HouseID   uint    `json:"house_id"`
}

type AddNewFacility struct {
	Name      string  `json:"name" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	HouseID   uint    `json:"house_id" validate:"required"`
}

type UpdateFacility struct {
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type RespondFacility struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	HouseID   uint    `json:"house_id"`
}

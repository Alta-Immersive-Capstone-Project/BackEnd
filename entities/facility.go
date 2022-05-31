package entities

import "gorm.io/gorm"

type Facility struct {
	gorm.Model
	Name       string  `json:"name"`
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
	DistrictID uint    `json:"district_id"`
}

type AddNewFacility struct {
	Name       string  `json:"name" validate:"required"`
	Longitude  float64 `json:"longitude" validate:"required"`
	Latitude   float64 `json:"latitude" validate:"required"`
	DistrictID uint    `json:"district_id" validate:"required"`
}

type UpdateFacility struct {
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type RespondFacility struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
	DistrictID uint    `json:"district_id"`
}

type NearFacility struct {
	Name   string  `json:"name"`
	Radius float64 `json:"radius"`
}

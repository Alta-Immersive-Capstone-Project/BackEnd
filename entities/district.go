package entities

import "gorm.io/gorm"

type District struct {
	gorm.Model
	Name      string `gorm:"type:varchar(50);unique"`
	Longitude float64
	Latitude  float64
	CityID    uint
	House     []House    `gorm:"foreingkey:DistrictID"`
	Facility  []Facility `gorm:"foreingkey:DistrictID"`
}

type AddDistrict struct {
	Name      string  `json:"name" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	CityID    uint    `json:"city_id" validate:"required"`
}

type UpdateDistrict struct {
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	CityID    uint    `json:"city_id"`
}

type RespondDistrict struct {
	ID        uint    `json:"dist_id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	CityID    uint    `json:"city_id"`
}

package entities

import "gorm.io/gorm"

type City struct {
	gorm.Model
	City      string     `json:"city_name"`
	Districts []District `gorm:"foreingkey:CityID"`
}
type CityResponse struct {
	ID   uint   `json:"id"`
	City string `json:"city_name"`
}
type AddCity struct {
	City string `json:"city_name" validate:"required"`
}

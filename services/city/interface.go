package city

import "kost/entities"

type CityRepo interface {
	CreateCity(Insert entities.AddCity) (entities.CityResponse, error)
	GetAllCity() ([]entities.CityResponse, error)
	GetIDCity(id uint) (entities.RespondRoom, error)
	UpdateCity(id uint, update entities.City) (entities.RespondRoom, error)
	DeleteCity(id uint) error
}

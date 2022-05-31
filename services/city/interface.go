package city

import "kost/entities"

type CityRepo interface {
	CreateCity(Insert entities.AddCity) (entities.CityResponse, error)
	GetAllCity() ([]entities.CityResponse, error)
	GetIDCity(id uint) (entities.CityResponse, error)
	UpdateCity(id uint, update entities.AddCity) (entities.CityResponse, error)
	DeleteCity(id uint) error
}

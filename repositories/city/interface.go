package city

import "kost/entities"

type CityRepo interface {
	CreateCity(new entities.City) (entities.City, error)
	GetAllCity() ([]entities.City, error)
	GetAllCityDistricts(cityID uint) ([]entities.City, error)
	GetCity(cityID uint) (entities.City, error)
	DeleteCity(cityID uint) error
	UpdateCity(id uint, new entities.AddCity) (entities.City, error)
}

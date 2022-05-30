package house

import "kost/entities"

type IHouseService interface {
	CreateHouse(Insert entities.AddHouse) (entities.HouseResponse, error)
	UpdateHouse(id uint, update entities.UpdateHouse) (entities.HouseResponse, error)
	DeleteHouse(id uint) error
	GetAllHouseByDist(dist_id uint) ([]entities.HouseResponse, error)
	GetHouseID(id uint) (entities.HouseResponse, error)
}

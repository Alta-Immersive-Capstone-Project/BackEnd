package house

import "kost/entities"

type IRepoHouse interface {
	CreateHouse(addNew entities.House) (entities.House, error)
	UpdateHouse(id uint, update entities.House) (entities.House, error)
	DeleteHouse(id uint) error
	GetHouseID(id uint) (entities.House, error)
	GetAllHouse(dist_id uint) ([]entities.House, error)
}

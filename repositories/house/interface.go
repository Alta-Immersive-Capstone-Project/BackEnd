package house

import "kost/entities"

type IRepoHouse interface {
	CreateHouse(addNew entities.House) (entities.House, error)
	UpdateHouse(id uint, update entities.House) (entities.House, error)
	DeleteHouse(id uint) error
	GetHouseID(id uint) (entities.House, error)
	GetAllHouseByDistrict(dist_id uint) ([]entities.House, error)
	GetAllHouseByCities(cid uint) ([]entities.House, error)
	GetAllHouseByDstAndCty(cid uint, dist_id uint) ([]entities.House, error)
	SelectAllHouse() ([]entities.House, error)
	FindHouseByTitle(name string) ([]entities.House, error)
	FindHouseByLocation(lat float64, long float64) ([]entities.House, error)
}

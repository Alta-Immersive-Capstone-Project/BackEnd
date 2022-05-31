package house

import "kost/entities"

type IRepoHouse interface {
	CreateHouse(addNew entities.House) (entities.House, error)
	UpdateHouse(id uint, update entities.House) (entities.House, error)
	DeleteHouse(id uint) error
	GetHouseID(id uint) (entities.House, error)
	GetAllHouseByDistrict(dist_id uint) ([]entities.HouseResponseJoin, error)
	GetAllHouseByCities(cid uint) ([]entities.HouseResponseJoin, error)
	GetAllHouseByDstAndCty(cid uint, dist_id uint) ([]entities.HouseResponseJoin, error)
	SelectAllHouse() ([]entities.HouseResponseJoin, error)
	FindHouseByTitle(name string) ([]entities.HouseResponseJoin, error)
	FindHouseByLocation(lat float64, long float64) ([]entities.HouseResponseJoin, error)
}

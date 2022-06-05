package house

import "kost/entities"

type IRepoHouse interface {
	CreateHouse(addNew entities.House) (entities.House, error)
	UpdateHouse(id uint, update entities.House) (entities.House, error)
	DeleteHouse(id uint) error
	GetAllHouseByDist(dist_id uint) ([]entities.HouseResponseGet, error)
	GetHouseID(id uint) (entities.HouseResponseGetByID, error)
	GetAllHouseByDistrict(dist_id uint) ([]entities.HouseResponseGet, error)
	GetAllHouseByCities(cid uint) ([]entities.HouseResponseGet, error)
	GetAllHouseByDstAndCty(cid uint, dist_id uint) ([]entities.HouseResponseGet, error)
	SelectAllHouse() ([]entities.HouseResponseGet, error)
	FindHouseByTitle(name string) ([]entities.HouseResponseGet, error)
	// FindHouseByLocation(lat float64, long float64) ([]entities.HouseResponseJoin, error)
}

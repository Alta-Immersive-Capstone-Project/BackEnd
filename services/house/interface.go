package house

import "kost/entities"

type IHouseService interface {
	CreateHouse(Insert entities.AddHouse) (entities.HouseResponse, error)
	UpdateHouse(id uint, update entities.UpdateHouse) (entities.HouseResponse, error)
	DeleteHouse(id uint) error
	GetHouseID(id uint) (entities.HouseResponse, error)
	GetAllHouseByDistrict(dist_id uint) ([]entities.HouseResponse, error)
	FindAllHouseByDistrict(dist_id uint) ([]entities.HouseResponseJoin, error)
	FindAllHouseByCities(cid uint) ([]entities.HouseResponseJoin, error)
	FindAllHouseByCtyAndDst(cid uint, dist_id uint) ([]entities.HouseResponseJoin, error)
	SelectAllHouse() ([]entities.HouseResponseJoin, error)
	FindHouseByTitle(title string) ([]entities.HouseResponseJoin, error)
	// FindHouseByLocation(lat float64, long float64) ([]entities.HouseResponseJoin, error)
}

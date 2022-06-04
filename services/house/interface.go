package house

import "kost/entities"

type IHouseService interface {
	CreateHouse(Insert entities.AddHouse, url string) (entities.HouseResponse, error)
	UpdateHouse(id uint, update entities.House) (entities.HouseResponse, error)
	DeleteHouse(id uint) error
	GetHouseID(id uint) (entities.HouseResponse, error)
	GetAllHouseByDistrict(dist_id uint) ([]entities.HouseResponseGetAll, error)
	FindAllHouseByDistrict(dist_id uint) ([]entities.HouseResponseGetAll, error)
	FindAllHouseByCities(cid uint) ([]entities.HouseResponseGetAll, error)
	FindAllHouseByCtyAndDst(cid uint, dist_id uint) ([]entities.HouseResponseGetAll, error)
	SelectAllHouse() ([]entities.HouseResponseGetAll, error)
	FindHouseByTitle(title string) ([]entities.HouseResponseGetAll, error)
	// FindHouseByLocation(lat float64, long float64) ([]entities.HouseResponseJoin, error)
}

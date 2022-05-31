package district

import "kost/entities"

type DistrictControl interface {
	CreateDist(Insert entities.AddDistrict) (entities.RespondDistrict, error)
	UpdateDist(id uint, update entities.UpdateDistrict) (entities.RespondDistrict, error)
	DeleteDist(id uint) error
	GetAllDist(cid uint) ([]entities.RespondDistrict, error)
	GetDistID(id uint) (entities.RespondDistrict, error)
	SelectAllDistrict() (entities.RespondDistrict, error)
}

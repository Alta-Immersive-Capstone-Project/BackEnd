package district

import "kost/entities"

type RepoDistrict interface {
	CreateDistrict(district entities.District) (entities.District, error)
	UpdateDistrict(id uint, update entities.District) (entities.District, error)
	DeleteDistrict(id uint) error
	GetDistrictID(id uint) (entities.District, error)
	GetAllDistrict(cid uint) ([]entities.District, error)
}

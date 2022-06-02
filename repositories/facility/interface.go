package facility

import "kost/entities"

type RepoFacility interface {
	CreateFacility(New entities.Facility) (entities.Facility, error)
	GetAllFacility(HouseID uint) ([]entities.Facility, error)
	GetFacilityID(id uint) (entities.Facility, error)
	UpdateFacility(id uint, UpdateFacility entities.Facility) (entities.Facility, error)
	DeleteFacility(id uint) error
	GetNearFacility(HouseID uint) ([]entities.NearFacility, error)
}

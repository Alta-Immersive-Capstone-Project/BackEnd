package facility

import "kost/entities"

type FacilityControl interface {
	CreateFacility(Insert entities.AddNewFacility) (entities.RespondFacility, error)
	GetAllFacility(HouseID uint) ([]entities.RespondFacility, error)
	GetFacilityID(id uint) (entities.RespondFacility, error)
	UpdateFacility(id uint, update entities.UpdateFacility) (entities.RespondFacility, error)
	DeleteFacility(id uint) error
	GetNearFacility(HouseID uint) ([]entities.NearFacility, error)
}

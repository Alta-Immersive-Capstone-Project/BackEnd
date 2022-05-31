package facility

import (
	"kost/entities"
	"kost/repositories/facility"

	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
)

type ServiceFacility struct {
	repo facility.RepoFacility
}

func NewServiceFacility(Repo facility.RepoFacility) *ServiceFacility {
	return &ServiceFacility{
		repo: Repo,
	}
}

// Create Facility
func (s *ServiceFacility) CreateFacility(Insert entities.AddNewFacility) (entities.RespondFacility, error) {

	var Add entities.Facility
	copier.Copy(&Add, &Insert)
	res, err := s.repo.CreateFacility(Add)
	if err != nil {
		log.Warn(err)
		return entities.RespondFacility{}, err
	}

	var result entities.RespondFacility

	copier.Copy(&result, &res)

	return result, nil
}

// Get All Facilities
func (s *ServiceFacility) GetAllFacility(DistrictID uint) ([]entities.RespondFacility, error) {

	res, err := s.repo.GetAllFacility(DistrictID)
	if err != nil {
		log.Warn(err)
		return []entities.RespondFacility{}, err
	}
	var result []entities.RespondFacility

	copier.Copy(&result, &res)
	return result, nil
}

// Get Facility By ID
func (s *ServiceFacility) GetFacilityID(id uint) (entities.RespondFacility, error) {

	res, err := s.repo.GetFacilityID(id)
	if err != nil {
		log.Warn(err)
		return entities.RespondFacility{}, err
	}

	var result entities.RespondFacility

	copier.Copy(&result, &res)
	return result, nil
}

// Update Facility By ID
func (s *ServiceFacility) UpdateFacility(id uint, update entities.UpdateFacility) (entities.RespondFacility, error) {

	var UpdateFacility entities.Facility
	copier.Copy(&UpdateFacility, &update)
	res, err := s.repo.UpdateFacility(id, UpdateFacility)
	if err != nil {
		log.Warn(err)
		return entities.RespondFacility{}, err
	}

	var result entities.RespondFacility

	copier.Copy(&result, &res)
	return result, nil
}

// Delete Facility By ID
func (s *ServiceFacility) DeleteFacility(id uint) error {

	err := s.repo.DeleteFacility(id)
	if err != nil {
		log.Warn(err)
		return err
	}
	return nil
}

func (s *ServiceFacility) GetNearFacility(HouseID uint) ([]entities.NearFacility, error) {
	result, err := s.repo.GetNearFacility(HouseID)
	if err != nil {
		log.Warn(err)
		return []entities.NearFacility{}, err
	}
	return result, nil
}

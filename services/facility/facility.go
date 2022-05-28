package facility

import (
	"kost/entities"
	"kost/repositories/facility"

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

	NewAdd := entities.Facility{
		Name:      Insert.Name,
		Longitude: Insert.Longitude,
		Latitude:  Insert.Latitude,
		HouseID:   Insert.HouseID,
	}

	res, err := s.repo.CreateFacility(NewAdd)
	if err != nil {
		log.Warn(err)
		return entities.RespondFacility{}, err
	}

	result := entities.RespondFacility{
		ID:        res.HouseID,
		Name:      res.Name,
		Longitude: res.Longitude,
		Latitude:  res.Latitude,
		HouseID:   res.HouseID,
	}

	return result, nil
}

// Get All Facilities
func (s *ServiceFacility) GetAllFacility(HouseID uint) ([]entities.RespondFacility, error) {

	res, err := s.repo.GetAllFacility(HouseID)
	if err != nil {
		log.Warn(err)
		return []entities.RespondFacility{}, err
	}
	var result []entities.RespondFacility
	for _, v := range res {
		facility := entities.RespondFacility{
			ID:        v.HouseID,
			Name:      v.Name,
			Longitude: v.Longitude,
			Latitude:  v.Latitude,
			HouseID:   v.HouseID,
		}
		result = append(result, facility)
	}
	return result, nil
}

// Get Facility By ID
func (s *ServiceFacility) GetFacilityID(id uint) (entities.RespondFacility, error) {

	res, err := s.repo.GetFacilityID(id)
	if err != nil {
		log.Warn(err)
		return entities.RespondFacility{}, err
	}

	result := entities.RespondFacility{
		ID:        res.HouseID,
		Name:      res.Name,
		Longitude: res.Longitude,
		Latitude:  res.Latitude,
		HouseID:   res.HouseID,
	}
	return result, nil
}

// Update Facility By ID
func (s *ServiceFacility) UpdateFacility(id uint, update entities.UpdateFacility) (entities.RespondFacility, error) {

	UpdateFacility := entities.Facility{
		Name:      update.Name,
		Longitude: update.Longitude,
		Latitude:  update.Latitude,
	}

	res, err := s.repo.UpdateFacility(id, UpdateFacility)
	if err != nil {
		log.Warn(err)
		return entities.RespondFacility{}, err
	}

	result := entities.RespondFacility{
		ID:        res.HouseID,
		Name:      res.Name,
		Longitude: res.Longitude,
		Latitude:  res.Latitude,
		HouseID:   res.HouseID,
	}
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

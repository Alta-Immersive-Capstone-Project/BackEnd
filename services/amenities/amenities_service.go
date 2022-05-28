package amenities

import (
	"kost/entities"
	"kost/repositories/amenities"

	"github.com/jinzhu/copier"

	"github.com/labstack/gommon/log"
)

type ServiceAmenities struct {
	repo amenities.RepoAmenities
}

func NewServiceAmenities(Repo amenities.RepoAmenities) *ServiceAmenities {
	return &ServiceAmenities{
		repo: Repo,
	}
}

// Create Amenities
func (s *ServiceAmenities) CreateAmenities(Insert entities.AddAmenities) (entities.RespondAmenities, error) {

	var New entities.Amenities
	copier.Copy(&New, &Insert)

	res, err := s.repo.CreateAmenities(New)
	if err != nil {
		log.Warn(err)
		return entities.RespondAmenities{}, err
	}

	var result entities.RespondAmenities
	copier.Copy(&result, &res)
	return result, nil
}

// Get All Facilities
func (s *ServiceAmenities) GetAllAmenities(HouseID uint) ([]entities.RespondAmenities, error) {

	res, err := s.repo.GetAllAmenities(HouseID)
	if err != nil {
		log.Warn(err)
		return []entities.RespondAmenities{}, err
	}
	var result []entities.RespondAmenities
	copier.Copy(&result, &res)
	return result, nil
}

// Get Amenities By ID
func (s *ServiceAmenities) GetAmenitiesID(id uint) (entities.RespondAmenities, error) {

	res, err := s.repo.GetAmenitiesID(id)
	if err != nil {
		log.Warn(err)
		return entities.RespondAmenities{}, err
	}
	var result entities.RespondAmenities

	copier.Copy(&result, &res)

	return result, nil
}

// Update Amenities By ID
func (s *ServiceAmenities) UpdateAmenities(id uint, update entities.UpdateAmenities) (entities.RespondAmenities, error) {

	var UpdateAmenities entities.Amenities
	copier.Copy(&UpdateAmenities, &update)

	res, err := s.repo.UpdateAmenities(id, UpdateAmenities)
	if err != nil {
		log.Warn(err)
		return entities.RespondAmenities{}, err
	}
	var result entities.RespondAmenities

	copier.Copy(&result, &res)

	return result, nil
}

// Delete Amenities By ID
func (s *ServiceAmenities) DeleteAmenities(id uint) error {

	err := s.repo.DeleteAmenities(id)
	if err != nil {
		log.Warn(err)
		return err
	}
	return nil
}

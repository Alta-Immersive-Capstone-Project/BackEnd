package city

import (
	"kost/entities"
	"kost/repositories/city"

	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
)

type ServiceCity struct {
	repo city.CityRepo
}

func NewServiceCity(Repo city.CityRepo) *ServiceCity {
	return &ServiceCity{
		repo: Repo,
	}
}

func (r *ServiceCity) CreateCity(Insert entities.AddCity) (entities.CityResponse, error) {

	newCity := entities.City{}
	copier.Copy(&newCity, &Insert)

	result, err := r.repo.CreateCity(newCity)
	if err != nil {
		log.Warn(err)
		return entities.CityResponse{}, err
	}
	var res entities.CityResponse
	copier.Copy(&res, &result)

	return res, nil
}

func (s *ServiceCity) GetAllCity() ([]entities.CityResponse, error) {

	res, err := s.repo.GetAllCity()
	if err != nil {
		log.Warn(err)
		return []entities.CityResponse{}, err
	}

	result := []entities.CityResponse{}
	copier.Copy(&result, &res)

	return result, nil
}
func (s *ServiceCity) GetIDCity(id uint) (entities.RespondRoom, error) {

	res, err := s.repo.GetCity(id)
	if err != nil {
		log.Warn(err)
		return entities.RespondRoom{}, err
	}

	result := entities.RespondRoom{}
	copier.Copy(&result, &res)

	return result, nil
}

func (s *ServiceCity) UpdateCity(id uint, update entities.City) (entities.RespondRoom, error) {

	res, err := s.repo.UpdateCity(id, update)
	if err != nil {
		log.Warn(err)
		return entities.RespondRoom{}, err
	}

	result := entities.RespondRoom{}
	copier.Copy(&result, &res)

	return result, nil
}

func (s *ServiceCity) DeleteCity(id uint) error {

	err := s.repo.DeleteCity(id)
	if err != nil {
		log.Warn(err)
		return err
	}

	return nil
}

package house

import (
	"kost/entities"
	"kost/repositories/house"

	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
)

type HouseService struct {
	repo house.IRepoHouse
}

func NewHouseService(Repo house.IRepoHouse) *HouseService {
	return &HouseService{
		repo: Repo,
	}
}

func (hs *HouseService) CreateHouse(Insert entities.HouseRequest) (entities.HouseResponse, error) {
	var house entities.House
	copier.Copy(&house, &Insert)
	res, err := hs.repo.CreateHouse(house)
	if err != nil {
		log.Warn(err)
		return entities.HouseResponse{}, err
	}
	var result entities.HouseResponse
	copier.Copy(&result, &res)
	return result, nil
}

func (hs *HouseService) UpdateHouse(id uint, update entities.HouseRequest) (entities.HouseResponse, error) {
	var UpdateHouse entities.House
	copier.Copy(&UpdateHouse, &update)
	res, err := hs.repo.UpdateHouse(id, UpdateHouse)
	if err != nil {
		log.Warn(err)
		return entities.HouseResponse{}, err
	}
	var result entities.HouseResponse

	copier.Copy(&result, &res)
	return result, nil
}

func (hs *HouseService) DeleteHouse(id uint) error {
	err := hs.repo.DeleteHouse(id)
	if err != nil {
		log.Warn(err)
		return err
	}
	return nil
}

func (hs *HouseService) GetAllHouseByDist(dist_id uint) ([]entities.HouseResponse, error) {
	res, err := hs.repo.GetAllHouse(dist_id)
	if err != nil {
		log.Warn(err)
		return []entities.HouseResponse{}, err
	}
	var result []entities.HouseResponse

	copier.Copy(&result, &res)
	return result, nil
}
func (hs *HouseService) GetHouseID(id uint) (entities.HouseResponse, error) {
	res, err := hs.repo.GetHouseID(id)
	if err != nil {
		log.Warn(err)
		return entities.HouseResponse{}, err
	}

	var result entities.HouseResponse

	copier.Copy(&result, &res)
	return result, nil
}

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

func (hs *HouseService) CreateHouse(Insert entities.AddHouse) (entities.HouseResponse, error) {
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

func (hs *HouseService) UpdateHouse(id uint, update entities.UpdateHouse) (entities.HouseResponse, error) {
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

func (hs *HouseService) GetAllHouseByDistrict(dist_id uint) ([]entities.HouseResponse, error) {
	res, err := hs.repo.GetAllHouseByDist(dist_id)
	if err != nil {
		log.Warn(err)
		return []entities.HouseResponse{}, err
	}
	var result []entities.HouseResponse

	copier.Copy(&result, &res)
	return result, nil
}

func (hs *HouseService) FindAllHouseByDistrict(dist_id uint) ([]entities.HouseResponseJoin, error) {
	res, err := hs.repo.GetAllHouseByDistrict(dist_id)
	if err != nil {
		log.Warn(err)
		return []entities.HouseResponseJoin{}, err
	}
	var result []entities.HouseResponseJoin
	copier.Copy(&result, &res)
	return result, nil
}

func (hs *HouseService) FindAllHouseByCities(cid uint) ([]entities.HouseResponseJoin, error) {
	res, err := hs.repo.GetAllHouseByCities(cid)
	if err != nil {
		log.Warn(err)
		return []entities.HouseResponseJoin{}, err
	}
	var result []entities.HouseResponseJoin
	copier.Copy(&result, &res)
	return result, nil
}
func (hs *HouseService) FindAllHouseByCtyAndDst(cid uint, dist_id uint) ([]entities.HouseResponseJoin, error) {
	res, err := hs.repo.GetAllHouseByDstAndCty(cid, dist_id)
	if err != nil {
		log.Warn(err)
		return []entities.HouseResponseJoin{}, err
	}
	var result []entities.HouseResponseJoin
	copier.Copy(&result, &res)
	return result, nil
}
func (hs *HouseService) SelectAllHouse() ([]entities.HouseResponseJoin, error) {
	res, err := hs.repo.SelectAllHouse()
	if err != nil {
		log.Warn(err)
		return []entities.HouseResponseJoin{}, err
	}
	var result []entities.HouseResponseJoin
	copier.Copy(&result, &res)
	return result, nil
}
func (hs *HouseService) FindHouseByTitle(title string) ([]entities.HouseResponseJoin, error) {
	res, err := hs.repo.FindHouseByTitle(title)
	if err != nil {
		log.Warn(err)
		return []entities.HouseResponseJoin{}, err
	}
	var result []entities.HouseResponseJoin
	copier.Copy(&result, &res)
	return result, nil
}

// func (hs *HouseService) FindHouseByLocation(lat float64, long float64) ([]entities.HouseResponseJoin, error) {
// 	res, err := hs.repo.FindHouseByLocation(lat, long)
// 	if err != nil {
// 		log.Warn(err)
// 		return []entities.HouseResponseJoin{}, err
// 	}
// 	var result []entities.HouseResponseJoin
// 	copier.Copy(&result, &res)
// 	return result, nil
// }

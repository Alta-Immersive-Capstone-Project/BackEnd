package district

import (
	"kost/entities"
	"kost/repositories/district"

	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
)

type DistrictService struct {
	repo district.RepoDistrict
}

func NewDistService(Repo district.RepoDistrict) *DistrictService {
	return &DistrictService{
		repo: Repo,
	}
}

func (ds *DistrictService) CreateDist(Insert entities.AddDistrict) (entities.RespondDistrict, error) {
	var District entities.District
	copier.Copy(&District, &Insert)
	res, err := ds.repo.CreateDistrict(District)
	if err != nil {
		log.Warn(err)
		return entities.RespondDistrict{}, err
	}
	var result entities.RespondDistrict
	copier.Copy(&result, &res)
	return result, nil
}

func (ds *DistrictService) UpdateDist(id uint, update entities.UpdateDistrict) (entities.RespondDistrict, error) {
	var UpdateDistrict entities.District
	copier.Copy(&UpdateDistrict, &update)
	res, err := ds.repo.UpdateDistrict(id, UpdateDistrict)
	if err != nil {
		log.Warn(err)
		return entities.RespondDistrict{}, err
	}
	var result entities.RespondDistrict

	copier.Copy(&result, &res)
	return result, nil
}
func (ds *DistrictService) DeleteDist(id uint) error {
	err := ds.repo.DeleteDistrict(id)
	if err != nil {
		log.Warn(err)
		return err
	}
	return nil
}
func (ds *DistrictService) GetAllDist(cid uint) ([]entities.RespondDistrict, error) {
	res, err := ds.repo.GetAllDistrict(cid)
	if err != nil {
		log.Warn(err)
		return []entities.RespondDistrict{}, err
	}
	var result []entities.RespondDistrict

	copier.Copy(&result, &res)
	return result, nil
}
func (ds *DistrictService) GetDistID(id uint) (entities.RespondDistrict, error) {
	res, err := ds.repo.GetDistrictID(id)
	if err != nil {
		log.Warn(err)
		return entities.RespondDistrict{}, err
	}

	var result entities.RespondDistrict

	copier.Copy(&result, &res)
	return result, nil
}
func (ds *DistrictService) SelectAllDistrict() ([]entities.RespondDistrict, error) {
	res, err := ds.repo.SelectAllDistrict()
	if err != nil {
		log.Warn(err)
		return []entities.RespondDistrict{}, err
	}

	var result []entities.RespondDistrict
	copier.Copy(&result, &res)
	return result, nil
}

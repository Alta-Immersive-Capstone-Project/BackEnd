package district

import (
	"kost/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type DistrictRepo struct {
	Db *gorm.DB
}

func NewDisttrictRepo(db *gorm.DB) *DistrictRepo {
	return &DistrictRepo{
		Db: db,
	}
}

func (dr *DistrictRepo) CreateDistrict(district entities.District) (entities.District, error) {
	err := dr.Db.Create(&district).Error
	if err != nil {
		log.Warn(err)
		return district, err
	}
	return district, nil
}

func (dr *DistrictRepo) UpdateDistrict(id uint, update entities.District) (entities.District, error) {
	var District entities.District
	err := dr.Db.Where("id = ?", id).First(&District).Updates(&update).Error
	if err != nil {
		log.Warn("Update Error", err)
		return District, err
	}
	return District, nil
}
func (dr *DistrictRepo) DeleteDistrict(id uint) error {
	var District entities.District
	err := dr.Db.Where("id = ?", id).First(&District).Delete(&District).Error
	if err != nil {
		log.Warn("Delete Error", err)
		return err
	}
	return nil
}
func (dr *DistrictRepo) GetDistrictID(id uint) (entities.District, error) {
	var District entities.District
	err := dr.Db.Where("id = ?", id).First(&District).Error
	if err != nil {
		log.Warn("Error Get By ID", err)
		return District, err
	}
	return District, nil
}
func (dr *DistrictRepo) GetAllDistrict(cid uint) ([]entities.District, error) {
	var Districts []entities.District
	err := dr.Db.Where("city_id = ?", cid).Find(&Districts).Error
	if err != nil {
		log.Warn("Error Get Data", err)
		return Districts, err
	}
	return Districts, nil
}

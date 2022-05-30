package house

import (
	"kost/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type HouseRepo struct {
	Db *gorm.DB
}

func NewHouseRepo(db *gorm.DB) *HouseRepo {
	return &HouseRepo{
		Db: db,
	}
}

func (hr *HouseRepo) CreateHouse(addNew entities.House) (entities.House, error) {
	err := hr.Db.Create(&addNew).Error
	if err != nil {
		log.Warn(err)
		return entities.House{}, err
	}
	return addNew, nil
}

func (hr *HouseRepo) UpdateHouse(id uint, update entities.House) (entities.House, error) {
	var house entities.House
	err := hr.Db.Where("id = ?", id).First(&house).Updates(&house).Error
	if err != nil {
		log.Warn("Update Error", err)
		return house, err
	}
	return house, nil
}

func (hr *HouseRepo) DeleteHouse(id uint) error {
	var house entities.House
	err := hr.Db.Where("id = ?", id).First(&house).Delete(&house).Error
	if err != nil {
		log.Warn("Delete Error", err)
		return err
	}
	return nil
}
func (hr *HouseRepo) GetHouseID(id uint) (entities.House, error) {
	var house entities.House
	err := hr.Db.Where("id = ?", id).First(&house).Error
	if err != nil {
		log.Warn("Error Get By ID", err)
		return house, err
	}
	return house, nil
}
func (hr *HouseRepo) GetAllHouse(dist_id uint) ([]entities.House, error) {
	var house []entities.House
	err := hr.Db.Where("district_id = ?", dist_id).Find(&house).Error
	if err != nil {
		log.Warn("Error Get Data", err)
		return house, err
	}
	return house, nil
}

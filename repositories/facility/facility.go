package facility

import (
	"kost/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type FacilityDB struct {
	Db *gorm.DB
}

// Depen
func NewFacilityDB(db *gorm.DB) *FacilityDB {
	return &FacilityDB{
		Db: db,
	}
}

// Create Facility
func (f *FacilityDB) CreateFacility(New entities.Facility) (entities.Facility, error) {
	if err := f.Db.Create(&New).Error; err != nil {
		log.Warn(err)
		return New, err
	}
	return New, nil
}

// Get All Facility
func (f *FacilityDB) GetAllFacility(HouseID uint) ([]entities.Facility, error) {
	var facilities []entities.Facility
	err := f.Db.Where("house_id = ?", HouseID).Find(&facilities).Error
	if err != nil {
		log.Warn("Error Get Data", err)
		return facilities, err
	}
	return facilities, nil
}

// Get Facility By ID
func (f *FacilityDB) GetFacilityID(id uint) (entities.Facility, error) {
	var facility entities.Facility
	err := f.Db.Where("id= ?", id).First(&facility).Error
	if err != nil {
		log.Warn("Error Get By ID", err)
		return facility, err
	}
	return facility, nil
}

// Update Facility By ID
func (f *FacilityDB) UpdateFacility(id uint, UpdateFacility entities.Facility) (entities.Facility, error) {
	var facility entities.Facility

	err := f.Db.Where("id =?", id).First(&facility).Updates(&UpdateFacility).Find(&facility).Error
	if err != nil {
		log.Warn("Update Error", err)
		return facility, err
	}

	return facility, nil
}

// Delete Facility By ID
func (f *FacilityDB) DeleteFacility(id uint) error {
	var delete entities.Facility

	err := f.Db.Where("id = ?", id).First(&delete).Delete(&delete).Error
	if err != nil {
		log.Warn("Delete Error", err)
		return err
	}
	return nil
}

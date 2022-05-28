package amenities

import (
	"kost/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type AmenitiesDB struct {
	Db *gorm.DB
}

// Depen
func NewAmenitiesDB(db *gorm.DB) *AmenitiesDB {
	return &AmenitiesDB{
		Db: db,
	}
}

// Create Amenities
func (f *AmenitiesDB) CreateAmenities(New entities.Amenities) (entities.Amenities, error) {
	if err := f.Db.Create(&New).Error; err != nil {
		log.Warn(err)
		return New, err
	}
	return New, nil
}

// Get All Amenities
func (f *AmenitiesDB) GetAllAmenities(RoomID uint) ([]entities.Amenities, error) {
	var facilities []entities.Amenities
	err := f.Db.Where("room_id = ?", RoomID).Find(&facilities).Error
	if err != nil {
		log.Warn("Error Get Data", err)
		return facilities, err
	}
	return facilities, nil
}

// Get Amenities By ID
func (f *AmenitiesDB) GetAmenitiesID(id uint) (entities.Amenities, error) {
	var facility entities.Amenities
	err := f.Db.Where("id= ?", id).First(&facility).Error
	if err != nil {
		log.Warn("Error Get By ID", err)
		return facility, err
	}
	return facility, nil
}

// Update Amenities By ID
func (f *AmenitiesDB) UpdateAmenities(id uint, UpdateAmenities entities.Amenities) (entities.Amenities, error) {
	var facility entities.Amenities

	err := f.Db.Where("id =?", id).First(&facility).Updates(&UpdateAmenities).Find(&facility).Error
	if err != nil {
		log.Warn("Update Error", err)
		return facility, err
	}

	return facility, nil
}

// Delete Amenities By ID
func (f *AmenitiesDB) DeleteAmenities(id uint) error {
	var delete entities.Amenities

	err := f.Db.Where("id = ?", id).First(&delete).Delete(&delete).Error
	if err != nil {
		log.Warn("Delete Error", err)
		return err
	}
	return nil
}

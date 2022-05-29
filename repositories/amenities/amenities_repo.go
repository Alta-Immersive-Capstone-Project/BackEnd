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

// Get Amenities By ID
func (f *AmenitiesDB) GetAmenitiesID(RoomID uint) (entities.Amenities, error) {
	var facility entities.Amenities
	err := f.Db.Where("room_id= ?", RoomID).First(&facility).Error
	if err != nil {
		log.Warn("Error Get By ID", err)
		return facility, err
	}
	return facility, nil
}

// Update Amenities By ID
func (f *AmenitiesDB) UpdateAmenities(RoomID uint, UpdateAmenities entities.Amenities) (entities.Amenities, error) {
	var facility entities.Amenities

	err := f.Db.Where("RoomID =?", RoomID).First(&facility).Updates(&UpdateAmenities).Find(&facility).Error
	if err != nil {
		log.Warn("Update Error", err)
		return facility, err
	}

	return facility, nil
}

// Delete Amenities By ID
func (f *AmenitiesDB) DeleteAmenities(RoomID uint) error {
	var delete entities.Amenities

	err := f.Db.Where("RoomID = ?", RoomID).First(&delete).Delete(&delete).Error
	if err != nil {
		log.Warn("Delete Error", err)
		return err
	}
	return nil
}

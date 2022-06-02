package utils

import (
	"kost/entities"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&entities.User{}, &entities.Image{}, &entities.City{}, &entities.District{}, &entities.House{}, &entities.Amenities{}, &entities.Room{}, &entities.Facility{}, &entities.Transaction{}, &entities.Review{},
	)
}

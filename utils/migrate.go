package utils

import (
	"kost/entities"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&entities.User{}, &entities.Room{}, &entities.Facility{}, &entities.Amenities{}, &entities.Transaction{}, &entities.Review{}, &entities.Image{},
	)
}

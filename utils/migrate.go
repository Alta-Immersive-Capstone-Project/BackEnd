package utils

import (
	"kost/entities"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
<<<<<<< HEAD
		&entities.User{}, &entities.Room{}, &entities.Facility{}, &entities.Amenities{}, &entities.Transaction{}, &entities.Review{}, &entities.Image{},
=======
		&entities.User{}, &entities.District{}, &entities.House{}, &entities.Amenities{}, &entities.Room{}, &entities.Facility{}, &entities.Transaction{}, &entities.Review{},
>>>>>>> 26df28703b223ed9633b240d0bc5db7286c471d2
	)
}

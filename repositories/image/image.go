package image

import (
	"kost/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type imageDB struct {
	Db *gorm.DB
}

func NewImageDB(db *gorm.DB) *imageDB {
	return &imageDB{
		Db: db,
	}
}

func (i *imageDB) CreateImage(new entities.Image) error {
	err := i.Db.Create(&new).Error
	if err != nil {
		log.Warn(err)
		return err
	}
	return nil
}

func (i *imageDB) DeleteImage(roomID, userID uint) error {
	var image entities.Image
	err := i.Db.Where("roomID = ? OR userID =?", roomID, userID).Delete(&image).Error
	if err != nil {
		log.Warn(err)
		return err
	}
	return nil
}

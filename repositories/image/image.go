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

func (i *imageDB) GetAllImage(roomID uint) ([]entities.Image, error) {
	var image []entities.Image
	err := i.Db.Where("room_id=?", roomID).Find(&image).Error
	if err != nil {
		log.Warn(err)
		return []entities.Image{}, err
	}
	return image, nil
}
func (i *imageDB) GetImage(ID uint) (entities.Image, error) {
	var image entities.Image
	err := i.Db.Where("id=?", ID).First(&image).Error
	if err != nil {
		log.Warn(err)
		return entities.Image{}, err
	}
	return image, nil
}
func (i *imageDB) DeleteImage(imageID uint) error {
	var image entities.Image
	err := i.Db.Where("id = ?", imageID).Delete(&image).Error
	if err != nil {
		log.Warn(err)
		return err
	}
	return nil
}

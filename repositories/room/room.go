package room

import (
	"kost/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type roomDB struct {
	Db *gorm.DB
}

func NewRoomDB(db *gorm.DB) *roomDB {
	return &roomDB{
		Db: db,
	}
}

func (r *roomDB) CreateRoom(new entities.Room) (entities.Room, error) {
	err := r.Db.Create(&new).Error
	if err != nil {
		log.Warn(err)
		return entities.Room{}, err
	}
	return new, nil
}

func (r *roomDB) GetAllRoom() ([]entities.Room, error) {
	var rooms []entities.Room
	err := r.Db.Find(&rooms).Error
	if err != nil {
		log.Warn(err)
		return []entities.Room{}, err
	}
	return rooms, nil
}

func (r *roomDB) GetRoomID(id uint) (entities.Room, error) {
	var room entities.Room
	err := r.Db.Where("id = ?", id).Preload("Images").Preload("Amenities").Preload("Reviews").First(&room).Error
	if err != nil {
		log.Warn(err)
		return entities.Room{}, err
	}
	return room, nil
}

func (r *roomDB) UpdateRoom(id uint, new entities.UpdateRoom) (entities.Room, error) {
	var room entities.Room
	err := r.Db.Where("id = ?", id).First(&room).Updates(&new).Find(&room).Error
	if err != nil {
		log.Warn(err)
		return entities.Room{}, err
	}
	return room, nil
}

func (r *roomDB) DeleteRoom(id uint) error {
	var room entities.Room

	err := r.Db.Where("id = ?", id).First(&room).Delete(&room).Error

	if err != nil {
		log.Warn(err)
		return err
	}
	return nil
}
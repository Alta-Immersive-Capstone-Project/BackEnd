package entities

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	RoomID uint   `json:"room_id" validate:"required"`
	Url    string `json:"url"`
}

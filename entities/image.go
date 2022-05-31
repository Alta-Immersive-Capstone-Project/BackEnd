package entities

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	RoomID uint   `json:"roomID" validate:"required"`
	Url    string `json:"url"`
}

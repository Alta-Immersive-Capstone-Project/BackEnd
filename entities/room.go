package entities

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	HouseID                uint          `json:"house_id"`
	Type                   string        `json:"type"`
	UserID                 uint          `json:"userid"`
	Price                  int32         `json:"price"`
	Additional_description string        `json:"additional_description"`
	Transactions           []Transaction `gorm:"foreingkey:RoomID"`
	Reviews                []Review      `gorm:"foreingkey:RoomID"`
	Images                 []Image       `gorm:"foreingkey:RoomID"`
	Amenities              Amenities     `gorm:"foreingkey:RoomID"`
}

type AddRoom struct {
	HouseID                uint   `json:"house_id" validate:"required" form:"house_id"`
	Type                   string `gorm:"type:varchar(100);not null" json:"type" validate:"required" form:"type"`
	Price                  int32  `gorm:"type:int;not null" json:"price" validate:"required" form:"price"`
	Additional_description string `gorm:"type:varchar(100);not null" json:"additional_description" validate:"required" form:"additional_description"`
}
type UpdateRoom struct {
	Type                   string `gorm:"type:varchar(100);not null" json:"type" form:"type"`
	Price                  int32  `gorm:"type:int;not null" json:"price" form:"price"`
	Additional_description string `gorm:"type:varchar(100);not null" json:"additional_description" form:"additional_description"`
}
type RespondRoom struct {
	ID                     uint   `json:"id"`
	HouseID                uint   `json:"house_id"`
	Type                   string `json:"type"`
	Price                  int32  `json:"price"`
	Additional_description string `json:"additional_description"`
}

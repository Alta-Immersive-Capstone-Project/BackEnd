package entities

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	HouseID                uint          `json:"house_id"`
	Type                   string        `json:"type"`
	UserID                 uint          `json:"userid"`
	Price                  int64         `json:"price"`
	Additional_description string        `json:"additional_description"`
	Transactions           []Transaction `gorm:"foreingkey:RoomID"`
	Images                 []Image       `gorm:"foreingkey:RoomID"`
	Amenities              Amenities     `gorm:"foreingkey:RoomID"`
}

type AddRoom struct {
	HouseID                uint   `json:"house_id" validate:"required" form:"house_id"`
	Type                   string `gorm:"type:varchar(100);not null" json:"type" validate:"required" form:"type"`
	Price                  int64  `gorm:"type:int;not null" json:"price" validate:"required" form:"price"`
	Additional_description string `gorm:"type:varchar(100);not null" json:"additional_description" validate:"required" form:"additional_description"`
}
type UpdateRoom struct {
	Type                   string `gorm:"type:varchar(100);not null" json:"type" form:"type"`
	Price                  int64  `gorm:"type:int;not null" json:"price" form:"price"`
	Additional_description string `gorm:"type:varchar(100);not null" json:"additional_description" form:"additional_description"`
}
type RespondRoom struct {
	ID                     uint   `json:"id"`
	HouseID                uint   `json:"house_id"`
	Type                   string `json:"type"`
	Price                  int64  `json:"price"`
	Additional_description string `json:"additional_description"`
}
type RespondRoomcreat struct {
	ID                     uint   `json:"id"`
	HouseID                uint   `json:"house_id"`
	Type                   string `json:"type"`
	Price                  int64  `json:"price"`
	Additional_description string `json:"additional_description"`
	Images                 []Images
}
type RespondRoomJoin struct {
	ID                     uint   `json:"id"`
	Type                   string `json:"type"`
	Price                  int32  `json:"price"`
	Additional_description string `json:"additional_description"`
	Images                 []Images
	Amenities              Amenitiest
}
type Images struct {
	ID  uint   `json:"id"`
	Url string `json:"url"`
}
type Amenitiest struct {
	ID          uint   `json:"id"`
	Bathroom    string `json:"bathroom"`
	Bed         string `json:"bed"`
	AC          string `json:"ac"`
	Wardrobe    string `json:"wardrobe"`
	Electricity string `json:"electricity"`
}

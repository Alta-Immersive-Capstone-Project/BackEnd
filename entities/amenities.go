package entities

import "gorm.io/gorm"

type Amenities struct {
	gorm.Model
	Bathroom    string `json:"bathroom"`
	Bed         string `json:"bed"`
	AC          string `json:"ac"`
	Wardrobe    string `json:"wardrobe"`
	Electricity string `json:"electricity"`
	RoomID      uint   `json:"room_id"`
}

type AddAmenities struct {
	RoomID      uint   `json:"room_id"`
	Bathroom    string `json:"bathroom" validate:"required"`
	Bed         string `json:"bed" validate:"required"`
	AC          string `json:"ac" validate:"required"`
	Wardrobe    string `json:"wardrobe" validate:"required"`
	Electricity string `json:"electricity" validate:"required"`
}

type UpdateAmenities struct {
	Bathroom    string `json:"bathroom"`
	Bed         string `json:"bed"`
	AC          string `json:"ac"`
	Wardrobe    string `json:"wardrobe"`
	Electricity string `json:"electricity"`
}

type RespondAmenities struct {
	ID          uint   `json:"id"`
	Bathroom    string `json:"bathroom"`
	Bed         string `json:"bed"`
	AC          string `json:"ac"`
	Wardrobe    string `json:"wardrobe"`
	Electricity string `json:"electricity"`
	RoomID      uint   `json:"room_id"`
}

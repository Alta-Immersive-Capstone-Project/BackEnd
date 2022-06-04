package entities

import (
	"gorm.io/gorm"
)

type House struct {
	gorm.Model
	DistrictID uint     `gorm:"not null json:district_id"`
	Title      string   `json:"title" gorm:"type:varchar(161);not null"`
	Brief      string   `json:"brief" gorm:"type:text;not null"`
	OwnerName  string   `json:"owner_name" gorm:"type:varchar(255);not null"`
	OwnerPhone string   `json:"owner_phone" gorm:"type:varchar(15);not null"`
	Address    string   `json:"address" gorm:"type:varchar(255);not null"`
	Type       string   `json:"type" gorm:"type:varchar(255);not null"`
	SlotRoom   int      `json:"slot_room" gorm:"type:int(3);not null"`
	Available  int      `json:"available" gorm:"type:int(3);not null"`
	Latitude   float64  `json:"latitude" gorm:"not null"`
	Longitude  float64  `json:"longitude" gorm:"not null"`
	Image      string   `json:"image"`
	Room       []Room   `gorm:"foreingkey:HouseID"`
	Review     []Review `gorm:"foreingkey:HouseID"`
}

type AddHouse struct {
	DistrictID uint    `json:"district_id" validate:"required" form:"district_id"`
	Title      string  `json:"title" validate:"required" form:"title"`
	Brief      string  `json:"brief" validate:"required" form:"brief"`
	OwnerName  string  `json:"owner_name" validate:"required" form:"owner_name"`
	OwnerPhone string  `json:"owner_phone" validate:"required" form:"owner_phone"`
	Address    string  `json:"address" validate:"required" form:"address"`
	Type       string  `json:"type" validate:"required" form:"type"`
	SlotRoom   int     `json:"slot_room" validate:"required" form:"slot_room"`
	Available  int     `json:"available" validate:"required" form:"available"`
	Latitude   float64 `json:"latitude" validate:"required" form:"latitude"`
	Longitude  float64 `json:"longitude" validate:"required" form:"longitude"`
}

type UpdateHouse struct {
	DistrictID uint    `json:"district_id" form:"district_id"`
	Title      string  `json:"title" form:"title"`
	Brief      string  `json:"brief" form:"brief"`
	OwnerName  string  `json:"owner_name" form:"owner_name"`
	OwnerPhone string  `json:"owner_phone" form:"owner_phone"`
	Address    string  `json:"address" form:"address"`
	Type       string  `json:"type" form:"type"`
	Image      string  `json:"image"`
	SlotRoom   int     `json:"slot_room" form:"slot_room"`
	Available  int     `json:"available" form:"available"`
	Latitude   float64 `json:"latitude" form:"latitude"`
	Longitude  float64 `json:"longitude" form:"longitude"`
}

type HouseResponse struct {
	ID         uint    `json:"house_id"`
	Title      string  `json:"title"`
	Brief      string  `json:"brief"`
	OwnerName  string  `json:"owner_name"`
	OwnerPhone string  `json:"owner_phone"`
	Address    string  `json:"address"`
	SlotRoom   int     `json:"slot_room"`
	Type       string  `json:"type"`
	Image      string  `json:"image"`
	Available  int     `json:"available"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	DistrictID uint    `json:"district_id"`
	District   string  `json:"district"`
	CityID     uint    `json:"city_id"`
	City       string  `json:"city_name"`
	Rooms      []RespondRoomJoin
}
type HouseResponseGetByID struct {
	ID         uint    `json:"house_id"`
	Title      string  `json:"title"`
	Brief      string  `json:"brief"`
	OwnerName  string  `json:"owner_name"`
	OwnerPhone string  `json:"owner_phone"`
	Address    string  `json:"address"`
	SlotRoom   int     `json:"slot_room"`
	Type       string  `json:"type"`
	Image      string  `json:"image"`
	Available  int     `json:"available"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	DistrictID uint    `json:"district_id"`
	District   string  `json:"district"`
	CityID     uint    `json:"city_id"`
	City       string  `json:"city"`
}
type HouseResponseGet struct {
	ID         uint    `json:"house_id"`
	Title      string  `json:"title"`
	Brief      string  `json:"brief"`
	OwnerName  string  `json:"owner_name"`
	OwnerPhone string  `json:"owner_phone"`
	Address    string  `json:"address"`
	SlotRoom   int     `json:"slot_room"`
	Type       string  `json:"type"`
	Image      string  `json:"image"`
	Available  int     `json:"available"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	DistrictID uint    `json:"district_id"`
	District   string  `json:"district"`
	CityID     uint    `json:"city_id"`
	City       string  `json:"city"`
	Price      int32   `json:"price"`
	Rating     float32 `json:"rating"`
}
type HouseResponseGetAll struct {
	ID         uint    `json:"house_id"`
	Title      string  `json:"title"`
	Address    string  `json:"address"`
	SlotRoom   int     `json:"slot_room"`
	Type       string  `json:"type"`
	Image      string  `json:"image"`
	Available  int     `json:"available"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	DistrictID uint    `json:"district_id"`
	District   string  `json:"district"`
	CityID     uint    `json:"city_id"`
	City       string  `json:"city"`
	Price      int32   `json:"price"`
	Rating     float32 `json:"rating"`
}

type HouseResponseJoin struct {
	ID         uint    `json:"house_id"`
	Title      string  `json:"title"`
	Brief      string  `json:"brief"`
	OwnerName  string  `json:"owner_name"`
	OwnerPhone string  `json:"owner_phone"`
	Address    string  `json:"address"`
	Available  int     `json:"available"`
	Type       string  `json:"type"`
	Image      string  `json:"image"`
	RoomID     uint    `json:"room_id"`
	RoomType   string  `json:"room_type"`
	Price      int32   `json:"price"`
	Rating     float32 `json:"rating"`
	DistrictID uint    `json:"district_id"`
	District   string  `json:"district"`
	// ImagesUrl  string  `json:"images_url"`
}

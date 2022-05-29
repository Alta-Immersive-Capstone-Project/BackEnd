package entities

import (
	"gorm.io/gorm"
)

type House struct {
	gorm.Model
	DistrictID uint       `gorm:"not null json:district_id"`
	Title      string     `json:"title" gorm:"type:varchar(161);not null;unique"`
	Brief      string     `json:"brief" gorm:"type:text;not null"`
	OwnerName  string     `json:"owner_name" gorm:"type:varchar(255);not null"`
	OwnerPhone string     `json:"owner_phone" gorm:"type:varchar(15);not null"`
	Address    string     `json:"address" gorm:"type:varchar(255);not null"`
	SlotRoom   int        `json:"slot_room" gorm:"type:int(3);not null"`
	Available  int        `json:"available" gorm:"type:int(3);not null"`
	Latitude   float64    `json:"latitude" gorm:"not null"`
	Longitude  float64    `json:"longitude" gorm:"not null"`
	Facility   []Facility `gorm:"foreingkey:HouseID"`
	Room       []Room     `gorm:"foreingkey:HouseID"`
}

type AddHouse struct {
	DistrictID uint    `json:"district_id"`
	Title      string  `json:"title"`
	Brief      string  `json:"brief"`
	OwnerName  string  `json:"owner_name"`
	OwnerPhone string  `json:"owner_phone"`
	Address    string  `json:"address"`
	SlotRoom   int     `json:"slot_room"`
	Available  int     `json:"available"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

type UpdateHouse struct {
	DistrictID uint    `json:"district_id"`
	Title      string  `json:"title"`
	Brief      string  `json:"brief"`
	OwnerName  string  `json:"owner_name"`
	OwnerPhone string  `json:"owner_phone"`
	Address    string  `json:"address"`
	SlotRoom   int     `json:"slot_room"`
	Available  int     `json:"available"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

type HouseResponse struct {
	ID         uint    `json:"house_id"`
	Title      string  `json:"title"`
	Brief      string  `json:"brief"`
	OwnerName  string  `json:"owner_name"`
	OwnerPhone string  `json:"owner_phone"`
	Address    string  `json:"address"`
	SlotRoom   int     `json:"slot_room"`
	Available  int     `json:"available"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	DistrictID uint    `json:"district_id"`
}

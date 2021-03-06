package room

import "kost/entities"

type RoomRepo interface {
	CreateRoom(new entities.Room) (entities.Room, error)
	GetAllRoom(id uint) ([]entities.Room, error)
	GetRoomID(id uint) (entities.Room, error)
	UpdateRoom(id uint, new entities.Room) (entities.Room, error)
	DeleteRoom(id uint) error
	GetbyHouse(id uint) ([]entities.RespondRoomJoin, error)
}

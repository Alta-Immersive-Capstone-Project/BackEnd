package room

import "kost/entities"

type RoomServices interface {
	CreateRoom(id uint, Insert entities.AddRoom) (entities.RespondRoomcreat, error)
	GetAllRoom(id uint) ([]entities.RespondRoom, error)
	GetIDRoom(id uint) (entities.Room, error)
	UpdateRoom(id uint, update entities.UpdateRoom) (entities.RespondRoom, error)
	DeleteRoom(id uint) error
}

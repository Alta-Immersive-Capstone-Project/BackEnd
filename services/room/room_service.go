package room

import (
	"kost/entities"
	"kost/repositories/image"
	"kost/repositories/room"

	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
)

type ServiceRoom struct {
	repo  room.RoomRepo
	image image.ImageRepo
}

func NewServiceRoom(Repo room.RoomRepo, image image.ImageRepo) *ServiceRoom {
	return &ServiceRoom{
		repo:  Repo,
		image: image,
	}
}

func (r *ServiceRoom) CreateRoom(id uint, Insert entities.AddRoom) (entities.RespondRoomcreat, error) {

	newRoom := entities.Room{}
	copier.Copy(&newRoom, &Insert)
	newRoom.UserID = id
	newRoom.Image = "https://belajar-be.s3.ap-southeast-1.amazonaws.com/room/1653973008.png"

	res, err := r.repo.CreateRoom(newRoom)
	if err != nil {
		log.Warn(err)
		return entities.RespondRoomcreat{}, err
	}

	result := entities.RespondRoomcreat{}
	copier.Copy(&result, &res)

	return result, nil
}

func (s *ServiceRoom) GetAllRoom() ([]entities.RespondRoom, error) {

	res, err := s.repo.GetAllRoom()
	if err != nil {
		log.Warn(err)
		return []entities.RespondRoom{}, err
	}

	result := []entities.RespondRoom{}
	copier.Copy(&result, &res)

	return result, nil
}
func (s *ServiceRoom) GetIDRoom(id uint) (entities.Room, error) {

	res, err := s.repo.GetRoomID(id)
	if err != nil {
		log.Warn(err)
		return entities.Room{}, err
	}

	return res, nil
}

func (s *ServiceRoom) UpdateRoom(id uint, update entities.UpdateRoom) (entities.RespondRoom, error) {

	var updateRoom entities.Room
	copier.Copy(&updateRoom, &update)

	res, err := s.repo.UpdateRoom(id, updateRoom)
	if err != nil {
		log.Warn(err)
		return entities.RespondRoom{}, err
	}

	result := entities.RespondRoom{}
	copier.Copy(&result, &res)

	return result, nil
}

func (s *ServiceRoom) DeleteRoom(id uint) error {

	err := s.repo.DeleteRoom(id)
	if err != nil {
		log.Warn(err)
		return err
	}

	return nil
}

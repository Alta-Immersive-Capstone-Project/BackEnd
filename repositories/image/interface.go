package image

import "kost/entities"

type ImageRepo interface {
	CreateImage(new entities.Image) error
	DeleteImage(imageID uint) error
	GetAllImage(roomID uint) ([]entities.Image, error)
	GetImage(ID uint) (entities.Image, error)
}

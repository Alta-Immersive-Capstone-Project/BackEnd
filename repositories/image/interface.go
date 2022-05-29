package image

import "kost/entities"

type ImageRepo interface {
	CreateImage(new entities.Image) error
	DeleteImage(roomID, userID uint) error
}

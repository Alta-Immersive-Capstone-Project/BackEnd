package image

import "mime/multipart"

type ImageService interface {
	InsertImage(files []*multipart.FileHeader, id uint) error
	DeleteImage(id uint) error
	DeleteImagebyID(id_room []int) error
}

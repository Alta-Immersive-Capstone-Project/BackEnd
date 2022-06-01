package image

import (
	"fmt"
	"kost/deliveries/helpers"
	"kost/entities"
	"kost/repositories/image"
	"kost/repositories/room"
	"kost/utils/s3"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

type ServiceImage struct {
	image image.ImageRepo
	repo  room.RoomRepo
	s3    s3.S3Control
}

func NewServiceImage(Repo room.RoomRepo, image image.ImageRepo, S3 s3.S3Control) *ServiceImage {
	return &ServiceImage{
		image: image,
		repo:  Repo,
		s3:    S3,
	}
}

func (i *ServiceImage) InsertImage(files []*multipart.FileHeader, id uint) error {
	for index, fileU := range files {
		src, err := fileU.Open()
		defer src.Close()
		if err != nil {
			return err
		}
		fileName := "room/" + strconv.Itoa(int(time.Now().Unix())) + ".png"
		res, err := i.s3.UploadFileToS3(fileName, *fileU)
		fmt.Println(res)
		if err != nil {
			return err
		}
		if index == 0 {
			i.repo.UpdateRoom(id, entities.Room{Image: res})
		}

		img := entities.Image{
			RoomID: id,
			Url:    res,
		}
		err = i.image.CreateImage(img)
		if err != nil {
			return err
		}
	}
	return nil
}
func (i *ServiceImage) DeleteImage(id_room uint) error {

	result, err := i.image.GetAllImage(id_room)
	if err != nil {
		return err
	}

	for _, res := range result {
		file := strings.Replace(res.Url, "https://belajar-be.s3.ap-southeast-1.amazonaws.com/", "", 1)
		err = i.s3.DeleteFromS3(file)
		if err != nil {
			return err
		}
		err = i.image.DeleteImage(res.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
func (i *ServiceImage) DeleteImagebyID(id_room []uint) error {

	for _, res := range id_room {
		fmt.Println(res)
		result, err := i.image.GetImage(uint(res))
		if err != nil {
			return err
		}
		file := strings.Replace(result.Url, "https://belajar-be.s3.ap-southeast-1.amazonaws.com/", "", 1)
		err = helpers.DeleteFromS3(file)
		if err != nil {
			return err
		}
		err = i.image.DeleteImage(result.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

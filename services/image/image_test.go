package image

import (
	"errors"
	"kost/entities"
	mocks "kost/mocks/repositories/image"
	roomMock "kost/mocks/repositories/room"
	s3mock "kost/mocks/utils/s3"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// func TestInsertImage(t *testing.T) {
// 	var file []*multipart.FileHeader
// 	t.Run("Success Create Image", func(t *testing.T) {
// 		Image := mocks.NewImageRepo(t)
// 		RoomRepo := roomMock.NewRoomRepo(t)
// 		s3 := s3mock.NewS3Control(t)
// 		Image.On("CreateImage", mock.Anything).Return(nil).Once()
// 		s3.On("UploadFileToS3", mock.Anything, mock.Anything).Return("url", nil).Once()
// 		imageService := NewServiceImage(RoomRepo, Image, s3)
// 		err := imageService.InsertImage(file, uint(1))
// 		assert.NoError(t, err)

// 		Image.AssertExpectations(t)
// 		RoomRepo.AssertExpectations(t)
// 		s3.AssertExpectations(t)
// 	})
// }

func TestDeleteImage(t *testing.T) {
	result := []entities.Image{{RoomID: uint(1)}}
	t.Run("Success Delete Image", func(t *testing.T) {
		Image := mocks.NewImageRepo(t)
		RoomRepo := roomMock.NewRoomRepo(t)
		s3 := s3mock.NewS3Control(t)
		Image.On("GetAllImage", uint(1)).Return(result, nil).Once()
		s3.On("DeleteFromS3", "").Return(nil).Once()
		Image.On("DeleteImage", mock.Anything).Return(nil).Once()
		imageService := NewServiceImage(RoomRepo, Image, s3)
		err := imageService.DeleteImage(uint(1))
		assert.NoError(t, err)

		Image.AssertExpectations(t)
		RoomRepo.AssertExpectations(t)
		s3.AssertExpectations(t)
	})
	t.Run("Error Get Image", func(t *testing.T) {
		Image := mocks.NewImageRepo(t)
		RoomRepo := roomMock.NewRoomRepo(t)
		s3 := s3mock.NewS3Control(t)
		Image.On("GetAllImage", uint(1)).Return(result, errors.New("Error Get Image")).Once()

		imageService := NewServiceImage(RoomRepo, Image, s3)
		err := imageService.DeleteImage(uint(1))
		assert.Error(t, err)

		Image.AssertExpectations(t)
		RoomRepo.AssertExpectations(t)
		s3.AssertExpectations(t)
	})
	t.Run("Error Delete S3", func(t *testing.T) {
		Image := mocks.NewImageRepo(t)
		RoomRepo := roomMock.NewRoomRepo(t)
		s3 := s3mock.NewS3Control(t)
		Image.On("GetAllImage", uint(1)).Return(result, nil).Once()
		s3.On("DeleteFromS3", "").Return(errors.New("DeleteFromS3")).Once()
		imageService := NewServiceImage(RoomRepo, Image, s3)
		err := imageService.DeleteImage(uint(1))
		assert.Error(t, err)

		Image.AssertExpectations(t)
		RoomRepo.AssertExpectations(t)
		s3.AssertExpectations(t)
	})
	t.Run("Error Delete Image", func(t *testing.T) {
		Image := mocks.NewImageRepo(t)
		RoomRepo := roomMock.NewRoomRepo(t)
		s3 := s3mock.NewS3Control(t)
		Image.On("GetAllImage", uint(1)).Return(result, nil).Once()
		s3.On("DeleteFromS3", "").Return(nil).Once()
		Image.On("DeleteImage", mock.Anything).Return(errors.New("Error Delete Image")).Once()
		imageService := NewServiceImage(RoomRepo, Image, s3)
		err := imageService.DeleteImage(uint(1))
		assert.Error(t, err)

		Image.AssertExpectations(t)
		RoomRepo.AssertExpectations(t)
		s3.AssertExpectations(t)
	})
}

func TestDeleteImageByID(t *testing.T) {
	result := entities.Image{RoomID: uint(1)}
	id := []uint{1}
	t.Run("Success Delete Image", func(t *testing.T) {
		Image := mocks.NewImageRepo(t)
		RoomRepo := roomMock.NewRoomRepo(t)
		s3 := s3mock.NewS3Control(t)
		Image.On("GetImage", uint(1)).Return(result, nil).Once()
		s3.On("DeleteFromS3", "").Return(nil).Once()
		Image.On("DeleteImage", mock.Anything).Return(nil).Once()
		imageService := NewServiceImage(RoomRepo, Image, s3)
		err := imageService.DeleteImagebyID(id)
		assert.NoError(t, err)

		Image.AssertExpectations(t)
		RoomRepo.AssertExpectations(t)
		s3.AssertExpectations(t)
	})
	t.Run("Error Get Image", func(t *testing.T) {
		Image := mocks.NewImageRepo(t)
		RoomRepo := roomMock.NewRoomRepo(t)
		s3 := s3mock.NewS3Control(t)
		Image.On("GetImage", mock.Anything).Return(result, errors.New("Error Get Image")).Once()

		imageService := NewServiceImage(RoomRepo, Image, s3)
		err := imageService.DeleteImagebyID(id)
		assert.Error(t, err)

		Image.AssertExpectations(t)
		RoomRepo.AssertExpectations(t)
		s3.AssertExpectations(t)
	})
	t.Run("Error Delete S3", func(t *testing.T) {
		Image := mocks.NewImageRepo(t)
		RoomRepo := roomMock.NewRoomRepo(t)
		s3 := s3mock.NewS3Control(t)
		Image.On("GetImage", mock.Anything).Return(result, nil).Once()
		s3.On("DeleteFromS3", "").Return(errors.New("DeleteFromS3")).Once()
		imageService := NewServiceImage(RoomRepo, Image, s3)
		err := imageService.DeleteImagebyID(id)
		assert.Error(t, err)

		Image.AssertExpectations(t)
		RoomRepo.AssertExpectations(t)
		s3.AssertExpectations(t)
	})
	t.Run("Error Delete Image", func(t *testing.T) {
		Image := mocks.NewImageRepo(t)
		RoomRepo := roomMock.NewRoomRepo(t)
		s3 := s3mock.NewS3Control(t)
		Image.On("GetImage", mock.Anything).Return(result, nil).Once()
		s3.On("DeleteFromS3", "").Return(nil).Once()
		Image.On("DeleteImage", mock.Anything).Return(errors.New("Error Delete Image")).Once()
		imageService := NewServiceImage(RoomRepo, Image, s3)
		err := imageService.DeleteImagebyID(id)
		assert.Error(t, err)

		Image.AssertExpectations(t)
		RoomRepo.AssertExpectations(t)
		s3.AssertExpectations(t)
	})
}

func TestGetImage(t *testing.T) {
	result := []entities.Image{{Model: gorm.Model{ID: uint(1)}}}
	id := uint(1)
	t.Run("Success Delete Image", func(t *testing.T) {
		Image := mocks.NewImageRepo(t)
		RoomRepo := roomMock.NewRoomRepo(t)
		s3 := s3mock.NewS3Control(t)
		Image.On("GetAllImage", uint(1)).Return(result, nil).Once()
		imageService := NewServiceImage(RoomRepo, Image, s3)
		respond, err := imageService.GetImage(id)
		assert.NoError(t, err)
		assert.Equal(t, result[0].ID, respond[0].ID)
		Image.AssertExpectations(t)
		RoomRepo.AssertExpectations(t)
		s3.AssertExpectations(t)
	})
	t.Run("Error Get Image", func(t *testing.T) {
		Image := mocks.NewImageRepo(t)
		RoomRepo := roomMock.NewRoomRepo(t)
		s3 := s3mock.NewS3Control(t)
		Image.On("GetAllImage", uint(1)).Return(result, errors.New("Error Get Image")).Once()
		imageService := NewServiceImage(RoomRepo, Image, s3)
		_, err := imageService.GetImage(id)
		assert.Error(t, err)

		Image.AssertExpectations(t)
		RoomRepo.AssertExpectations(t)
		s3.AssertExpectations(t)
	})
}

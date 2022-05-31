package reviews_test

import (
	"errors"
	"kost/entities"
	repoMock "kost/mocks/repositories/reviews"
	review "kost/services/reviews"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAddComment(t *testing.T) {
	repo := repoMock.NewReviewModel(t)
	insertData := entities.ReviewRequest{RoomID: 1, Comment: "test", Rating: 5}
	returnData := entities.Review{Model: gorm.Model{ID: uint(1), CreatedAt: time.Now()}, UserID: 1, RoomID: 1, Comment: "test", Rating: 5}

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("Create", mock.Anything).Return(returnData, nil).Once()
		srv := review.NewReviewService(repo)

		res, err := srv.AddComment(1, insertData)
		assert.NoError(t, err)
		assert.Equal(t, returnData.ID, res.ID)
		assert.Equal(t, returnData.RoomID, res.RoomID)
		assert.Equal(t, returnData.Comment, res.Comment)
		assert.Equal(t, returnData.Rating, res.Rating)
		assert.Equal(t, returnData.CreatedAt, res.CreatedAt)
		repo.AssertExpectations(t)
	})

	t.Run("Error Insert", func(t *testing.T) {
		repo.On("Create", mock.Anything).Return(entities.Review{}, errors.New("there is some error")).Once()
		srv := review.NewReviewService(repo)

		_, err := srv.AddComment(1, insertData)
		assert.Error(t, err)
		assert.EqualError(t, err, "there is some error")
		repo.AssertExpectations(t)
	})
}

func TestGetByRoomID(t *testing.T) {
	repo := new(repoMock.ReviewModel)
	userData := entities.User{Model: gorm.Model{ID: uint(1)}, Name: "test"}
	returnData := []entities.Review{{Model: gorm.Model{ID: uint(1), CreatedAt: time.Now()}, UserID: uint(1), RoomID: uint(1), Comment: "test", Rating: 5}}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetByRoomID", uint(1)).Return(returnData, nil).Once()
		repo.On("GetByUserID", uint(1)).Return(userData, nil).Once()
		srv := review.NewReviewService(repo)

		res, err := srv.GetByRoomID(uint(1))
		assert.NoError(t, err)
		// assert.Equal(t, returnData[0].ID, res[0].ID)
		// assert.Equal(t, returnData[0].UserID, res[0].UserID)
		assert.Equal(t, userData.Name, res[0].Name)
		// assert.Equal(t, returnData[0].RoomID, res[0].RoomID)
		assert.Equal(t, returnData[0].Comment, res[0].Comment)
		assert.Equal(t, returnData[0].Rating, res[0].Rating)
		assert.Equal(t, returnData[0].CreatedAt, res[0].CreatedAt)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get All", func(t *testing.T) {
		repo.On("GetByRoomID", uint(1)).Return(nil, errors.New("data not found")).Once()
		srv := review.NewReviewService(repo)

		res, err := srv.GetByRoomID(uint(1))
		assert.Error(t, err)
		assert.EqualError(t, err, "data not found")
		assert.Equal(t, []entities.ReviewGetResponse{}, res)
		repo.AssertExpectations(t)
	})
}

func TestGetRating(t *testing.T) {
	repo := new(repoMock.ReviewModel)
	returnData := float32(5)

	t.Run("Success Get Rating", func(t *testing.T) {
		repo.On("GetRating", uint(1)).Return(returnData, nil).Once()
		srv := review.NewReviewService(repo)

		res, _ := srv.GetRating(uint(1))
		assert.Equal(t, returnData, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get Rating", func(t *testing.T) {
		repo.On("GetRating", uint(1)).Return(float32(0), errors.New("data not found")).Once()
		srv := review.NewReviewService(repo)

		res, _ := srv.GetRating(uint(1))
		assert.Equal(t, float32(0), res)
		repo.AssertExpectations(t)
	})
}

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
	insertData := entities.ReviewRequest{HouseID: 1, Comment: "test", Rating: 5}
	returnData := entities.Review{Model: gorm.Model{ID: uint(1), CreatedAt: time.Now()}, UserID: 1, HouseID: 1, Comment: "test", Rating: 5}

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("Create", mock.Anything).Return(returnData, nil).Once()
		srv := review.NewReviewService(repo)

		res, err := srv.AddComment(1, insertData)
		assert.NoError(t, err)
		assert.Equal(t, returnData.ID, res.ID)
		assert.Equal(t, returnData.HouseID, res.HouseID)
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

func TestGetByHouseID(t *testing.T) {
	repo := new(repoMock.ReviewModel)
	returnData := []entities.ReviewJoin{{Name: "test", Comment: "test", Rating: 5, CreatedAt: time.Now()}}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetByHouseID", uint(1)).Return(returnData, nil).Once()
		srv := review.NewReviewService(repo)

		res, err := srv.GetByHouseID(uint(1))
		assert.NoError(t, err)
		assert.Equal(t, returnData[0].Name, res[0].Name)
		assert.Equal(t, returnData[0].Comment, res[0].Comment)
		assert.Equal(t, returnData[0].Rating, res[0].Rating)
		assert.Equal(t, returnData[0].CreatedAt, res[0].CreatedAt)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get All", func(t *testing.T) {
		repo.On("GetByHouseID", uint(1)).Return(nil, errors.New("data not found")).Once()
		srv := review.NewReviewService(repo)

		res, err := srv.GetByHouseID(uint(1))
		assert.Error(t, err)
		assert.EqualError(t, err, "data not found")
		assert.Equal(t, []entities.ReviewJoin{}, res)
		repo.AssertExpectations(t)
	})
}

func TestGetRating(t *testing.T) {
	repo := new(repoMock.ReviewModel)
	returnArray := []int{1, 2, 3, 4, 5}
	returnData := float32(5)

	t.Run("Success Get Rating", func(t *testing.T) {
		repo.On("GetRating", uint(1)).Return(returnArray, returnData, nil).Once()
		srv := review.NewReviewService(repo)

		count, total := srv.GetRating(uint(1))
		assert.Equal(t, returnArray, count)
		assert.Equal(t, returnData, total)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get Rating", func(t *testing.T) {
		repo.On("GetRating", uint(1)).Return([]int{}, float32(0), errors.New("data not found")).Once()
		srv := review.NewReviewService(repo)

		count, total := srv.GetRating(uint(1))
		assert.Equal(t, []int{}, count)
		assert.Equal(t, float32(0), total)
		repo.AssertExpectations(t)
	})
}

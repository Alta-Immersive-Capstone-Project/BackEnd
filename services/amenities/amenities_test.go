package amenities

import (
	"errors"
	"kost/entities"

	mocks "kost/mocks/repositories/amenities"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// Create Sample Data
var MockAmenities = []entities.Amenities{
	{
		Model:       gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		RoomID:      1,
		Bathroom:    "yes",
		Bed:         "yes",
		AC:          "yes",
		Wardrobe:    "yes",
		Electricity: "yes",
	},
	{
		Model:       gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		RoomID:      1,
		Bathroom:    "yes",
		Bed:         "yes",
		AC:          "no",
		Wardrobe:    "no",
		Electricity: "yes",
	},
}

// TEST CREATE FACILITY SERVICE
func TestCreateAmenities(t *testing.T) {

	var NewAmenities = entities.AddAmenities{
		RoomID:      1,
		Bathroom:    "yes",
		Bed:         "yes",
		AC:          "yes",
		Wardrobe:    "yes",
		Electricity: "yes",
	}
	t.Run("Success Create Amenities", func(t *testing.T) {

		AmenitiesRepo := mocks.NewRepoAmenities(t)

		AmenitiesRepo.On("CreateAmenities", mock.Anything).Return(MockAmenities[0], nil).Once()

		AmenitiesService := NewServiceAmenities(AmenitiesRepo)
		result, err := AmenitiesService.CreateAmenities(NewAmenities)
		assert.NoError(t, err)
		assert.Equal(t, MockAmenities[0].Bathroom, result.Bathroom)

		AmenitiesRepo.AssertExpectations(t)
	})
	t.Run("Error Create Amenities", func(t *testing.T) {
		AmenitiesRepo := mocks.NewRepoAmenities(t)

		AmenitiesRepo.On("CreateAmenities", mock.Anything).Return(entities.Amenities{}, errors.New("Error Create Amenities")).Once()

		AmenitiesService := NewServiceAmenities(AmenitiesRepo)
		result, err := AmenitiesService.CreateAmenities(NewAmenities)
		assert.Error(t, err)
		assert.NotEqual(t, MockAmenities[0].Bathroom, result.Bathroom)

		AmenitiesRepo.AssertExpectations(t)
	})
}

// TEST GET FACILITY BY ID SERVICES
func TestGetAmenitiesID(t *testing.T) {
	t.Run("Success Get Amenities ID", func(t *testing.T) {
		AmenitiesRepo := mocks.NewRepoAmenities(t)
		AmenitiesRepo.On("GetAmenitiesID", uint(1)).Return(MockAmenities[0], nil).Once()
		AmenitiesService := NewServiceAmenities(AmenitiesRepo)
		result, err := AmenitiesService.GetAmenitiesID(1)
		assert.NoError(t, err)
		assert.Equal(t, MockAmenities[0].Bathroom, result.Bathroom)

		AmenitiesRepo.AssertExpectations(t)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		AmenitiesRepo := mocks.NewRepoAmenities(t)
		AmenitiesRepo.On("GetAmenitiesID", uint(1)).Return(entities.Amenities{}, errors.New("Error Access Database")).Once()
		AmenitiesService := NewServiceAmenities(AmenitiesRepo)
		result, err := AmenitiesService.GetAmenitiesID(1)
		assert.Error(t, err)
		assert.NotEqual(t, MockAmenities[0].Bathroom, result.Bathroom)

		AmenitiesRepo.AssertExpectations(t)
	})
}

// TEST UPDATE FACILITY SERVICE
func TestUpdateAmenities(t *testing.T) {
	var respon = entities.Amenities{
		Model:       gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		RoomID:      1,
		Bathroom:    "yes",
		Bed:         "yes",
		AC:          "yes",
		Wardrobe:    "yes",
		Electricity: "yes",
	}
	var UpdateAmenities = entities.UpdateAmenities{
		Bathroom:    "yes",
		Bed:         "yes",
		AC:          "yes",
		Wardrobe:    "yes",
		Electricity: "yes",
	}
	t.Run("Success Update Amenities", func(t *testing.T) {

		AmenitiesRepo := mocks.NewRepoAmenities(t)

		AmenitiesRepo.On("UpdateAmenities", uint(1), mock.Anything).Return(respon, nil).Once()

		AmenitiesService := NewServiceAmenities(AmenitiesRepo)
		result, err := AmenitiesService.UpdateAmenities(uint(1), UpdateAmenities)
		assert.NoError(t, err)
		assert.Equal(t, respon.Bathroom, result.Bathroom)
		AmenitiesRepo.AssertExpectations(t)
	})
	t.Run("Error Update Amenities", func(t *testing.T) {

		AmenitiesRepo := mocks.NewRepoAmenities(t)

		AmenitiesRepo.On("UpdateAmenities", uint(1), mock.Anything).Return(entities.Amenities{}, errors.New("Error Access Database")).Once()

		AmenitiesService := NewServiceAmenities(AmenitiesRepo)
		result, err := AmenitiesService.UpdateAmenities(uint(1), UpdateAmenities)
		assert.Error(t, err)
		assert.NotEqual(t, respon.Bathroom, result.Bathroom)

		AmenitiesRepo.AssertExpectations(t)
	})
}

// TEST GET FACILITY BY ID SERVICES
func TestDeleteAmenities(t *testing.T) {
	t.Run("Success Delete Amenities ID", func(t *testing.T) {
		AmenitiesRepo := mocks.NewRepoAmenities(t)
		AmenitiesRepo.On("DeleteAmenities", uint(1)).Return(nil).Once()
		AmenitiesService := NewServiceAmenities(AmenitiesRepo)
		err := AmenitiesService.DeleteAmenities(1)
		assert.NoError(t, err)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		AmenitiesRepo := mocks.NewRepoAmenities(t)
		AmenitiesRepo.On("DeleteAmenities", uint(1)).Return(errors.New("Error Access Database")).Once()
		AmenitiesService := NewServiceAmenities(AmenitiesRepo)
		err := AmenitiesService.DeleteAmenities(1)
		assert.Error(t, err)
	})
}

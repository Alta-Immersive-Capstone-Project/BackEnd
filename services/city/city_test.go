package city

import (
	"errors"
	"kost/entities"
	mocks "kost/mocks/repositories/city"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var mockCities = []entities.City{
	{
		Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		City:  "Sleman",
	},
	{
		Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		City:  "Kulon Progo",
	},
	{
		Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		City:  "Bantul",
	},
}

func TestCreateCity(t *testing.T) {
	var NewCity = entities.AddCity{
		City: "Sleman",
	}

	t.Run("Success Create City", func(t *testing.T) {
		CityRepo := mocks.NewCityRepo(t)
		CityRepo.On("CreateCity", mock.Anything).Return(mockCities[0], nil).Once()
		cityService := NewServiceCity(CityRepo)

		result, err := cityService.CreateCity(NewCity)
		assert.NoError(t, err)
		assert.Equal(t, mockCities[0].City, result.City)

		CityRepo.AssertExpectations(t)
	})

	t.Run("Error Create City", func(t *testing.T) {
		CityRepo := mocks.NewCityRepo(t)
		CityRepo.On("CreateCity", mock.Anything).Return(entities.City{}, errors.New("error create city")).Once()
		cityService := NewServiceCity(CityRepo)

		result, err := cityService.CreateCity(NewCity)
		assert.Error(t, err)
		assert.NotEqual(t, mockCities[0].City, result.City)

		CityRepo.AssertExpectations(t)
	})
}

func TestUpdateCity(t *testing.T) {
	var response = entities.City{
		Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		City:  "Gunung Kidul",
	}
	var UpdateCity = entities.City{
		City: "Gunung Kidul",
	}

	t.Run("Success Update City", func(t *testing.T) {
		CityRepo := mocks.NewCityRepo(t)
		CityRepo.On("UpdateCity", uint(1), mock.Anything).Return(response, nil).Once()

		cityService := NewServiceCity(CityRepo)
		result, err := cityService.UpdateCity(uint(1), UpdateCity)
		assert.NoError(t, err)
		assert.Equal(t, response.City, result.City)

		CityRepo.AssertExpectations(t)
	})

	t.Run("Error Update City", func(t *testing.T) {
		CityRepo := mocks.NewCityRepo(t)
		CityRepo.On("UpdateCity", uint(1), mock.Anything).Return(entities.CityResponse{}, errors.New("error update cities")).Once()

		cityService := NewServiceCity(CityRepo)
		result, err := cityService.UpdateCity(uint(1), UpdateCity)
		assert.Error(t, err)
		assert.NotEqual(t, response.City, result.City)

		CityRepo.AssertExpectations(t)
	})
}

func TestDeleteCity(t *testing.T) {
	t.Run("Success Delete City", func(t *testing.T) {
		CityRepo := mocks.NewCityRepo(t)
		CityRepo.On("DeleteCity", uint(1)).Return(nil).Once()

		cityService := NewServiceCity(CityRepo)
		err := cityService.DeleteCity(uint(1))
		assert.NoError(t, err)
	})
	t.Run("Error Delete City", func(t *testing.T) {
		CityRepo := mocks.NewCityRepo(t)
		CityRepo.On("DeleteCity", uint(1)).Return(errors.New("error access database")).Once()

		cityService := NewServiceCity(CityRepo)
		err := cityService.DeleteCity(uint(1))
		assert.Error(t, err)
	})
}

func TestGetAllCity(t *testing.T) {
	t.Run("Success Get All City", func(t *testing.T) {
		CityRepo := mocks.NewCityRepo(t)
		CityRepo.On("GetAllCity", mock.Anything).Return(mockCities, nil).Once()
		cityService := NewServiceCity(CityRepo)

		result, err := cityService.GetAllCity()
		assert.NoError(t, err)
		assert.NotNil(t, result)

		CityRepo.AssertExpectations(t)
	})
	t.Run("Error Get All City", func(t *testing.T) {
		CityRepo := mocks.NewCityRepo(t)
		CityRepo.On("GetAllCity", mock.Anything).Return(entities.CityResponse{}, errors.New("error access database")).Once()
		cityService := NewServiceCity(CityRepo)

		_, err := cityService.GetAllCity()
		assert.Error(t, err)

		CityRepo.AssertExpectations(t)
	})
}

func TestGetIDCity(t *testing.T) {
	t.Run("Success Get ID City", func(t *testing.T) {
		CityRepo := mocks.NewCityRepo(t)
		CityRepo.On("GetIDCity", uint(1)).Return(mockCities[0].City, nil).Once()
		cityService := NewServiceCity(CityRepo)

		result, err := cityService.GetIDCity(uint(1))
		assert.NoError(t, err)
		assert.Equal(t, mockCities[0].City, result.City)

		CityRepo.AssertExpectations(t)
	})
	t.Run("Error Get ID City", func(t *testing.T) {
		CityRepo := mocks.NewCityRepo(t)
		CityRepo.On("GetIDCity", mock.Anything).Return(entities.CityResponse{}, errors.New("error access database")).Once()
		cityService := NewServiceCity(CityRepo)
		result, err := cityService.GetIDCity(uint(1))

		assert.Error(t, err)
		assert.NotEqual(t, mockCities[0].City, result.City)

		CityRepo.AssertExpectations(t)
	})
}

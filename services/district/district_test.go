package district

import (
	"errors"
	"kost/entities"
	"testing"
	"time"

	mocks "kost/mocks/repositories/district"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var MockDistrict = []entities.District{
	{
		Model:     gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:      "Catur Tunggal",
		Longitude: -6.168273696181832,
		Latitude:  106.86491520706296,
		CityID:    1,
	},
	{
		Model:     gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:      "Condong Catur",
		Longitude: -6.168273696181832,
		Latitude:  106.86491520706296,
		CityID:    1,
	},
	{
		Model:     gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:      "Minomartani",
		Longitude: -6.168273696181832,
		Latitude:  106.86491520706296,
		CityID:    1,
	},
}

func TestCreateDist(t *testing.T) {
	var NewDistrict = entities.AddDistrict{
		Name:      "Catur Tunggal",
		Longitude: -6.168273696181832,
		Latitude:  106.86491520706296,
		CityID:    1,
	}

	t.Run("Success Create District", func(t *testing.T) {
		DistrictRepo := mocks.NewRepoDistrict(t)
		DistrictRepo.On("CreateDistrict", mock.Anything).Return(MockDistrict[0], nil).Once()
		DistrictService := NewDistService(DistrictRepo)

		result, err := DistrictService.CreateDist(NewDistrict)
		assert.NoError(t, err)
		assert.Equal(t, MockDistrict[0].Name, result.Name)

		DistrictRepo.AssertExpectations(t)
	})

	t.Run("Error Create District", func(t *testing.T) {
		DistrictRepo := mocks.NewRepoDistrict(t)
		DistrictRepo.On("CreateDistrict", mock.Anything).Return(entities.District{}, errors.New("error create district")).Once()
		DistrictService := NewDistService(DistrictRepo)

		result, err := DistrictService.CreateDist(NewDistrict)
		assert.Error(t, err)
		assert.NotEqual(t, MockDistrict[0].Name, result.Name)

		DistrictRepo.AssertExpectations(t)
	})
}

func TestUpdateDistrict(t *testing.T) {
	var response = entities.District{
		Model:     gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:      "Maguwoharjo",
		Longitude: -6.168273696181832,
		Latitude:  106.86491520706296,
		CityID:    1,
	}
	var UpdateDistrict = entities.UpdateDistrict{
		Name:      "Maguwoharjo",
		Longitude: -6.168273696181832,
		Latitude:  106.86491520706296,
		CityID:    1,
	}
	t.Run("Success Update District", func(t *testing.T) {
		DistrictRepo := mocks.NewRepoDistrict(t)

		DistrictRepo.On("UpdateDistrict", uint(1), mock.Anything).Return(response, nil).Once()

		DistrictService := NewDistService(DistrictRepo)
		result, err := DistrictService.UpdateDist(uint(1), UpdateDistrict)
		assert.NoError(t, err)
		assert.Equal(t, response.Name, result.Name)

		DistrictRepo.AssertExpectations(t)
	})
	t.Run("Error Update District", func(t *testing.T) {
		DistrictRepo := mocks.NewRepoDistrict(t)

		DistrictRepo.On("UpdateDistrict", uint(1), mock.Anything).Return(entities.District{}, errors.New("error access database")).Once()

		DistrictService := NewDistService(DistrictRepo)
		result, err := DistrictService.UpdateDist(uint(1), UpdateDistrict)
		assert.Error(t, err)
		assert.NotEqual(t, response.Name, result.Name)

		DistrictRepo.AssertExpectations(t)

	})
}

func TestDeleteDistrict(t *testing.T) {
	t.Run("Success Delete District", func(t *testing.T) {
		DistrictRepo := mocks.NewRepoDistrict(t)
		DistrictRepo.On("DeleteDistrict", uint(1)).Return(nil).Once()
		DistrictService := NewDistService(DistrictRepo)
		err := DistrictService.DeleteDist(uint(1))
		assert.NoError(t, err)

		DistrictRepo.AssertExpectations(t)
	})

	t.Run("Error Delete District", func(t *testing.T) {
		DistrictRepo := mocks.NewRepoDistrict(t)
		DistrictRepo.On("DeleteDistrict", uint(1)).Return(errors.New("Error Access Database")).Once()
		DistrictService := NewDistService(DistrictRepo)
		err := DistrictService.DeleteDist(uint(1))
		assert.Error(t, err)

		DistrictRepo.AssertExpectations(t)
	})

}

func TestGetAllDistrict(t *testing.T) {
	t.Run("Success Get All District", func(t *testing.T) {
		DistrictRepo := mocks.NewRepoDistrict(t)
		DistrictRepo.On("GetAllDistrict", uint(1)).Return(MockDistrict, nil).Once()
		DistrictService := NewDistService(DistrictRepo)

		result, err := DistrictService.GetAllDist(uint(1))
		assert.NoError(t, err)
		assert.NotNil(t, result)

		DistrictRepo.AssertExpectations(t)
	})
	t.Run("Error Get All District", func(t *testing.T) {
		DistrictRepo := mocks.NewRepoDistrict(t)
		DistrictRepo.On("GetAllDistrict", uint(1)).Return(nil, errors.New("Error Access Database")).Once()
		DistrictService := NewDistService(DistrictRepo)

		result, err := DistrictService.GetAllDist(uint(1))
		assert.Error(t, err)
		assert.NotNil(t, result)

		DistrictRepo.AssertExpectations(t)
	})
}

func TestGetDistrictID(t *testing.T) {
	t.Run("Success Get District ID", func(t *testing.T) {
		DistrictRepo := mocks.NewRepoDistrict(t)
		DistrictRepo.On("GetDistrictID", uint(1)).Return(MockDistrict[0], nil).Once()
		DistrictService := NewDistService(DistrictRepo)

		result, err := DistrictService.GetDistID(uint(1))
		assert.NoError(t, err)
		assert.Equal(t, MockDistrict[0].Name, result.Name)

		DistrictRepo.AssertExpectations(t)
	})
	t.Run("Error Get District ID", func(t *testing.T) {
		DistrictRepo := mocks.NewRepoDistrict(t)
		DistrictRepo.On("GetDistrictID", uint(1)).Return(entities.District{}, errors.New("error access database")).Once()
		DistrictService := NewDistService(DistrictRepo)

		result, err := DistrictService.GetDistID(uint(1))
		assert.Error(t, err)
		assert.NotEqual(t, MockDistrict[0].Name, result.Name)

		DistrictRepo.AssertExpectations(t)
	})
}

func TestSelectAllDistrict(t *testing.T) {
	t.Run("Success Select All District", func(t *testing.T) {
		DistrictRepo := mocks.NewRepoDistrict(t)
		DistrictRepo.On("SelectAllDistrict", mock.Anything).Return(MockDistrict, nil).Once()
		DistrictService := NewDistService(DistrictRepo)
		res, err := DistrictService.SelectAllDistrict()
		assert.NoError(t, err)
		assert.Equal(t, MockDistrict, res)

		DistrictRepo.AssertExpectations(t)
	})
	t.Run("Error Select All District", func(t *testing.T) {
		DistrictRepo := mocks.NewRepoDistrict(t)
		DistrictRepo.On("SelectAllDistrict", mock.Anything).Return(entities.RespondDistrict{}, errors.New("error access database")).Once()
		DistrictService := NewDistService(DistrictRepo)
		_, err := DistrictService.SelectAllDistrict()
		assert.Error(t, err)

		DistrictRepo.AssertExpectations(t)
	})
}

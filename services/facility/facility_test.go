package facility

import (
	"errors"
	"kost/entities"

	mocks "kost/mocks/repositories/facility"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// Create Sample Data
var MockFacility = []entities.Facility{
	{
		Model:      gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:       "Rumah Sakit Primaya",
		Longitude:  -6.168273696181832,
		Latitude:   106.86491520706296,
		DistrictID: 1,
	},
	{
		Model:      gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:       "Grand Chandra Karya",
		Longitude:  -6.19395219376102,
		Latitude:   106.85925435178284,
		DistrictID: 1,
	},
}

// TEST CREATE FACILITY SERVICE
func TestCreateFacility(t *testing.T) {

	var NewFacility = entities.AddNewFacility{
		Name:       "Rumah Sakit Primaya",
		Longitude:  -6.168273696181832,
		Latitude:   106.86491520706296,
		DistrictID: 1,
	}
	t.Run("Success Create Facility", func(t *testing.T) {

		FacilityRepo := mocks.NewRepoFacility(t)

		FacilityRepo.On("CreateFacility", mock.Anything).Return(MockFacility[0], nil).Once()

		FacilityService := NewServiceFacility(FacilityRepo)
		result, err := FacilityService.CreateFacility(NewFacility)
		assert.NoError(t, err)
		assert.Equal(t, MockFacility[0].Name, result.Name)

		FacilityRepo.AssertExpectations(t)
	})
	t.Run("Error Create Facility", func(t *testing.T) {
		FacilityRepo := mocks.NewRepoFacility(t)

		FacilityRepo.On("CreateFacility", mock.Anything).Return(entities.Facility{}, errors.New("Error Create Facility")).Once()

		FacilityService := NewServiceFacility(FacilityRepo)
		result, err := FacilityService.CreateFacility(NewFacility)
		assert.Error(t, err)
		assert.NotEqual(t, MockFacility[0].Name, result.Name)

		FacilityRepo.AssertExpectations(t)
	})
}

// TEST GET ALL FACILITY SERVICES
func TestGetAllFacility(t *testing.T) {
	t.Run("Success Get All Facility", func(t *testing.T) {
		FacilityRepo := mocks.NewRepoFacility(t)
		FacilityRepo.On("GetAllFacility", uint(1)).Return(MockFacility, nil).Once()
		FacilityService := NewServiceFacility(FacilityRepo)
		result, err := FacilityService.GetAllFacility(1)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		FacilityRepo.AssertExpectations(t)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		FacilityRepo := mocks.NewRepoFacility(t)
		FacilityRepo.On("GetAllFacility", uint(1)).Return(nil, errors.New("Error Access Database")).Once()
		FacilityService := NewServiceFacility(FacilityRepo)
		_, err := FacilityService.GetAllFacility(1)
		assert.Error(t, err)
		FacilityRepo.AssertExpectations(t)
	})
}

// TEST GET FACILITY BY ID SERVICES
func TestGetFacilityID(t *testing.T) {
	t.Run("Success Get Facility ID", func(t *testing.T) {
		FacilityRepo := mocks.NewRepoFacility(t)
		FacilityRepo.On("GetFacilityID", uint(1)).Return(MockFacility[0], nil).Once()
		FacilityService := NewServiceFacility(FacilityRepo)
		result, err := FacilityService.GetFacilityID(1)
		assert.NoError(t, err)
		assert.Equal(t, MockFacility[0].Name, result.Name)

		FacilityRepo.AssertExpectations(t)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		FacilityRepo := mocks.NewRepoFacility(t)
		FacilityRepo.On("GetFacilityID", uint(1)).Return(entities.Facility{}, errors.New("Error Access Database")).Once()
		FacilityService := NewServiceFacility(FacilityRepo)
		result, err := FacilityService.GetFacilityID(1)
		assert.Error(t, err)
		assert.NotEqual(t, MockFacility[0].Name, result.Name)

		FacilityRepo.AssertExpectations(t)
	})
}

// TEST UPDATE FACILITY SERVICE
func TestUpdateFacility(t *testing.T) {
	var respon = entities.Facility{
		Model:      gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:       "Mall Ciputra",
		Longitude:  -6.168273696181832,
		Latitude:   106.86491520706296,
		DistrictID: 1,
	}
	var NewFacility = entities.UpdateFacility{
		Name:      "Mall Ciputra",
		Longitude: -6.168273696181832,
		Latitude:  106.86491520706296,
	}
	t.Run("Success Update Facility", func(t *testing.T) {

		FacilityRepo := mocks.NewRepoFacility(t)

		FacilityRepo.On("UpdateFacility", uint(1), mock.Anything).Return(respon, nil).Once()

		FacilityService := NewServiceFacility(FacilityRepo)
		result, err := FacilityService.UpdateFacility(uint(1), NewFacility)
		assert.NoError(t, err)
		assert.Equal(t, respon.Name, result.Name)
		FacilityRepo.AssertExpectations(t)
	})
	t.Run("Error Update Facility", func(t *testing.T) {

		FacilityRepo := mocks.NewRepoFacility(t)

		FacilityRepo.On("UpdateFacility", uint(1), mock.Anything).Return(entities.Facility{}, errors.New("Error Access Database")).Once()

		FacilityService := NewServiceFacility(FacilityRepo)
		result, err := FacilityService.UpdateFacility(uint(1), NewFacility)
		assert.Error(t, err)
		assert.NotEqual(t, respon.Name, result.Name)

		FacilityRepo.AssertExpectations(t)
	})
}

// TEST GET FACILITY BY ID SERVICES
func TestDeleteFacility(t *testing.T) {
	t.Run("Success Delete Facility ID", func(t *testing.T) {
		FacilityRepo := mocks.NewRepoFacility(t)
		FacilityRepo.On("DeleteFacility", uint(1)).Return(nil).Once()
		FacilityService := NewServiceFacility(FacilityRepo)
		err := FacilityService.DeleteFacility(1)
		assert.NoError(t, err)

	})
	t.Run("Error Access Database", func(t *testing.T) {
		FacilityRepo := mocks.NewRepoFacility(t)
		FacilityRepo.On("DeleteFacility", uint(1)).Return(errors.New("Error Access Database")).Once()
		FacilityService := NewServiceFacility(FacilityRepo)
		err := FacilityService.DeleteFacility(1)
		assert.Error(t, err)
	})
}

//
func TestGetNearFacility(t *testing.T) {
	result := []entities.NearFacility{{
		Name:   "Mall Ciputra",
		Radius: 100,
	}}
	t.Run("Success Get Near Facility", func(t *testing.T) {
		FacilityRepo := mocks.NewRepoFacility(t)
		FacilityRepo.On("GetNearFacility", uint(1)).Return(result, nil).Once()
		FacilityService := NewServiceFacility(FacilityRepo)
		respon, err := FacilityService.GetNearFacility(uint(1))
		assert.NoError(t, err)
		assert.Equal(t, respon[0].Name, result[0].Name)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		FacilityRepo := mocks.NewRepoFacility(t)
		FacilityRepo.On("GetNearFacility", uint(1)).Return(nil, errors.New("Error Access Database")).Once()
		FacilityService := NewServiceFacility(FacilityRepo)
		_, err := FacilityService.GetNearFacility(uint(1))
		assert.Error(t, err)
	})
}

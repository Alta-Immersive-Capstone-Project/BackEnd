package house

import (
	"errors"
	"kost/entities"
	mocks "kost/mocks/repositories/house"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// Create Mock Data
var MockHouse = []entities.House{
	{
		Model:      gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		DistrictID: 1,
		Title:      "Kost A Maguwoharjo Ring road Utara",
		Brief:      "Brief information about the house",
		OwnerName:  "Adi Yudho",
		OwnerPhone: "+62823987654321",
		Address:    "Jl. Maguwoharjo RT.11 RW.11",
		SlotRoom:   6,
		Available:  3,
		Longitude:  -6.168273696181832,
		Latitude:   106.86491520706296,
	},
	{
		Model:      gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		DistrictID: 1,
		Title:      "Kost A12 Kedaman Maguwo",
		Brief:      "Brief information about the house",
		OwnerName:  "Ana Mana",
		OwnerPhone: "+6282398761111",
		Address:    "Jl. Maguwoharjo RT.12 RW.11",
		SlotRoom:   7,
		Available:  2,
		Longitude:  -6.168273696181832,
		Latitude:   106.86491520706296,
	},
}

func TestCreateHouse(t *testing.T) {
	var NewHouse = entities.AddHouse{
		DistrictID: 1,
		Title:      "Kost A12 Kedaman Maguwo",
		Brief:      "Brief information about the house",
		OwnerName:  "Ana Mana",
		OwnerPhone: "+6282398761111",
		Address:    "Jl. Maguwoharjo RT.12 RW.11",
		SlotRoom:   7,
		Available:  2,
		Longitude:  -6.168273696181832,
		Latitude:   106.86491520706296,
	}

	t.Run("Success Create House", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("CreateHouse", mock.Anything).Return(MockHouse[1], nil).Once()

		HouseService := NewHouseService(HouseRepo)
		result, err := HouseService.CreateHouse(NewHouse)
		assert.NoError(t, err)
		assert.Equal(t, MockHouse[1].Title, result.Title)
		assert.Equal(t, MockHouse[1].OwnerName, result.OwnerName)

		HouseRepo.AssertExpectations(t)
	})

	t.Run("Error Create House", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("CreateHouse", mock.Anything).Return(entities.House{}, errors.New("error create house")).Once()

		HouseService := NewHouseService(HouseRepo)
		result, err := HouseService.CreateHouse(NewHouse)
		assert.Error(t, err)
		assert.NotEqual(t, MockHouse[1].Title, result.Title)
		assert.NotEqual(t, MockHouse[1].OwnerName, result.OwnerName)

		HouseRepo.AssertExpectations(t)
	})
}

func TestUpdateHouse(t *testing.T) {
	var UpdateHouse = entities.UpdateHouse{
		DistrictID: 1,
		Title:      "Kost Maguwo JBay A12 ",
		Brief:      "Brief information about the house",
		OwnerName:  "Ana Mana",
		OwnerPhone: "+6282398761111",
		Address:    "Jl. Maguwoharjo RT.12 RW.11",
		SlotRoom:   7,
		Available:  1,
		Longitude:  -6.168273696181832,
		Latitude:   106.86491520706296,
	}
	var response = entities.House{
		Model:      gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		DistrictID: 1,
		Title:      "Kost Maguwo JBay A12 ",
		Brief:      "Brief information about the house",
		OwnerName:  "Ana Mana",
		OwnerPhone: "+6282398761111",
		Address:    "Jl. Maguwoharjo RT.12 RW.11",
		SlotRoom:   7,
		Available:  1,
		Longitude:  -6.168273696181832,
		Latitude:   106.86491520706296,
	}

	t.Run("Success Update House", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("UpdateHouse", uint(2), mock.Anything).Return(response, nil).Once()

		HouseService := NewHouseService(HouseRepo)
		result, err := HouseService.UpdateHouse(uint(2), UpdateHouse)
		assert.NoError(t, err)
		assert.Equal(t, response.Title, result.Title)
		assert.Equal(t, response.OwnerName, result.OwnerName)

		HouseRepo.AssertExpectations(t)
	})

	t.Run("Error Update House", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("UpdateHouse", uint(2), mock.Anything).Return(entities.House{}, errors.New("error update house")).Once()

		HouseService := NewHouseService(HouseRepo)
		result, err := HouseService.UpdateHouse(uint(2), UpdateHouse)
		assert.Error(t, err)
		assert.NotEqual(t, response.Title, result.Title)
		assert.NotEqual(t, response.OwnerName, result.OwnerName)

		HouseRepo.AssertExpectations(t)
	})
}

func TestDeleteHouse(t *testing.T) {
	t.Run("Success Delete House", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("DeleteHouse", uint(2)).Return(nil).Once()

		HouseService := NewHouseService(HouseRepo)
		err := HouseService.DeleteHouse(uint(2))
		assert.NoError(t, err)

		HouseRepo.AssertExpectations(t)
	})

	t.Run("Error Delete House", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("DeleteHouse", uint(2)).Return(errors.New("error access database")).Once()

		HouseService := NewHouseService(HouseRepo)
		err := HouseService.DeleteHouse(uint(2))
		assert.Error(t, err)

		HouseRepo.AssertExpectations(t)
	})
}

func TestGetAllHouse(t *testing.T) {
	t.Run("Success Get All Houses", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("GetAllHouse", uint(1)).Return(MockHouse, nil).Once()
		HouseService := NewHouseService(HouseRepo)

		result, err := HouseService.GetAllHouseByDist(uint(1))
		assert.NoError(t, err)
		assert.NotNil(t, result)

		HouseRepo.AssertExpectations(t)
	})
	t.Run("Error Get All Houses", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("GetAllHouse", uint(1)).Return(nil, errors.New("Error Access Database")).Once()
		HouseService := NewHouseService(HouseRepo)

		result, err := HouseService.GetAllHouseByDist(uint(1))
		assert.Error(t, err)
		assert.NotNil(t, result)

		HouseRepo.AssertExpectations(t)
	})
}

func TestGetHouseID(t *testing.T) {
	t.Run("Success Get House ID", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("GetHouseID", uint(1)).Return(MockHouse[0], nil).Once()
		HouseService := NewHouseService(HouseRepo)

		result, err := HouseService.GetHouseID(uint(1))
		assert.NoError(t, err)
		assert.Equal(t, MockHouse[0].OwnerName, result.OwnerName)

		HouseRepo.AssertExpectations(t)
	})
	t.Run("Error Get House ID", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("GetHouseID", uint(1)).Return(entities.House{}, errors.New("error access database")).Once()
		DistrictService := NewHouseService(HouseRepo)

		result, err := DistrictService.GetHouseID(uint(1))
		assert.Error(t, err)
		assert.NotEqual(t, MockHouse[0].OwnerName, result.OwnerName)

		HouseRepo.AssertExpectations(t)
	})
}

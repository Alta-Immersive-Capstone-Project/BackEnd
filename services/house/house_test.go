package house

import (
	"errors"
	"kost/entities"
	mocks "kost/mocks/repositories/house"
	roomMock "kost/mocks/repositories/room"
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

var MockHouseJoin = []entities.HouseResponseJoin{
	{
		ID:         1,
		Title:      "Kost A Maguwoharjo Ring road Utara",
		Brief:      "Brief information about the house",
		OwnerName:  "Adi Yudho",
		OwnerPhone: "+62823987654321",
		Address:    "Jl. Maguwoharjo RT.11 RW.11",
		Available:  3,
		DistrictID: 1,
		District:   "Maguwoharjo",
		RoomID:     1,
		Type:       "Putra",
		Price:      700000,
		Rating:     4.5,
	},
	{
		ID:         2,
		Title:      "Kost A12 Kedaman Maguwo",
		Brief:      "Brief information about the house",
		OwnerName:  "Ana Mana",
		OwnerPhone: "+6282398761111",
		Address:    "Jl. Maguwoharjo RT.12 RW.11",
		Available:  2,
		DistrictID: 1,
		District:   "Maguwoharjo",
		RoomID:     1,
		Type:       "Putri",
		Price:      600000,
		Rating:     4,
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
		RoomRepo := roomMock.NewRoomRepo(t)
		HouseService := NewHouseService(HouseRepo, RoomRepo)
		result, err := HouseService.CreateHouse(NewHouse)
		assert.NoError(t, err)
		assert.Equal(t, MockHouse[1].Title, result.Title)
		assert.Equal(t, MockHouse[1].OwnerName, result.OwnerName)

		HouseRepo.AssertExpectations(t)
	})

	t.Run("Error Create House", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("CreateHouse", mock.Anything).Return(entities.House{}, errors.New("error create house")).Once()
		RoomRepo := roomMock.NewRoomRepo(t)
		HouseService := NewHouseService(HouseRepo, RoomRepo)
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
		RoomRepo := roomMock.NewRoomRepo(t)
		HouseService := NewHouseService(HouseRepo, RoomRepo)
		result, err := HouseService.UpdateHouse(uint(2), UpdateHouse)
		assert.NoError(t, err)
		assert.Equal(t, response.Title, result.Title)
		assert.Equal(t, response.OwnerName, result.OwnerName)

		HouseRepo.AssertExpectations(t)
	})

	t.Run("Error Update House", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("UpdateHouse", uint(2), mock.Anything).Return(entities.House{}, errors.New("error update house")).Once()
		RoomRepo := roomMock.NewRoomRepo(t)
		HouseService := NewHouseService(HouseRepo, RoomRepo)
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
		RoomRepo := roomMock.NewRoomRepo(t)
		HouseService := NewHouseService(HouseRepo, RoomRepo)
		err := HouseService.DeleteHouse(uint(2))
		assert.NoError(t, err)

		HouseRepo.AssertExpectations(t)
	})

	t.Run("Error Delete House", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("DeleteHouse", uint(2)).Return(errors.New("error access database")).Once()
		RoomRepo := roomMock.NewRoomRepo(t)
		HouseService := NewHouseService(HouseRepo, RoomRepo)
		err := HouseService.DeleteHouse(uint(2))
		assert.Error(t, err)

		HouseRepo.AssertExpectations(t)
	})
}

func TestGetAllHouseByDist(t *testing.T) {
	t.Run("Success Get All Houses By Dist", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("GetAllHouseByDist", uint(1)).Return(MockHouse, nil).Once()
		RoomRepo := roomMock.NewRoomRepo(t)
		HouseService := NewHouseService(HouseRepo, RoomRepo)
		result, err := HouseService.GetAllHouseByDistrict(uint(1))
		assert.NoError(t, err)
		assert.NotNil(t, result)

		HouseRepo.AssertExpectations(t)
	})
	t.Run("Error Get All Houses By Dist", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("GetAllHouseByDist", uint(1)).Return(nil, errors.New("Error Access Database")).Once()
		RoomRepo := roomMock.NewRoomRepo(t)
		HouseService := NewHouseService(HouseRepo, RoomRepo)
		result, err := HouseService.GetAllHouseByDistrict(uint(1))
		assert.Error(t, err)
		assert.NotNil(t, result)

		HouseRepo.AssertExpectations(t)
	})
}

func TestGetHouseID(t *testing.T) {
	t.Run("Success Get House ID", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("GetHouseID", uint(1)).Return(MockHouse[0], nil).Once()
		RoomRepo := roomMock.NewRoomRepo(t)
		HouseService := NewHouseService(HouseRepo, RoomRepo)

		result, err := HouseService.GetHouseID(uint(1))
		assert.NoError(t, err)
		assert.Equal(t, MockHouse[0].OwnerName, result.OwnerName)

		HouseRepo.AssertExpectations(t)
	})
	t.Run("Error Get House ID", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("GetHouseID", uint(1)).Return(entities.House{}, errors.New("error access database")).Once()
		RoomRepo := roomMock.NewRoomRepo(t)
		DistrictService := NewHouseService(HouseRepo, RoomRepo)

		result, err := DistrictService.GetHouseID(uint(1))
		assert.Error(t, err)
		assert.NotEqual(t, MockHouse[0].OwnerName, result.OwnerName)

		HouseRepo.AssertExpectations(t)
	})
}

func TestFindAllHouseByDistrict(t *testing.T) {
	t.Run("Success Find All Houses By Dist", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("GetAllHouseByDistrict", uint(1)).Return(MockHouseJoin, nil).Once()
		RoomRepo := roomMock.NewRoomRepo(t)
		HouseService := NewHouseService(HouseRepo, RoomRepo)

		result, err := HouseService.FindAllHouseByDistrict(uint(1))
		assert.NoError(t, err)
		assert.NotNil(t, result)

		HouseRepo.AssertExpectations(t)
	})
	t.Run("Error Get All Houses By Dist", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("GetAllHouseByDistrict", uint(1)).Return(nil, errors.New("Error Access Database")).Once()
		RoomRepo := roomMock.NewRoomRepo(t)
		HouseService := NewHouseService(HouseRepo, RoomRepo)

		result, err := HouseService.FindAllHouseByDistrict(uint(1))
		assert.Error(t, err)
		assert.NotNil(t, result)

		HouseRepo.AssertExpectations(t)
	})
}

func TestFindAllHouseByCities(t *testing.T) {
	t.Run("Success Find All Houses By Cities", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("GetAllHouseByCities", uint(1)).Return(MockHouseJoin, nil).Once()
		RoomRepo := roomMock.NewRoomRepo(t)
		HouseService := NewHouseService(HouseRepo, RoomRepo)

		result, err := HouseService.FindAllHouseByCities(uint(1))
		assert.NoError(t, err)
		assert.NotNil(t, result)

		HouseRepo.AssertExpectations(t)
	})
	t.Run("Error Get All Houses By Cities", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("GetAllHouseByCities", uint(1)).Return(nil, errors.New("Error Access Database")).Once()
		RoomRepo := roomMock.NewRoomRepo(t)
		HouseService := NewHouseService(HouseRepo, RoomRepo)

		result, err := HouseService.FindAllHouseByCities(uint(1))
		assert.Error(t, err)
		assert.NotNil(t, result)

		HouseRepo.AssertExpectations(t)
	})
}

func TestFindAllHouseByCtyAndDst(t *testing.T) {
	t.Run("Success Find All Houses By Cities And District", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("GetAllHouseByDstAndCty", uint(1), uint(2)).Return(MockHouseJoin, nil).Once()
		RoomRepo := roomMock.NewRoomRepo(t)
		HouseService := NewHouseService(HouseRepo, RoomRepo)

		result, err := HouseService.FindAllHouseByCtyAndDst(uint(1), uint(2))
		assert.NoError(t, err)
		assert.NotNil(t, result)

		HouseRepo.AssertExpectations(t)
	})
	t.Run("Error Get All Houses By Cities And Dsitrict", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("GetAllHouseByDstAndCty", uint(1), uint(2)).Return(nil, errors.New("error access database")).Once()
		RoomRepo := roomMock.NewRoomRepo(t)
		HouseService := NewHouseService(HouseRepo, RoomRepo)

		_, err := HouseService.FindAllHouseByCtyAndDst(uint(1), uint(2))
		assert.Error(t, err)

		HouseRepo.AssertExpectations(t)
	})
}

func TestSelectAllHouse(t *testing.T) {
	t.Run("Success Find All Houses By Cities", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("SelectAllHouse", mock.Anything).Return(MockHouseJoin, nil).Once()
		RoomRepo := roomMock.NewRoomRepo(t)
		HouseService := NewHouseService(HouseRepo, RoomRepo)

		result, err := HouseService.SelectAllHouse()
		assert.NoError(t, err)
		assert.NotNil(t, result)

		HouseRepo.AssertExpectations(t)
	})
	t.Run("Error Get All Houses By Cities", func(t *testing.T) {
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("SelectAllHouse", mock.Anything).Return(nil, errors.New("Error Access Database")).Once()
		RoomRepo := roomMock.NewRoomRepo(t)
		HouseService := NewHouseService(HouseRepo, RoomRepo)

		result, err := HouseService.SelectAllHouse()
		assert.Error(t, err)
		assert.Nil(t, result)

		HouseRepo.AssertExpectations(t)
	})
}

func TestFindHouseByTitle(t *testing.T) {
	t.Run("Success Find House By Title", func(t *testing.T) {
		title := "Maguwo"
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("FindHouseByTitle", title).Return(MockHouseJoin, nil).Once()
		RoomRepo := roomMock.NewRoomRepo(t)
		HouseService := NewHouseService(HouseRepo, RoomRepo)

		result, err := HouseService.FindHouseByTitle(title)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		HouseRepo.AssertExpectations(t)
	})
	t.Run("Error Get All Houses By Cities", func(t *testing.T) {
		title := "Maguwo"
		HouseRepo := mocks.NewIRepoHouse(t)
		HouseRepo.On("FindHouseByTitle", title).Return([]entities.HouseResponseJoin{}, errors.New("Error Access Database")).Once()
		RoomRepo := roomMock.NewRoomRepo(t)
		HouseService := NewHouseService(HouseRepo, RoomRepo)

		result, err := HouseService.FindHouseByTitle(title)
		assert.Error(t, err)
		assert.Nil(t, result)

		HouseRepo.AssertExpectations(t)
	})
}

package room_test

import (
	"errors"
	"fmt"
	"kost/entities"
	repoMock "kost/mocks/repositories/room"
	room "kost/services/room"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateRoom(t *testing.T) {
	repo := repoMock.NewRoomRepo(t)
	insertData := entities.AddRoom{
		HouseID:                1,
		Type:                   "4x3",
		Price:                  2000000,
		Additional_description: "free wifi",
	}
	returnData := entities.Room{
		HouseID:                1,
		Type:                   "4x3",
		UserID:                 1,
		Price:                  2000000,
		Additional_description: "free wifi",
	}

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("CreateRoom", mock.Anything).Return(returnData, nil).Once()
		srv := room.NewServiceRoom(repo)

		res, err := srv.CreateRoom(1, insertData)
		fmt.Println(res)
		assert.NoError(t, err)
		assert.Equal(t, returnData.ID, res.ID)
		assert.Equal(t, returnData.HouseID, res.HouseID)
		assert.Equal(t, returnData.Type, res.Type)
		assert.Equal(t, returnData.Price, res.Price)
		assert.Equal(t, returnData.Additional_description, res.Additional_description)
		repo.AssertExpectations(t)
	})

	t.Run("Error Insert", func(t *testing.T) {
		repo.On("CreateRoom", mock.Anything).Return(entities.Room{}, errors.New("there is some error")).Once()
		srv := room.NewServiceRoom(repo)

		_, err := srv.CreateRoom(1, insertData)
		assert.Error(t, err)
		assert.EqualError(t, err, "there is some error")
		repo.AssertExpectations(t)
	})
}
func TestGetAllRoom(t *testing.T) {
	repo := repoMock.NewRoomRepo(t)
	returnData := []entities.Room{{
		HouseID:                1,
		Type:                   "4x3",
		UserID:                 1,
		Price:                  2000000,
		Additional_description: "free wifi",
	}}

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("GetAllRoom", uint(1)).Return(returnData, nil).Once()
		srv := room.NewServiceRoom(repo)

		res, err := srv.GetAllRoom(uint(1))
		assert.NoError(t, err)
		assert.Equal(t, returnData[0].ID, res[0].ID)
		assert.Equal(t, returnData[0].HouseID, res[0].HouseID)
		assert.Equal(t, returnData[0].Type, res[0].Type)
		assert.Equal(t, returnData[0].Price, res[0].Price)
		assert.Equal(t, returnData[0].Additional_description, res[0].Additional_description)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get All", func(t *testing.T) {
		repo.On("GetAllRoom", uint(1)).Return(nil, errors.New("there is some error")).Once()
		srv := room.NewServiceRoom(repo)

		_, err := srv.GetAllRoom(uint(1))
		assert.Error(t, err)
		assert.EqualError(t, err, "there is some error")
		repo.AssertExpectations(t)
	})
}
func TestGetIDRoom(t *testing.T) {
	repo := repoMock.NewRoomRepo(t)
	returnData := entities.Room{
		HouseID:                1,
		Type:                   "4x3",
		UserID:                 1,
		Price:                  2000000,
		Additional_description: "free wifi",
	}

	t.Run("Success GetIDRoom", func(t *testing.T) {
		repo.On("GetRoomID", uint(1)).Return(returnData, nil).Once()
		srv := room.NewServiceRoom(repo)

		res, err := srv.GetIDRoom(uint(1))
		assert.NoError(t, err)
		assert.Equal(t, returnData.ID, res.ID)
		assert.Equal(t, returnData.HouseID, res.HouseID)
		assert.Equal(t, returnData.UserID, res.UserID)
		assert.Equal(t, returnData.Type, res.Type)
		assert.Equal(t, returnData.Price, res.Price)
		assert.Equal(t, returnData.Additional_description, res.Additional_description)
		repo.AssertExpectations(t)
	})

	t.Run("Error GetIDRoom", func(t *testing.T) {
		repo.On("GetRoomID", uint(1)).Return(entities.Room{}, errors.New("there is some error")).Once()
		srv := room.NewServiceRoom(repo)

		_, err := srv.GetIDRoom(uint(1))
		assert.Error(t, err)
		assert.EqualError(t, err, "there is some error")
		repo.AssertExpectations(t)
	})
}
func TestUpdateRoom(t *testing.T) {
	repo := repoMock.NewRoomRepo(t)
	insertData := entities.UpdateRoom{
		Type:                   "4x3",
		Price:                  2000000,
		Additional_description: "free wifi",
	}
	returnData := entities.Room{
		HouseID:                1,
		Type:                   "4x3",
		UserID:                 1,
		Price:                  2000000,
		Additional_description: "free wifi",
	}

	t.Run("Success UpdateRoom", func(t *testing.T) {
		repo.On("UpdateRoom", uint(1), mock.Anything).Return(returnData, nil).Once()
		srv := room.NewServiceRoom(repo)

		res, err := srv.UpdateRoom(uint(1), insertData)
		assert.NoError(t, err)
		assert.Equal(t, returnData.ID, res.ID)
		assert.Equal(t, returnData.HouseID, res.HouseID)
		assert.Equal(t, returnData.Type, res.Type)
		assert.Equal(t, returnData.Price, res.Price)
		assert.Equal(t, returnData.Additional_description, res.Additional_description)
		repo.AssertExpectations(t)
	})

	t.Run("Error UpdateRoom", func(t *testing.T) {
		repo.On("UpdateRoom", uint(1), mock.Anything).Return(entities.Room{}, errors.New("there is some error")).Once()
		srv := room.NewServiceRoom(repo)

		_, err := srv.UpdateRoom(uint(1), insertData)
		assert.Error(t, err)
		assert.EqualError(t, err, "there is some error")
		repo.AssertExpectations(t)
	})
}
func TestDeleteRoom(t *testing.T) {
	repo := repoMock.NewRoomRepo(t)

	t.Run("Success DeleteRoom", func(t *testing.T) {
		repo.On("DeleteRoom", mock.Anything).Return(nil).Once()
		srv := room.NewServiceRoom(repo)

		err := srv.DeleteRoom(1)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Error DeleteRoom", func(t *testing.T) {
		repo.On("DeleteRoom", mock.Anything).Return(errors.New("there is some error")).Once()
		srv := room.NewServiceRoom(repo)

		err := srv.DeleteRoom(1)
		assert.Error(t, err)
		assert.EqualError(t, err, "there is some error")
		repo.AssertExpectations(t)
	})
}

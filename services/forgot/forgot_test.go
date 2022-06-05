package forgot

import (
	"errors"
	"kost/entities"
	mocks "kost/mocks/repositories/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindByUser(t *testing.T) {
	respond := entities.User{Name: "Trial", Email: "test@example.com"}
	t.Run("Success Get Data", func(t *testing.T) {
		Repo := mocks.NewUserRepositoryInterface(t)
		Repo.On("FindByUser", mock.Anything).Return(respond, nil).Once()
		ForgetService := NewforgotService(Repo)

		result, err := ForgetService.FindUserByEmail("test@example.com")
		assert.NoError(t, err)
		assert.Equal(t, respond.Name, result.Name)

		Repo.AssertExpectations(t)
	})
	t.Run("Error Get Data", func(t *testing.T) {
		Repo := mocks.NewUserRepositoryInterface(t)
		Repo.On("FindByUser", mock.Anything).Return(respond, errors.New("Error FindByUser")).Once()
		ForgetService := NewforgotService(Repo)

		result, err := ForgetService.FindUserByEmail("test@example.com")
		assert.Error(t, err)
		assert.NotEqual(t, respond.Name, result.Name)

		Repo.AssertExpectations(t)
	})
}

func TestResetPassword(t *testing.T) {
	password := entities.User{Password: "password"}
	respond := entities.User{Name: "Trial", Email: "test@example.com", Password: "password"}
	t.Run("Success Reset", func(t *testing.T) {
		Repo := mocks.NewUserRepositoryInterface(t)
		Repo.On("UpdateUser", uint(1), password).Return(respond, nil).Once()
		ForgetService := NewforgotService(Repo)

		result, err := ForgetService.ResetPassword(1, "password")
		assert.NoError(t, err)
		assert.Equal(t, respond.Name, result.Name)

		Repo.AssertExpectations(t)
	})

	t.Run("Error Reset Password", func(t *testing.T) {
		Repo := mocks.NewUserRepositoryInterface(t)
		Repo.On("UpdateUser", uint(1), password).Return(respond, errors.New("Error Reset Password"))
		ForgetService := NewforgotService(Repo)

		_, err := ForgetService.ResetPassword(1, "password")
		assert.Error(t, err)
		assert.Equal(t, "Error Reset Password", err.Error())

		Repo.AssertExpectations(t)
	})
}

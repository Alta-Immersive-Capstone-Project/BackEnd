package user_test

import (
	"errors"
	"kost/entities"
	repo "kost/mocks/repositories/user"

	"kost/services/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var MockUser = []entities.User{
	{
		Email:    "kiki@gmail.com",
		Password: "#$%$$$!!#@",
		Name:     "kiki",
		Gender:   "girl",
		Phone:    "0811345456",
		Avatar:   "https://belajar-be.s3.ap-southeast-1.amazonaws.com/Avatar/1653973235.png",
		Role:     "admin",
	},
	{
		Email:    "boy@gmail.com",
		Password: "#$%$$9876@",
		Name:     "boy",
		Gender:   "man",
		Phone:    "08117855450",
		Avatar:   "https://belajar-be.s3.ap-southeast-1.amazonaws.com/Avatar/1653973235.png",
		Role:     "customer",
	},
	{
		Email:    "siti@gmail.com",
		Password: "#$786659876@",
		Name:     "siti",
		Gender:   "girl",
		Phone:    "081179743450",
		Avatar:   "a picture",
		Role:     "consultant",
	},
}

func TestCreateUser(t *testing.T) {
	var NewUser = entities.CreateUserRequest{
		Email:    "kiki@gmail.com",
		Password: "#$%$$$!!#@",
		Name:     "kiki",
		Gender:   "girl",
		Phone:    "0811345456",
		Role:     "admin",
	}

	Url := "https://belajar-be.s3.ap-southeast-1.amazonaws.com/Avatar/1653973235.png"

	t.Run("Success", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("InsertUser", mock.Anything).Return(MockUser[0], nil).Once()

		UserService := user.NewUserService(userRepo)
		res, err := UserService.CreateUser(NewUser, Url)
		assert.NoError(t, err)
		assert.NotEqual(t, "", res.Token)
		assert.Equal(t, MockUser[0].Email, res.User.Email)
		assert.Equal(t, MockUser[0].Name, res.User.Name)
		userRepo.AssertExpectations(t)

	})

	t.Run("Error Access Database", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("InsertUser", mock.Anything).Return(entities.User{}, errors.New("Error Access Database")).Once()

		UserService := user.NewUserService(userRepo)
		res, err := UserService.CreateUser(NewUser, Url)
		assert.Error(t, err)
		assert.NotEqual(t, MockUser[0].Name, res.User.Name)
		assert.NotEqual(t, MockUser[0].Email, res.User.Email)
		assert.Equal(t, "", res.Token)

		userRepo.AssertExpectations(t)

	})
}

func TestGetByID(t *testing.T) {

	t.Run("Success", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("GetUserID", mock.Anything).Return(MockUser[1], nil).Once()

		UserService := user.NewUserService(userRepo)

		res, err := UserService.GetbyID(1)

		assert.NoError(t, err)
		assert.Equal(t, MockUser[1].Name, res.Name)
		assert.Equal(t, MockUser[1].Role, "customer")
		userRepo.AssertExpectations(t)
	})

	t.Run("Error Access Database", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("GetUserID", mock.Anything).Return(entities.User{}, errors.New("error access database"))

		srv := user.NewUserService(userRepo)

		res, err := srv.GetbyID(1)
		assert.Error(t, err)
		assert.NotEqual(t, MockUser[0].Name, res.Name)
		userRepo.AssertExpectations(t)
	})
}

func TestUpdateInternal(t *testing.T) {
	userRepo := repo.NewUserRepositoryInterface(t)
	var respon = entities.InternalResponse{
		Email:  "kikifatmala@gmail.com",
		Name:   "kiki Saputri",
		Gender: "girl",
		Avatar: "a picture",
		Phone:  "08117855450",
		Role:   "admin",
	}
	var internalRequest = entities.UpdateInternalRequest{
		Name: "kiki Saputri",
	}
	var userInternalRequest = entities.User{
		Name:     "kiki Saputri",
		Email:    "kiki@gmail.com",
		Password: "#$%$$$!!#@",
		Gender:   "girl",
		Phone:    "0811345456",
		Avatar:   "https://belajar-be.s3.ap-southeast-1.amazonaws.com/Avatar/1653973235.png",
		Role:     "admin",
	}
	var userData = entities.User{
		Email:    "kikifatmala@gmail.com",
		Password: "$2a$04$aFJ./S730/7TWneKFiS0ruCHNs9g97yumrB5RNx53gRqTDThpeQLa",
		Name:     "kiki Saputri",
		Gender:   "girl",
		Avatar:   "a picture",
		Phone:    "08117855450",
		Role:     "admin",
	}
	Url := "https://belajar-be.s3.ap-southeast-1.amazonaws.com/Avatar/1653973235.png"
	t.Run("Success", func(t *testing.T) {

		userRepo.On("GetUserID", uint(1)).Return(MockUser[0], nil)
		userRepo.On("UpdateUser", uint(1), userInternalRequest).Return(userData, nil).Once()
		UserService := user.NewUserService(userRepo)

		res, err := UserService.UpdateInternal(internalRequest, 1, Url)

		assert.NoError(t, err)
		assert.Equal(t, respon.Name, res.Name)
		assert.Equal(t, respon.Email, res.Email)
		userRepo.AssertExpectations(t)
	})

	t.Run("Error Update User", func(t *testing.T) {

		userRepo.On("GetUserID", uint(1)).Return(MockUser[0], nil)
		userRepo.On("UpdateUser", uint(1), userInternalRequest).Return(entities.User{}, errors.New("Error Update Data")).Once()

		UserService := user.NewUserService(userRepo)

		res, err := UserService.UpdateInternal(internalRequest, 1, Url)

		assert.Error(t, err)
		assert.NotEqual(t, respon.Role, res.Name)
		assert.NotEqual(t, respon.Gender, res.Email)
		userRepo.AssertExpectations(t)
	})
}

func TestUpdateCustomer(t *testing.T) {
	var respon = entities.CustomerResponse{

		Email: "boy@gmail.com",

		Name:   "Boy William",
		Gender: "man",
		Phone:  "08117855450",
		Avatar: "https://belajar-be.s3.ap-southeast-1.amazonaws.com/Avatar/1653973235.png",
	}
	var customerRequest = entities.UpdateCustomerRequest{
		Name: "Boy William",
	}
	var userCustomerRequest = entities.User{
		Name:     "Boy William",
		Email:    "boy@gmail.com",
		Password: "#$%$$9876@",

		Gender: "man",
		Phone:  "08117855450",
		Avatar: "https://belajar-be.s3.ap-southeast-1.amazonaws.com/Avatar/1653973235.png",
		Role:   "customer",
	}
	var userData = entities.User{
		Email:    "boy@gmail.com",
		Password: "#$%$$9876@",
		Name:     "Boy William",
		Gender:   "man",
		Phone:    "08117855450",
		Avatar:   "https://belajar-be.s3.ap-southeast-1.amazonaws.com/Avatar/1653973235.png",
		Role:     "customer",
	}
	Url := "https://belajar-be.s3.ap-southeast-1.amazonaws.com/Avatar/1653973235.png"

	t.Run("Success", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("GetUserID", uint(2)).Return(MockUser[1], nil)
		userRepo.On("UpdateUser", uint(2), userCustomerRequest).Return(userData, nil).Once()
		UserService := user.NewUserService(userRepo)

		res, err := UserService.UpdateCustomer(customerRequest, 2, Url)

		assert.NoError(t, err)
		assert.Equal(t, respon.Name, res.Name)
		assert.Equal(t, respon.Email, res.Email)
		userRepo.AssertExpectations(t)
	})

	t.Run("Error Update User", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("GetUserID", uint(2)).Return(MockUser[1], nil)
		userRepo.On("UpdateUser", uint(2), userCustomerRequest).Return(entities.User{}, errors.New("Error Update Data")).Once()
		UserService := user.NewUserService(userRepo)

		res, err := UserService.UpdateCustomer(customerRequest, 2, Url)

		assert.Error(t, err)
		assert.NotEqual(t, respon.Gender, res.Name)
		assert.NotEqual(t, respon.Phone, res.Email)
		userRepo.AssertExpectations(t)
	})
}

func TestDeleteInternal(t *testing.T) {

	t.Run("Success Delete Data", func(t *testing.T) {
		userR := repo.NewUserRepositoryInterface(t)

		userR.On("DeleteUser", uint(2)).Return(nil)

		srvDelete := user.NewUserService(userR)

		err := srvDelete.DeleteInternal(2)

		assert.NoError(t, err)

		userR.AssertExpectations(t)
	})

	t.Run("Error Access Database", func(t *testing.T) {
		userR := repo.NewUserRepositoryInterface(t)

		userR.On("DeleteUser", uint(2)).Return(errors.New("Error Access Database"))
		srvDelete := user.NewUserService(userR)

		err := srvDelete.DeleteInternal(2)

		assert.Error(t, err)

		userR.AssertExpectations(t)
	})
}

func TestDeleteCustomer(t *testing.T) {

	t.Run("Success Delete Data", func(t *testing.T) {
		userR := repo.NewUserRepositoryInterface(t)

		userR.On("DeleteUser", uint(1)).Return(nil)

		srvDelete := user.NewUserService(userR)

		err := srvDelete.DeleteCustomer(1)

		assert.NoError(t, err)

		userR.AssertExpectations(t)
	})

	t.Run("Error Access Database", func(t *testing.T) {
		userR := repo.NewUserRepositoryInterface(t)

		userR.On("DeleteUser", uint(1)).Return(errors.New("error")).Once()
		srvDelete := user.NewUserService(userR)

		err := srvDelete.DeleteCustomer(1)

		assert.Error(t, err)

		userR.AssertExpectations(t)
	})
}

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
	var NewCustomer = entities.CreateUserRequest{
		Email:    "kiki@gmail.com",
		Password: "#$%$$$!!#@",
		Name:     "kiki",
		Gender:   "girl",
		Phone:    "0811345456",
		Role:     "",
	}

	Url := "https://belajar-be.s3.ap-southeast-1.amazonaws.com/Avatar/1653973235.png"

	t.Run("Success Admin", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("InsertUser", mock.Anything).Return(MockUser[0], nil).Once()

		UserService := user.NewUserService(userRepo)
		res, err := UserService.CreateUser(NewUser, Url)
		assert.NoError(t, err)
		assert.Equal(t, MockUser[0].Email, res.Email)
		assert.Equal(t, MockUser[0].Name, res.Name)
		userRepo.AssertExpectations(t)

	})
	t.Run("Success Costumer", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("InsertUser", mock.Anything).Return(MockUser[0], nil).Once()

		UserService := user.NewUserService(userRepo)
		res, err := UserService.CreateUser(NewCustomer, Url)
		assert.NoError(t, err)
		assert.Equal(t, MockUser[0].Email, res.Email)
		assert.Equal(t, MockUser[0].Name, res.Name)
		userRepo.AssertExpectations(t)

	})

	t.Run("Error Access Database", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("InsertUser", mock.Anything).Return(entities.User{}, errors.New("Error Access Database")).Once()

		UserService := user.NewUserService(userRepo)
		_, err := UserService.CreateUser(NewCustomer, Url)
		assert.Error(t, err)

		userRepo.AssertExpectations(t)

	})
}

func TestGetAllMember(t *testing.T) {

	t.Run("Success", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("GetAllUser", mock.Anything).Return(MockUser, nil).Once()

		UserService := user.NewUserService(userRepo)

		res, err := UserService.GetAllMember()

		assert.NoError(t, err)
		assert.Equal(t, MockUser[1].Name, res[1].Name)
		assert.Equal(t, MockUser[1].Role, res[1].Role)
		userRepo.AssertExpectations(t)
	})

	t.Run("Error Access Database", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("GetAllUser", mock.Anything).Return(MockUser, errors.New("error access database")).Once()

		srv := user.NewUserService(userRepo)

		res, err := srv.GetAllMember()
		assert.Error(t, err)
		assert.Len(t, res, 0)
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

	var internalRequestWithPassword = entities.UpdateInternalRequest{
		Name:     "kiki Saputri",
		Password: "susi",
	}

	var userDataWithPassword = entities.User{
		Email:    "kiki@gmail.com",
		Password: "$2a$04$pPQj0kmfwt1LLcN7SPPm4.cUBKBydWKMYMpvXNRjNG6swfrq.H/OS",
		Name:     "kiki Saputri",
		Gender:   "girl",
		Avatar:   "https://belajar-be.s3.ap-southeast-1.amazonaws.com/Avatar/1653973235.png",
		Phone:    "0811345456",
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

		userRepo.On("GetUserID", uint(1)).Return(MockUser[0], nil).Once()
		userRepo.On("UpdateUser", uint(1), mock.Anything).Return(userData, nil).Once()
		UserService := user.NewUserService(userRepo)

		res, err := UserService.UpdateInternal(internalRequest, 1, Url)

		assert.NoError(t, err)
		assert.Equal(t, respon.Name, res.Name)
		assert.Equal(t, respon.Email, res.Email)
		userRepo.AssertExpectations(t)
	})

	t.Run("Success with password", func(t *testing.T) {

		userRepo.On("GetUserID", uint(1)).Return(MockUser[0], nil).Once()
		userRepo.On("UpdateUser", uint(1), mock.Anything).Return(userDataWithPassword, nil).Once()
		UserService := user.NewUserService(userRepo)

		res, err := UserService.UpdateInternal(internalRequestWithPassword, 1, Url)

		assert.NoError(t, err)
		assert.Equal(t, userDataWithPassword.Name, res.Name)
		assert.Equal(t, userDataWithPassword.Email, res.Email)
		userRepo.AssertExpectations(t)
	})

	t.Run("Error Get ID", func(t *testing.T) {
		userRepo.On("GetUserID", uint(1)).Return(entities.User{}, errors.New("Error Get ID")).Once()
		UserService := user.NewUserService(userRepo)

		_, err := UserService.UpdateInternal(internalRequest, 1, Url)

		assert.Error(t, err)
		userRepo.AssertExpectations(t)
	})

	t.Run("Error Update User", func(t *testing.T) {

		userRepo.On("GetUserID", uint(1)).Return(MockUser[0], nil).Once()
		userRepo.On("UpdateUser", uint(1), mock.Anything).Return(entities.User{}, errors.New("Error Update Data")).Once()

		UserService := user.NewUserService(userRepo)

		res, err := UserService.UpdateInternal(internalRequest, 1, Url)

		assert.Error(t, err)
		assert.NotEqual(t, respon.Role, res.Name)
		assert.NotEqual(t, respon.Gender, res.Email)
		userRepo.AssertExpectations(t)
	})
}

func TestUpdateCustomer(t *testing.T) {
	userRepo := repo.NewUserRepositoryInterface(t)
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
	var customerRequestWithPassword = entities.UpdateCustomerRequest{
		Name:     "Boy William",
		Password: "susi",
	}
	var userDataWithPassword = entities.User{
		Email:    "boy@gmail.com",
		Password: "$2a$04$yRAN5vClxfY4abOFiya5/On9g0lrDS.aoowraZVHDtm.DiabsGm8q",
		Name:     "Boy William",
		Gender:   "man",
		Avatar:   "https://belajar-be.s3.ap-southeast-1.amazonaws.com/Avatar/1653973235.png",
		Phone:    "0811345456",
		Role:     "customer",
	}

	var userData = entities.User{
		Email:    "boy@gmail.com",
		Password: "$2a$04$yRAN5vClxfY4abOFiya5/On9g0lrDS.aoowraZVHDtm.DiabsGm8q",
		Name:     "Boy William",
		Gender:   "man",
		Phone:    "08117855450",
		Avatar:   "https://belajar-be.s3.ap-southeast-1.amazonaws.com/Avatar/1653973235.png",
		Role:     "customer",
	}
	Url := "https://belajar-be.s3.ap-southeast-1.amazonaws.com/Avatar/1653973235.png"

	t.Run("Success", func(t *testing.T) {

		userRepo.On("GetUserID", uint(2)).Return(MockUser[1], nil).Once()
		userRepo.On("UpdateUser", uint(2), mock.Anything).Return(userData, nil).Once()
		UserService := user.NewUserService(userRepo)

		res, err := UserService.UpdateCustomer(customerRequest, 2, Url)

		assert.NoError(t, err)
		assert.Equal(t, respon.Name, res.Name)
		assert.Equal(t, respon.Email, res.Email)
		userRepo.AssertExpectations(t)
	})

	t.Run("Success with password", func(t *testing.T) {

		userRepo.On("GetUserID", uint(2)).Return(MockUser[1], nil).Once()
		userRepo.On("UpdateUser", uint(2), mock.Anything).Return(userDataWithPassword, nil).Once()
		UserService := user.NewUserService(userRepo)

		res, err := UserService.UpdateCustomer(customerRequestWithPassword, 2, Url)

		assert.NoError(t, err)
		assert.Equal(t, userDataWithPassword.Name, res.Name)
		assert.Equal(t, userDataWithPassword.Email, res.Email)
		userRepo.AssertExpectations(t)
	})

	t.Run("Error Get ID", func(t *testing.T) {
		userRepo.On("GetUserID", uint(2)).Return(entities.User{}, errors.New("Error Get ID")).Once()

		UserService := user.NewUserService(userRepo)

		_, err := UserService.UpdateCustomer(customerRequest, 2, Url)

		assert.Error(t, err)
		userRepo.AssertExpectations(t)
	})

	t.Run("Error Update User", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("GetUserID", uint(2)).Return(MockUser[1], nil).Once()
		userRepo.On("UpdateUser", uint(2), mock.Anything).Return(entities.User{}, errors.New("Error Update Data")).Once()
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

		userR.On("DeleteUser", uint(2)).Return(nil).Once()

		srvDelete := user.NewUserService(userR)

		err := srvDelete.DeleteInternal(2)

		assert.NoError(t, err)

		userR.AssertExpectations(t)
	})

	t.Run("Error Access Database", func(t *testing.T) {
		userR := repo.NewUserRepositoryInterface(t)

		userR.On("DeleteUser", uint(2)).Return(errors.New("Error Access Database")).Once()
		srvDelete := user.NewUserService(userR)

		err := srvDelete.DeleteInternal(2)

		assert.Error(t, err)

		userR.AssertExpectations(t)
	})
}

func TestDeleteCustomer(t *testing.T) {

	t.Run("Success Delete Data", func(t *testing.T) {
		userR := repo.NewUserRepositoryInterface(t)

		userR.On("DeleteUser", uint(1)).Return(nil).Once()

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

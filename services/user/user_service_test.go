package user_test

import (
	"errors"
	"kost/entities"
	repo "kost/mocks/repositories/user"
	storagee "kost/mocks/services/storage"
	"kost/services/user"
	"mime/multipart"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var MockUser = []entities.User{
	{
		Model:    gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Email:    "kiki@gmail.com",
		Password: "#$%$$$!!#@",
		Name:     "kiki",
		Gender:   "girl",
		Phone:    "0811345456",
		Avatar:   "a picture",
		Role:     "admin",
	},
	{
		Model:    gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Email:    "boy@gmail.com",
		Password: "#$%$$9876@",
		Name:     "boy",
		Gender:   "man",
		Phone:    "08117855450",
		Avatar:   "a picture",
		Role:     "customer",
	},
	{
		Model:    gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
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

	file := map[string]*multipart.FileHeader{
		"avatar": {
			Filename: "avatar.jpg",
			Size:     155 * 1024,
		},
	}

	t.Run("Success", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("InsertUser", mock.Anything).Return(MockUser[0], nil).Once()

		storageService := storagee.NewStorageInterface(t)
		storageService.On("UploadFromRequest", mock.Anything, mock.Anything).Return("example.com/images.png", nil).Once()

		UserService := user.NewUserService(userRepo, validator.New())
		res, err := UserService.CreateUser(NewUser, file, storageService)
		assert.NoError(t, err)
		assert.NotEqual(t, "", res.Token)
		assert.Equal(t, MockUser[0].Email, res.User.Email)
		assert.Equal(t, MockUser[0].Name, res.User.Name)
		userRepo.AssertExpectations(t)
		storageService.AssertExpectations(t)
		storageService.AssertExpectations(t)
	})

	t.Run("Error Access Database", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("InsertUser", mock.Anything).Return(entities.User{}, errors.New("Error Access Database")).Once()

		storageService := storagee.NewStorageInterface(t)
		storageService.On("UploadFromRequest", mock.Anything, mock.Anything).Return("example.com/images.png", nil).Once()

		UserService := user.NewUserService(userRepo, validator.New())
		res, err := UserService.CreateUser(NewUser, file, storageService)
		assert.Error(t, err)
		assert.NotEqual(t, MockUser[0].Name, res.User.Name)
		assert.NotEqual(t, MockUser[0].Email, res.User.Email)
		assert.Equal(t, "", res.Token)

		userRepo.AssertExpectations(t)
		storageService.AssertExpectations(t)
	})
}

func TestGetCustomer(t *testing.T) {

	t.Run("Success", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("GetUserID", mock.Anything).Return(MockUser[1], nil).Once()

		UserService := user.NewUserService(userRepo, validator.New())

		res, err := UserService.GetCustomer(1)

		assert.NoError(t, err)
		assert.Equal(t, MockUser[1].Name, res.Name)
		assert.Equal(t, MockUser[1].Role, "customer")
		userRepo.AssertExpectations(t)
	})

	t.Run("Error Access Database", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("GetUserID", mock.Anything).Return(entities.User{}, errors.New("error access database"))

		srv := user.NewUserService(userRepo, validator.New())

		res, err := srv.GetCustomer(1)
		assert.Error(t, err)
		assert.NotEqual(t, MockUser[0].Name, res.Name)
		userRepo.AssertExpectations(t)
	})
}

func TestUpdateInternal(t *testing.T) {
	var respon = entities.User{
		Model:    gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Email:    "kikifatmala@gmail.com",
		Password: "#$%$$$!!#@",
		Name:     "kiki sumiati",
		Gender:   "girl",
		Phone:    "0811345456",
		Avatar:   "a picture",
		Role:     "admin",
	}
	var Update = entities.UpdateInternalRequest{
		Email:    "kikifatmala@gmail.com",
		Password: "#$%$$$!!#@",
		Name:     "kiki sumiati",
		Gender:   "girl",
		Phone:    "0811345456",
		Role:     "admin",
	}
	file := map[string]*multipart.FileHeader{
		"avatar": {
			Filename: "avatar.jpg",
			Size:     155 * 1024,
		},
	}

	t.Run("Success", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("GetUserID", mock.Anything).Return(MockUser[0], nil)
		userRepo.On("UpdateUser", 1, mock.Anything).Return(respon, nil).Once()

		storageService := storagee.NewStorageInterface(t)
		storageService.On("Delete", mock.Anything).Return(nil).Once()
		storageService.On("UploadFromRequest", mock.Anything, mock.Anything).Return("example.com/images.png", nil).Once()

		UserService := user.NewUserService(userRepo, validator.New())

		res, err := UserService.UpdateInternal(Update, 1, file, storageService)

		assert.NoError(t, err)
		assert.Equal(t, respon.Name, res.Name)
		assert.Equal(t, respon.Email, res.Email)
		userRepo.AssertExpectations(t)
	})

	t.Run("Error Update User", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("GetUserID", mock.Anything).Return(MockUser[0], nil)
		userRepo.On("UpdateUser", 1, mock.Anything).Return(entities.User{}, errors.New("Error Access Database")).Once()

		storageService := storagee.NewStorageInterface(t)
		storageService.On("Delete", mock.Anything).Return(nil).Once()
		storageService.On("UploadFromRequest", mock.Anything, mock.Anything).Return("example.com/images.png", nil).Once()

		UserService := user.NewUserService(userRepo, validator.New())

		res, err := UserService.UpdateInternal(Update, 1, file, storageService)

		assert.Error(t, err)
		assert.NotEqual(t, respon.Role, res.Name)
		assert.NotEqual(t, respon.Gender, res.Email)
		userRepo.AssertExpectations(t)
	})
}

func TestUpdateCustomer(t *testing.T) {
	var responMethod = entities.User{
		Model:    gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Email:    "boyhandsome@gmail.com",
		Password: "#$%$$9876@",
		Name:     "boy permana",
		Gender:   "man",
		Phone:    "08117855450",
		Avatar:   "a picture",
		Role:     "customer",
	}
	var UpdateCustomer = entities.UpdateCustomerRequest{
		Email:    "boyhandsome@gmail.com",
		Password: "#$%$$9876@",
		Name:     "boy permana",
		Gender:   "man",
		Phone:    "08117855450",
	}
	file := map[string]*multipart.FileHeader{
		"avatar": {
			Filename: "avatar.jpg",
			Size:     155 * 1024,
		},
	}

	t.Run("Success", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("GetUserID", mock.Anything).Return(MockUser[1], nil)
		userRepo.On("UpdateUser", 2, mock.Anything).Return(responMethod, nil).Once()

		storageService := storagee.NewStorageInterface(t)
		storageService.On("Delete", mock.Anything).Return(nil).Once()
		storageService.On("UploadFromRequest", mock.Anything, mock.Anything).Return("example.com/images.png", nil).Once()

		UserService := user.NewUserService(userRepo, validator.New())

		res, err := UserService.UpdateCustomer(UpdateCustomer, 2, file, storageService)

		assert.NoError(t, err)
		assert.Equal(t, responMethod.Name, res.Name)
		assert.Equal(t, responMethod.Email, res.Email)
		userRepo.AssertExpectations(t)
	})

	t.Run("Error Update User", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("GetUserID", mock.Anything).Return(MockUser[1], nil)
		userRepo.On("UpdateUser", 2, mock.Anything).Return(entities.User{}, errors.New("Error Access Database")).Once()

		storageService := storagee.NewStorageInterface(t)
		storageService.On("Delete", mock.Anything).Return(nil).Once()
		storageService.On("UploadFromRequest", mock.Anything, mock.Anything).Return("example.com/images.png", nil).Once()

		UserService := user.NewUserService(userRepo, validator.New())

		res, err := UserService.UpdateCustomer(UpdateCustomer, 2, file, storageService)

		assert.Error(t, err)
		assert.NotEqual(t, responMethod.Gender, res.Name)
		assert.NotEqual(t, responMethod.Phone, res.Email)
		userRepo.AssertExpectations(t)
	})
}

func TestGetInternal(t *testing.T) {

	t.Run("Success", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("GetUserID", mock.Anything).Return(MockUser[0], nil).Once()

		UserService := user.NewUserService(userRepo, validator.New())

		res, err := UserService.GetInternal(1)

		assert.NoError(t, err)
		assert.Equal(t, MockUser[0].Name, res.Name)
		assert.Equal(t, MockUser[0].Role, "admin")
		userRepo.AssertExpectations(t)
	})

	t.Run("Error Access Database", func(t *testing.T) {

		userRepo := repo.NewUserRepositoryInterface(t)
		userRepo.On("GetUserID", mock.Anything).Return(entities.User{}, errors.New("error access database"))

		srv := user.NewUserService(userRepo, validator.New())

		res, err := srv.GetInternal(1)
		assert.Error(t, err)
		assert.NotEqual(t, MockUser[0].Name, res.Name)
		assert.NotEqual(t, MockUser[0].Role, "customer")
		userRepo.AssertExpectations(t)
	})
}

func TestDeleteInternal(t *testing.T) {

	t.Run("Success Delete Data", func(t *testing.T) {
		userR := repo.NewUserRepositoryInterface(t)
		userR.On("GetUserID", mock.Anything).Return(MockUser[0], nil).Once()

		srvDelete := user.NewUserService(userR, validator.New())

		storageDelete := storagee.NewStorageInterface(t)
		storageDelete.On("Delete", mock.Anything).Return(nil).Once()
		userR.On("DeleteUser", mock.Anything).Return(nil).Once()

		err := srvDelete.DeleteInternal(1, storageDelete)

		assert.NoError(t, err)

		userR.AssertExpectations(t)
	})

	t.Run("Error Access Database", func(t *testing.T) {
		userR := repo.NewUserRepositoryInterface(t)
		userR.On("GetUserID", mock.Anything).Return(MockUser[0], nil).Once()

		srvDelete := user.NewUserService(userR, validator.New())

		storageDelete := storagee.NewStorageInterface(t)
		storageDelete.On("Delete", mock.Anything).Return(nil).Once()
		userR.On("DeleteUser", mock.Anything).Return(errors.New("Error Access Database")).Once()

		err := srvDelete.DeleteInternal(1, storageDelete)

		assert.Error(t, err)

		userR.AssertExpectations(t)
	})
}

func TestDeleteCustomer(t *testing.T) {

	t.Run("Success Delete Data", func(t *testing.T) {
		userR := repo.NewUserRepositoryInterface(t)
		userR.On("GetUserID", mock.Anything).Return(MockUser[0], nil).Once()

		srvDelete := user.NewUserService(userR, validator.New())

		storageDelete := storagee.NewStorageInterface(t)
		storageDelete.On("Delete", mock.Anything).Return(nil).Once()
		userR.On("DeleteUser", mock.Anything).Return(nil).Once()

		err := srvDelete.DeleteInternal(1, storageDelete)

		assert.NoError(t, err)

		userR.AssertExpectations(t)
	})

	t.Run("Error Access Database", func(t *testing.T) {
		userR := repo.NewUserRepositoryInterface(t)
		userR.On("GetUserID", mock.Anything).Return(MockUser[0], nil).Once()

		srvDelete := user.NewUserService(userR, validator.New())

		storageDelete := storagee.NewStorageInterface(t)
		storageDelete.On("Delete", mock.Anything).Return(nil).Once()
		userR.On("DeleteUser", mock.Anything).Return(errors.New("Error Access Database")).Once()

		err := srvDelete.DeleteInternal(1, storageDelete)

		assert.Error(t, err)

		userR.AssertExpectations(t)
	})
}

package user

import (
	"kost/deliveries/helpers"
	"kost/deliveries/middlewares"
	"kost/entities"
	userRepository "kost/repositories/user"
	storageProvider "kost/services/storage"
	"mime/multipart"
	"net/url"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type UserService struct {
	userRepo userRepository.UserRepositoryInterface
	validate *validator.Validate
}

func NewUserService(repository userRepository.UserRepositoryInterface, valid *validator.Validate) *UserService {
	return &UserService{
		userRepo: repository,
		validate: valid,
	}
}

func (us *UserService) CreateUser(internalRequest entities.CreateUserRequest, files map[string]*multipart.FileHeader, storageProvider storageProvider.StorageInterface) (entities.InternalAuthResponse, error) {

	// Konversi user request menjadi domain untuk diteruskan ke repository
	user := entities.User{}
	copier.Copy(&user, &internalRequest)

	// Password hashing menggunakan bcrypt
	hashedPassword, _ := helpers.HashPassword(user.Password)
	user.Password = hashedPassword

	// Upload avatar if exists
	for field, file := range files {
		switch field {
		case "avatar":
			filename := uuid.New().String() + file.Filename
			fileURL, err := storageProvider.UploadFromRequest("users/avatar/"+filename, file)
			if err != nil {
				return entities.InternalAuthResponse{}, err
			}
			user.Avatar = fileURL
		}
	}

	if user.Role == "" {
		user.Role = "customer"
	}
	// Insert ke sistem melewati repository
	user, err := us.userRepo.InsertUser(user)
	if err != nil {
		return entities.InternalAuthResponse{}, err
	}

	// Konversi hasil repository menjadi user response
	userRes := entities.InternalResponse{}
	copier.Copy(&userRes, &user)

	// generate token
	token, err := middlewares.CreateToken(int(user.ID), user.Name, user.Role)
	if err != nil {
		return entities.InternalAuthResponse{}, err
	}

	// Buat auth response untuk dimasukkan token dan user
	authRes := entities.InternalAuthResponse{
		Token: token,
		User:  userRes,
	}
	return authRes, nil
}

func (us *UserService) GetCustomer(id int) (entities.CustomerResponse, error) {

	user, err := us.userRepo.GetUserID(id)
	if err != nil {
		return entities.CustomerResponse{}, err
	} else if user.Role != "customer" {
		return entities.CustomerResponse{}, err
	}
	userRes := entities.CustomerResponse{}
	copier.Copy(&userRes, &user)

	return userRes, err
}
func (us *UserService) GetInternal(id int) (entities.InternalResponse, error) {

	user, err := us.userRepo.GetUserID(id)
	if err != nil {
		return entities.InternalResponse{}, err
	} else if user.Role == "customer" {
		return entities.InternalResponse{}, err
	}
	userRes := entities.InternalResponse{}
	copier.Copy(&userRes, &user)

	return userRes, err
}

func (us *UserService) GetAllMember() ([]entities.User, error) {

	user, err := us.userRepo.GetAllUser()
	if err != nil {
		return []entities.User{}, err
	} else if user[0].Role == "customer" {
		return []entities.User{}, err
	}
	// userRes := []entities.InternalResponse{}

	return user, err
}

func (us *UserService) UpdateInternal(internalRequest entities.UpdateInternalRequest, id int, files map[string]*multipart.FileHeader, storageProvider storageProvider.StorageInterface) (entities.InternalResponse, error) {

	// Get user by ID via repository
	user, err := us.userRepo.GetUserID(id)
	if err != nil {
		return entities.InternalResponse{}, err
	}
	// Avatar
	for field, file := range files {
		switch field {
		case "avatar":
			// Delete avatar lama jika ada yang baru
			if user.Avatar != "" {
				u, _ := url.Parse(user.Avatar)
				objectPathS3 := strings.TrimPrefix(u.Path, "/")
				storageProvider.Delete(objectPathS3)
			}

			// Upload file to S3
			filename := uuid.New().String() + file.Filename
			fileURL, err := storageProvider.UploadFromRequest("users/avatar/"+filename, file)
			if err != nil {
				return entities.InternalResponse{}, err
			}
			user.Avatar = fileURL
		}
	}

	// Konversi dari request ke domain entities user - mengabaikan nilai kosong pada request
	copier.CopyWithOption(&user, &internalRequest, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	// Hanya hash password jika password juga diganti (tidak kosong)
	if internalRequest.Password != "" {
		hashedPassword, _ := helpers.HashPassword(user.Password)
		user.Password = hashedPassword
	}

	// Update via repository
	user, err = us.userRepo.UpdateUser(id, user)
	// Konversi user domain menjadi user response
	userRes := entities.InternalResponse{}
	copier.Copy(&userRes, &user)

	return userRes, err
}

func (us *UserService) UpdateCustomer(customerRequest entities.UpdateCustomerRequest, id int, files map[string]*multipart.FileHeader, storage storageProvider.StorageInterface) (entities.CustomerResponse, error) {

	// Get user by ID via repository
	user, err := us.userRepo.GetUserID(id)
	if err != nil || user.Role != "customer" {
		return entities.CustomerResponse{}, err
	}
	// Avatar
	for field, file := range files {
		switch field {
		case "avatar":
			// Delete avatar lama jika ada yang baru
			if user.Avatar != "" {
				u, _ := url.Parse(user.Avatar)
				objectPathS3 := strings.TrimPrefix(u.Path, "/")
				storage.Delete(objectPathS3)
			}

			// Upload file to S3
			filename := uuid.New().String() + file.Filename
			fileURL, err := storage.UploadFromRequest("users/avatar/"+filename, file)
			if err != nil {
				return entities.CustomerResponse{}, err
			}
			user.Avatar = fileURL
		}
	}

	// Konversi dari request ke domain entities user - mengabaikan nilai kosong pada request
	copier.CopyWithOption(&user, &customerRequest, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	// Hanya hash password jika password juga diganti (tidak kosong)
	if customerRequest.Password != "" {
		hashedPassword, _ := helpers.HashPassword(user.Password)
		user.Password = hashedPassword
	}

	// Update via repository
	user, err = us.userRepo.UpdateUser(id, user)
	// Konversi user domain menjadi user response
	userRes := entities.CustomerResponse{}
	copier.Copy(&userRes, &user)

	return userRes, err
}

func (us *UserService) DeleteInternal(id int, storage storageProvider.StorageInterface) error {

	// Cari user berdasarkan ID via repo
	user, err := us.userRepo.GetUserID(id)
	if err != nil || user.Role == "customer" {
		return err
	}

	// Delete avatar lama jika ada yang baru
	if user.Avatar != "" {
		u, _ := url.Parse(user.Avatar)
		objectPathS3 := strings.TrimPrefix(u.Path, "/")
		storage.Delete(objectPathS3)
	}
	// Delete user order

	// Delete via repository
	err = us.userRepo.DeleteUser(id)
	return err
}

func (us *UserService) DeleteCustomer(id int, storage storageProvider.StorageInterface) error {

	// Cari user berdasarkan ID via repo
	user, err := us.userRepo.GetUserID(id)
	if err != nil || user.Role != "customer" {
		return err
	}

	// Delete avatar lama jika ada yang baru
	if user.Avatar != "" {
		u, _ := url.Parse(user.Avatar)
		objectPathS3 := strings.TrimPrefix(u.Path, "/")
		storage.Delete(objectPathS3)
	}
	// Delete user order

	// Delete via repository
	err = us.userRepo.DeleteUser(id)
	return err
}

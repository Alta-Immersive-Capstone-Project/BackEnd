package user

import (
	"kost/deliveries/helpers"
	_middleware "kost/deliveries/middlewares"
	"kost/deliveries/validations"
	"kost/entities"
	web "kost/entities/web"
	userRepository "kost/repositories/user"
	storageProvider "kost/services/storage"
	"mime/multipart"

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

func (us *UserService) CreateInternal(internalRequest entities.CreateInternalRequest, files map[string]*multipart.FileHeader, storageProvider storageProvider.StorageInterface) (entities.InternalAuthResponse, error) {

	// Validation
	err := validations.ValidateCreateInternalRequest(us.validate, internalRequest, files)
	if err != nil {
		return entities.InternalAuthResponse{}, err
	}

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
				return entities.InternalAuthResponse{}, web.WebError{Code: 500, ProductionMessage: "server error", DevelopmentMessage: "Cannot upload file image"}
			}
			user.Avatar = fileURL
		}
	}

	// Insert ke sistem melewati repository
	user, err = us.userRepo.InsertUser(user)
	if err != nil {
		return entities.InternalAuthResponse{}, err
	}

	// Konversi hasil repository menjadi user response
	userRes := entities.InternalResponse{}
	copier.Copy(&userRes, &user)

	// generate token
	token, err := _middleware.CreateToken(int(user.ID), user.Name, user.Role)
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

package user

import (
	"kost/deliveries/helpers"
	"kost/entities"
	userRepository "kost/repositories/user"

	"github.com/jinzhu/copier"
)

type UserService struct {
	userRepo userRepository.UserRepositoryInterface
}

func NewUserService(repository userRepository.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepo: repository,
	}
}

var linkUrl string = "https://belajar-be.s3.ap-southeast-1.amazonaws.com/Avatar/1653973235.png"

func (us *UserService) CreateUser(internalRequest entities.CreateUserRequest, url string) (entities.User, error) {

	// Konversi user request menjadi domain untuk diteruskan ke repository
	user := entities.User{}
	copier.Copy(&user, &internalRequest)
	if url != "" {
		user.Avatar = url
	}
	// Insert ke sistem melewati repository
	respond, err := us.userRepo.InsertUser(user)
	if err != nil {
		return user, err
	}

	return respond, nil
}

func (us *UserService) GetbyID(id uint) (entities.CustomerResponse, error) {

	user, err := us.userRepo.GetUserID(id)
	if err != nil {
		return entities.CustomerResponse{}, err
	}
	userRes := entities.CustomerResponse{}
	copier.Copy(&userRes, &user)

	return userRes, err
}

func (us *UserService) GetAllMember() ([]entities.GetAll, error) {

	user, err := us.userRepo.GetAllUser()
	if err != nil {
		return []entities.GetAll{}, err
	}
	var userRes []entities.GetAll
	copier.Copy(&userRes, &user)
	return userRes, err
}

func (us *UserService) UpdateInternal(internalRequest entities.UpdateInternalRequest, id uint, url string) (entities.InternalResponse, error) {

	// Get user by ID via repository
	user, err := us.userRepo.GetUserID(id)
	if err != nil {
		return entities.InternalResponse{}, err
	}

	// Konversi dari request ke domain entities user - mengabaikan nilai kosong pada request
	copier.CopyWithOption(&user, &internalRequest, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	// Hanya hash password jika password juga diganti (tidak kosong)
	if internalRequest.Password != "" {
		hashedPassword, _ := helpers.HashPassword(internalRequest.Password)
		user.Password = hashedPassword
	}
	if url != "" {
		user.Avatar = url
	}

	// Update via repository
	user, err = us.userRepo.UpdateUser(id, user)
	if err != nil {
		return entities.InternalResponse{}, err
	}
	// Konversi user domain menjadi user response
	userRes := entities.InternalResponse{}
	copier.Copy(&userRes, &user)

	return userRes, err
}

func (us *UserService) UpdateCustomer(customerRequest entities.UpdateCustomerRequest, id uint, url string) (entities.CustomerResponse, error) {

	// Get user by ID via repository
	user, err := us.userRepo.GetUserID(id)
	if err != nil {
		return entities.CustomerResponse{}, err
	}

	// Konversi dari request ke domain entities user - mengabaikan nilai kosong pada request
	copier.CopyWithOption(&user, &customerRequest, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	// Hanya hash password jika password juga diganti (tidak kosong)
	if customerRequest.Password != "" {
		hashedPassword, _ := helpers.HashPassword(user.Password)
		user.Password = hashedPassword
	}
	if url != "" {
		user.Avatar = url
	}

	// Update via repository
	user, err = us.userRepo.UpdateUser(id, user)
	// Konversi user domain menjadi user response
	userRes := entities.CustomerResponse{}
	copier.Copy(&userRes, &user)

	return userRes, err
}

func (us *UserService) DeleteInternal(id uint) error {

	// Cari user berdasarkan ID via repo

	// Delete via repository
	err := us.userRepo.DeleteUser(id)
	return err
}

func (us *UserService) DeleteCustomer(id uint) error {

	// Cari user berdasarkan ID via repo

	// Delete via repository
	err := us.userRepo.DeleteUser(id)
	return err
}

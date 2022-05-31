package forgot

import (
	"kost/deliveries/helpers"
	"kost/deliveries/middlewares"
	"kost/entities"
	userRepository "kost/repositories/user"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

type forgotService struct {
	userRepo userRepository.UserRepositoryInterface
	validate *validator.Validate
}

func NewforgotService(repository userRepository.UserRepositoryInterface, valid *validator.Validate) *forgotService {
	return &forgotService{
		userRepo: repository,
		validate: valid,
	}
}

func (f *forgotService) GetToken(email string) (entities.InternalAuthResponse, error) {
	user := entities.User{}
	copier.Copy(&user, &email)

	user, err := f.userRepo.FindByUser(email)
	if err != nil {
		return entities.InternalAuthResponse{}, err
	}

	userRes := entities.CustomerResponse{}
	copier.Copy(&userRes, &user)

	token, err := middlewares.CreateToken(1, user.Name, user.Role)
	if err != nil {
		return entities.InternalAuthResponse{}, err
	}

	authRes := entities.InternalAuthResponse{
		Token: token,
	}
	return authRes, nil
}

func (f *forgotService) ResetPassword(reqNew entities.ForgotPassword) (entities.User, error) {
	user := entities.User{}
	id, _, _ := middlewares.ReadToken(reqNew.Token)
	hashedPassword, _ := helpers.HashPassword(reqNew.Password)
	user.Password = hashedPassword

	res, err := f.userRepo.UpdateUser(id, user)
	if err != nil {
		return entities.User{}, err
	}

	return res, nil
}

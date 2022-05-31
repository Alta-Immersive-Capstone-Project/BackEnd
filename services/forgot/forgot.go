package forgot

import (
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

func (f *forgotService) SendEmail(email string) (entities.InternalAuthResponse, error) {
	user := entities.User{}
	copier.Copy(&user, &email)

	user, err := f.userRepo.FindByUser(email)
	if err != nil {
		return entities.InternalAuthResponse{}, err
	}

	userRes := entities.CustomerResponse{}
	copier.Copy(&userRes, &user)

	token, err := middlewares.CreateToken(int(user.ID), user.Name, user.Role)
	if err != nil {
		return entities.InternalAuthResponse{}, err
	}

	authRes := entities.InternalAuthResponse{
		Token: token,
	}
	return authRes, nil
}

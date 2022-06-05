package forgot

import (
	"kost/entities"
	userRepository "kost/repositories/user"

	"github.com/jinzhu/copier"
)

type forgotService struct {
	userRepo userRepository.UserRepositoryInterface
}

func NewforgotService(repository userRepository.UserRepositoryInterface) *forgotService {
	return &forgotService{
		userRepo: repository,
	}
}

func (f *forgotService) FindUserByEmail(email string) (entities.User, error) {
	user := entities.User{}
	copier.Copy(&user, &email)

	Respond, err := f.userRepo.FindByUser(email)
	if err != nil {
		return entities.User{}, err
	}

	return Respond, nil
}

func (f *forgotService) ResetPassword(id int, password string) (entities.User, error) {

	user := entities.User{Password: password}
	res, err := f.userRepo.UpdateUser(uint(id), user)
	if err != nil {
		return entities.User{}, err
	}

	return res, nil
}

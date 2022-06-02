package user

import "kost/entities"

type UserRepositoryInterface interface {
	InsertUser(newUser entities.User) (entities.User, error)
	GetUserID(id uint) (entities.User, error)
	GetAllUser() ([]entities.User, error)
	UpdateUser(id uint, user entities.User) (entities.User, error)
	DeleteUser(id uint) error
	FindByUser(value string) (entities.User, error)
}

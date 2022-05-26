package user

import "backend/be8/entities"

type UserRepositoryInterface interface {
	InsertUser(newUser entities.User) (entities.User, error)
	GetUserID(ID int) (entities.User, error)
	UpdateUser(ID int, update entities.User) (entities.User, error)
	DeleteUser(ID int) (entities.User, error)
}

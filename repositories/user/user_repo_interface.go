package user
<<<<<<< HEAD

import "backend/be8/entities"

type UserRepositoryInterface interface {
	InsertUser(newUser entities.User) (entities.User, error)
	GetUserID(ID int) (entities.User, error)
	UpdateUser(ID int, update entities.User) (entities.User, error)
	DeleteUser(ID int) (entities.User, error)
}
=======
>>>>>>> c4ca72a1ed7c4b21751f758877248c248310c2f9

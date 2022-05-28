package user
<<<<<<< HEAD

import "backend/be8/entities"

type UserRepositoryInterface interface {
	InsertUser(newUser entities.User) (entities.User, error)
	GetUserID(id int) (entities.User, error)
	UpdateUser(id int, user entities.User) (entities.User, error)
	DeleteUser(id int) error
	FindByUser(field string, value string) (entities.User, error)
}
=======
>>>>>>> c4ca72a1ed7c4b21751f758877248c248310c2f9

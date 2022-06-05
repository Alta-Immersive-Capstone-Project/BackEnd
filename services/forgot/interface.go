package forgot

import "kost/entities"

type ForgotInterface interface {
	FindUserByEmail(email string) (entities.User, error)
	ResetPassword(id int, password string) (entities.User, error)
}

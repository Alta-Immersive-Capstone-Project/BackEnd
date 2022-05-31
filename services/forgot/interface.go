package forgot

import "kost/entities"

type ForgotInterface interface {
	GetToken(email string) (entities.InternalAuthResponse, error)
	ResetPassword(id int, password string) (entities.CustomerResponse, error)
}

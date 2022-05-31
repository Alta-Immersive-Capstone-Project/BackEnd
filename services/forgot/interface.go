package forgot

import "kost/entities"

type ForgotInterface interface {
	GetToken(email string) (entities.InternalAuthResponse, error)
	ResetPassword(reqNew entities.ForgotPassword) (entities.User, error)
}

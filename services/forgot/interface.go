package forgot

import "kost/entities"

type ForgotInterface interface {
	SendEmail(email string) (entities.InternalAuthResponse, error)
}

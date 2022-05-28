package auth

import (
	"kost/entities"
)

type AuthServiceInterface interface {
	Login(AuthReq entities.AuthRequest) (interface{}, error)
}

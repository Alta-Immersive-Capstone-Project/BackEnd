package auth

import (
	"backend/be8/entities"
)

type AuthServiceInterface interface {
	Login(AuthReq entities.AuthRequest) (interface{}, error)
}

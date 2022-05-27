package auth

import (
	"backend/be8/deliveries/helpers"
	middleware "backend/be8/deliveries/middlewares"
	"backend/be8/entities"
	web "backend/be8/entities/web"
	userRepository "backend/be8/repositories/user"

	"github.com/jinzhu/copier"
)

type AuthService struct {
	userRepo userRepository.UserRepositoryInterface
}

func NewAuthService(userRepo userRepository.UserRepositoryInterface) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

/*
 * Auth Service - Login
 * -------------------------------
 * Mencari user berdasarkan ID
 */
func (as AuthService) Login(authReq entities.AuthRequest) (interface{}, error) {

	// Get user by username via repository
	user, err := as.userRepo.FindByUser("email", authReq.Email)
	if err != nil {
		return entities.CustomerAuthResponse{}, web.WebError{Code: 401, Message: "Invalid Credential"}
	}

	// Verify password
	if !helpers.CheckPasswordHash(authReq.Password, user.Password) {
		return entities.CustomerAuthResponse{}, web.WebError{Code: 401, Message: "Invalid password"}
	}

	if user.Role != "customer" {

		userRes := entities.InternalResponse{}
		copier.Copy(&userRes, &user)

		// Create token
		token, err := middleware.CreateToken(int(userRes.ID), userRes.Name, userRes.Role)
		if err != nil {
			return entities.InternalAuthResponse{}, web.WebError{Code: 500, Message: "Error create token"}
		}

		return entities.InternalAuthResponse{
			Token: token,
			User:  userRes,
		}, nil
	}

	// Konversi menjadi customer response
	userRes := entities.CustomerResponse{}
	copier.Copy(&userRes, &user)

	// Create token
	token, err := middleware.CreateToken(int(userRes.ID), userRes.Name, "customer")
	if err != nil {
		return entities.CustomerAuthResponse{}, web.WebError{Code: 500, Message: "Error create token"}
	}

	return entities.CustomerAuthResponse{
		Token: token,
		User:  userRes,
	}, nil
}

package auth

import (
	"kost/deliveries/helpers"
	middleware "kost/deliveries/middlewares"
	"kost/entities"
	userRepository "kost/repositories/user"

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
		return entities.CustomerAuthResponse{}, err
	}

	// Verify password
	if !helpers.CheckPasswordHash(authReq.Password, user.Password) {
		return entities.CustomerAuthResponse{}, err
	}

	if user.Role != "customer" {

		userRes := entities.InternalResponse{}
		copier.Copy(&userRes, &user)

		// Create token
		token, err := middleware.CreateToken(int(userRes.ID), userRes.Name, userRes.Role)
		if err != nil {
			return entities.InternalAuthResponse{}, err
		}

		return token, nil
	}

	// Konversi menjadi customer response
	userRes := entities.CustomerResponse{}
	copier.Copy(&userRes, &user)

	// Create token
	token, err := middleware.CreateToken(int(userRes.ID), userRes.Name, "customer")
	if err != nil {
		return entities.CustomerAuthResponse{}, err
	}

	return token, nil
}

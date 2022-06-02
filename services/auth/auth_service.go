package auth

import (
	"kost/deliveries/helpers"
	middleware "kost/deliveries/middlewares"
	"kost/entities"
	userRepository "kost/repositories/user"
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
func (as AuthService) Login(authReq entities.AuthRequest) (string, error) {

	// Get user by username via repository
	user, err := as.userRepo.FindByUser(authReq.Email)
	if err != nil {
		return "", err
	}

	// Verify password
	if !helpers.CheckPasswordHash(authReq.Password, user.Password) {
		return "", err
	}

	if user.Role != "customer" {

		// Create token
		token, err := middleware.CreateToken(int(user.ID), user.Name, user.Role)
		if err != nil {
			return "", err
		}
		return token, nil
	}

	// Create token
	token, err := middleware.CreateToken(int(user.ID), user.Name, "customer")
	if err != nil {
		return "", err
	}

	return token, nil
}

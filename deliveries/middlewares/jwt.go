package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId float64, role, email string) (string, error) {
	claims := jwt.MapClaims{}

	claims["userId"] = userId
	claims["role"] = role
	claims["email"] = email

	claims["expired"] = time.Now().Add(time.Hour * 3).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("K05T"))
}

func ExtractTokenUserId(e echo.Context) float64 {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return userId
	}
	return 0
}

func ExtractTokenRole(e echo.Context) string {
	role := e.Get("user").(*jwt.Token)
	if role.Valid {
		claims := role.Claims.(jwt.MapClaims)
		role := claims["role"].(string)
		return role
	}
	return ""
}

func ExtractTokenEmail(e echo.Context) string {
	email := e.Get("user").(*jwt.Token)
	if email.Valid {
		claims := email.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		return email
	}
	return ""
}

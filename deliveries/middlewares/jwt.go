package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const secretJwt = "BACKENDBE8"

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(secretJwt),
		SigningMethod: jwt.SigningMethodHS256.Name,
	})
}

func CreateToken(id int, name string, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["name"] = name
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretJwt))
}

func ReadToken(token interface{}) (int, string, error) {
	tokenID := token.(*jwt.Token)
	claims := tokenID.Claims.(jwt.MapClaims)
	id := int(claims["id"].(float64))
	role := claims["role"].(string)
	return id, role, nil
}

func ExtractTokenUserId(e echo.Context) float64 {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(float64)
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

func ExtractTokenName(e echo.Context) string {
	name := e.Get("user").(*jwt.Token)
	if name.Valid {
		claims := name.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		return name
	}
	return ""
}

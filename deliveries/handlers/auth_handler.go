package handlers

import (
	"kost/deliveries/helpers"
	"kost/entities"
	authService "kost/services/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService *authService.AuthService
}

func NewAuthHandler(service *authService.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: service,
	}
}

func (handler AuthHandler) Login(c echo.Context) error {
	// Populate request input
	authReq := entities.AuthRequest{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	// define link hateoas
	// call auth service login
	authRes, err := handler.authService.Login(authReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
	}

	// send response
	return c.JSON(http.StatusOK, helpers.LoginOK(authRes))
}

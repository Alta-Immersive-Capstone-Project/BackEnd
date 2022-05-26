package handlers

import (
	"backend/be8/configs"
	"backend/be8/deliveries/helpers"
	"backend/be8/entities"
	"backend/be8/entities/web"
	authService "backend/be8/services/auth"

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
	links := map[string]string{"self": configs.Get().App.BaseURL + "/api/auth"}

	// call auth service login
	authRes, err := handler.authService.Login(authReq)
	if err != nil {
		return helpers.WebErrorResponse(c, err, links)
	}

	// send response
	return c.JSON(200, web.SuccessResponse{
		Status: "OK",
		Code:   200,
		Error:  nil,
		Links:  links,
		Data:   authRes,
	})
}

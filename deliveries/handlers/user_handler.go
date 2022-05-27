package handlers

import (
	"backend/be8/deliveries/helpers"
	"strconv"

	middleware "backend/be8/deliveries/middlewares"
	"backend/be8/entities"
	"backend/be8/entities/web"
	storageProvider "backend/be8/services/storage"
	userService "backend/be8/services/user"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService     *userService.UserService
	storageProvider storageProvider.StorageInterface
}

func NewUserHandler(service *userService.UserService, storageProvider storageProvider.StorageInterface) *UserHandler {
	return &UserHandler{
		userService:     service,
		storageProvider: storageProvider,
	}
}

/*
 * User Handler - Create
 * -------------------------------
 * Registrasi User kedalam sistem dan
 * mengembalikan token
 */
func (handler UserHandler) CreateInternal(c echo.Context) error {

	// Bind request ke user request
	userReq := entities.CreateUserRequest{}
	c.Bind(&userReq)

	token := c.Get("user")
	_, role, err := middleware.ReadToken(token)

	if err != nil || role != "admin" {
		return c.JSON(http.StatusUnauthorized, web.ErrorResponse{
			Code:   http.StatusUnauthorized,
			Status: "ERROR",
			Error:  "unauthorized",
		})
	}

	// Read files
	files := map[string]*multipart.FileHeader{}
	avatar, _ := c.FormFile("avatar")
	if avatar != nil {
		files["avatar"] = avatar
	}

	// registrasi user via call user service
	userRes, err := handler.userService.CreateUser(userReq, files, handler.storageProvider)
	if err != nil {
		return helpers.WebErrorResponse(c, err)
	}

	// response
	return c.JSON(http.StatusCreated, web.SuccessResponse{
		Status: "OK",
		Code:   http.StatusCreated,
		Error:  nil,
		Data:   userRes,
	})
}

func (handler UserHandler) CreateCustomer(c echo.Context) error {

	// Bind request ke user request
	userReq := entities.CreateUserRequest{}
	c.Bind(&userReq)

	// Read files
	files := map[string]*multipart.FileHeader{}
	avatar, _ := c.FormFile("avatar")
	if avatar != nil {
		files["avatar"] = avatar
	}

	// registrasi user via call user service
	userRes, err := handler.userService.CreateUser(userReq, files, handler.storageProvider)
	if err != nil {
		return helpers.WebErrorResponse(c, err)
	}

	// response
	return c.JSON(http.StatusCreated, web.SuccessResponse{
		Status: "OK",
		Code:   http.StatusCreated,
		Error:  nil,
		Data:   userRes,
	})
}

func (handler UserHandler) UpdateInternal(c echo.Context) error {

	// Bind request to user request
	userReq := entities.UpdateInternalRequest{}
	c.Bind(&userReq)

	// Get token
	token := c.Get("user")
	id, _ := strconv.Atoi(c.Param("id"))
	_, role, err := middleware.ReadToken(token)

	if err != nil || role != "admin" {
		return c.JSON(http.StatusUnauthorized, web.ErrorResponse{
			Code:   http.StatusUnauthorized,
			Status: "ERROR",
			Error:  "unauthorized",
		})
	}

	files := map[string]*multipart.FileHeader{}
	avatar, _ := c.FormFile("avatar")
	if avatar != nil {
		files["avatar"] = avatar
	}

	// Update via user service call
	userRes, err := handler.userService.UpdateInternal(userReq, id, files, handler.storageProvider)
	if err != nil {
		return helpers.WebErrorResponse(c, err)
	}

	// response
	return c.JSON(200, web.SuccessResponse{
		Status: "OK",
		Code:   200,
		Error:  nil,
		Data:   userRes,
	})
}

func (handler UserHandler) UpdateCustomer(c echo.Context) error {

	// Bind request to user request

	userReq := entities.UpdateCustomerRequest{}
	c.Bind(&userReq)

	// Get token
	token := c.Get("user")
	id, _ := strconv.Atoi(c.Param("id"))
	idToken, role, err := middleware.ReadToken(token)

	if id != idToken || role != "customer" || err != nil {
		return c.JSON(http.StatusUnauthorized, web.ErrorResponse{
			Code:   http.StatusUnauthorized,
			Status: "ERROR",
			Error:  "unauthorized",
		})
	}

	files := map[string]*multipart.FileHeader{}
	avatar, _ := c.FormFile("avatar")
	if avatar != nil {
		files["avatar"] = avatar
	}

	// Update via user service call
	userRes, err := handler.userService.UpdateCustomer(userReq, id, files, handler.storageProvider)
	if err != nil {
		return helpers.WebErrorResponse(c, err)
	}

	// response
	return c.JSON(200, web.SuccessResponse{
		Status: "OK",
		Code:   200,
		Error:  nil,
		Data:   userRes,
	})
}

func (handler UserHandler) DeleteInternal(c echo.Context) error {

	token := c.Get("user")
	id, _ := strconv.Atoi(c.Param("id"))
	_, role, err := middleware.ReadToken(token)

	if err != nil || role != "admin" {
		return c.JSON(http.StatusUnauthorized, web.ErrorResponse{
			Code:   http.StatusUnauthorized,
			Status: "ERROR",
			Error:  "unauthorized",
		})
	}

	// call delete service
	err = handler.userService.DeleteInternal(id, handler.storageProvider)
	if err != nil {
		return helpers.WebErrorResponse(c, err)
	}

	// response
	return c.JSON(200, web.SuccessResponse{
		Status: "OK",
		Code:   200,
		Error:  nil,

		Data: map[string]interface{}{
			"id": id,
		},
	})
}

func (handler UserHandler) DeleteCustomer(c echo.Context) error {

	token := c.Get("user")
	id, _ := strconv.Atoi(c.Param("id"))
	idToken, role, err := middleware.ReadToken(token)
	if id != idToken || role != "customer" || err != nil {
		return c.JSON(http.StatusUnauthorized, web.ErrorResponse{
			Code:   http.StatusUnauthorized,
			Status: "ERROR",
			Error:  "unauthorized",
		})
	}

	// call delete service
	err = handler.userService.DeleteCustomer(id, handler.storageProvider)
	if err != nil {
		return helpers.WebErrorResponse(c, err)
	}

	// response
	return c.JSON(200, web.SuccessResponse{
		Status: "OK",
		Code:   200,
		Error:  nil,

		Data: map[string]interface{}{
			"id": id,
		},
	})
}

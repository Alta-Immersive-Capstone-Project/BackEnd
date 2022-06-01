package handlers

import (
	"kost/deliveries/helpers"
	validation "kost/deliveries/validations"
	"strconv"

	middleware "kost/deliveries/middlewares"
	"kost/entities"
	storageProvider "kost/services/storage"
	userService "kost/services/user"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserHandler struct {
	userService     userService.UserServiceInterface
	storageProvider storageProvider.StorageInterface
	valid           validation.Validation
}

func NewUserHandler(service userService.UserServiceInterface, storageProvider storageProvider.StorageInterface, Valid validation.Validation) *UserHandler {
	return &UserHandler{
		userService:     service,
		storageProvider: storageProvider,
		valid:           Valid,
	}
}

/*
 * User Handler - Create
 * -------------------------------
 * Registrasi User kedalam sistem dan
 * mengembalikan token
 */
func (handler *UserHandler) CreateInternal() echo.HandlerFunc {
	return func(c echo.Context) error {

		// Bind request ke user request
		var userReq entities.CreateUserRequest
		err := c.Bind(&userReq)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
		}

		err = handler.valid.Validation(userReq)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
		}
		token := c.Get("user")
		_, role, err := middleware.ReadToken(token)

		if err != nil || role != "admin" {
			return c.JSON(http.StatusUnauthorized, helpers.ErrorAuthorize())
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
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}

		// response
		return c.JSON(http.StatusCreated, helpers.StatusCreate("Success Create "+userRes.User.Role, userRes))
	}
}

func (handler *UserHandler) CreateCustomer() echo.HandlerFunc {
	return func(c echo.Context) error {

		// Bind request ke user request
		var userReq entities.CreateUserRequest
		err := c.Bind(&userReq)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
		}

		err = handler.valid.Validation(userReq)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
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
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}

		// response
		return c.JSON(http.StatusCreated, helpers.StatusCreate("Success Create "+userRes.User.Role, userRes))
	}
}
func (handler *UserHandler) GetCustomerByID() echo.HandlerFunc {
	return func(c echo.Context) error {

		// Get token

		id, _ := strconv.Atoi(c.Param("id"))
		idToken, role := middleware.ExtractTokenRoleID(c)

		if id != int(idToken) || role != "customer" {
			log.Warn("error")
			return c.JSON(http.StatusUnauthorized, helpers.ErrorAuthorize())
		}

		files := map[string]*multipart.FileHeader{}
		avatar, _ := c.FormFile("avatar")
		if avatar != nil {
			files["avatar"] = avatar
		}

		// Update via user service call
		userRes, err := handler.userService.GetCustomer(int(idToken))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}

		// response
		return c.JSON(http.StatusOK, helpers.StatusGetDataID("Success Get Data"+userRes.Name, userRes))

	}
}

func (handler *UserHandler) GetAllMember() echo.HandlerFunc {
	return func(c echo.Context) error {

		// Get token
		token := c.Get("user")

		_, role, err := middleware.ReadToken(token)

		if role != "admin" || err != nil {
			return c.JSON(http.StatusUnauthorized, helpers.ErrorAuthorize())
		}

		files := map[string]*multipart.FileHeader{}
		avatar, _ := c.FormFile("avatar")
		if avatar != nil {
			files["avatar"] = avatar
		}

		// Update via user service call
		userRes, err := handler.userService.GetAllMember()
		if err != nil {
			log.Warn("")
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}

		// response
		return c.JSON(http.StatusOK, helpers.StatusGetAll("Success Get All Member", userRes))
	}

}
func (handler *UserHandler) UpdateInternal() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Bind request to user request
		userReq := entities.UpdateInternalRequest{}
		c.Bind(&userReq)

		// Get token
		token := c.Get("user")
		id, _ := strconv.Atoi(c.Param("id"))
		_, role, err := middleware.ReadToken(token)

		if err != nil || role != "admin" {
			return c.JSON(http.StatusUnauthorized, helpers.ErrorAuthorize())
		}

		files := map[string]*multipart.FileHeader{}
		avatar, _ := c.FormFile("avatar")
		if avatar != nil {
			files["avatar"] = avatar
		}

		// Update via user service call
		userRes, err := handler.userService.UpdateInternal(userReq, id, files, handler.storageProvider)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}

		// response
		return c.JSON(http.StatusOK, helpers.StatusUpdate("Success Update "+userRes.Role, userRes))
	}
}

func (handler *UserHandler) UpdateCustomer() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Bind request to user request

		userReq := entities.UpdateCustomerRequest{}
		c.Bind(&userReq)

		// Get token
		token := c.Get("user")
		id, _ := strconv.Atoi(c.Param("id"))
		idToken, role, err := middleware.ReadToken(token)

		if id != idToken || role != "customer" || err != nil {
			return c.JSON(http.StatusUnauthorized, helpers.ErrorAuthorize())
		}

		files := map[string]*multipart.FileHeader{}
		avatar, _ := c.FormFile("avatar")
		if avatar != nil {
			files["avatar"] = avatar
		}

		// Update via user service call
		userRes, err := handler.userService.UpdateCustomer(userReq, id, files, handler.storageProvider)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}

		// response
		return c.JSON(http.StatusOK, helpers.StatusUpdate("Success Update Customer", userRes))
	}
}

func (handler *UserHandler) DeleteInternal() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		id, _ := strconv.Atoi(c.Param("id"))
		_, role, err := middleware.ReadToken(token)

		if err != nil || role != "admin" {
			return c.JSON(http.StatusUnauthorized, helpers.ErrorAuthorize())
		}

		// call delete service
		err = handler.userService.DeleteInternal(id, handler.storageProvider)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}

		// response
		return c.JSON(http.StatusOK, helpers.StatusDelete())
	}
}

func (handler *UserHandler) DeleteCustomer() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		id, _ := strconv.Atoi(c.Param("id"))
		idToken, role, err := middleware.ReadToken(token)
		if id != idToken || role != "customer" || err != nil {
			return c.JSON(http.StatusUnauthorized, helpers.ErrorAuthorize())
		}

		// call delete service
		err = handler.userService.DeleteCustomer(id, handler.storageProvider)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}

		// response
		return c.JSON(http.StatusOK, helpers.StatusDelete())

	}
}

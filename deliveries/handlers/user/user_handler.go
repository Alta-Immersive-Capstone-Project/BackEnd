package handlers

import (
	"kost/deliveries/helpers"
	validation "kost/deliveries/validations"
	"kost/utils/s3"
	"strconv"
	"strings"
	"time"

	middleware "kost/deliveries/middlewares"
	"kost/entities"
	userService "kost/services/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserHandler struct {
	userService userService.UserServiceInterface
	s3          s3.S3Control
	valid       validation.Validation
}

func NewUserHandler(service userService.UserServiceInterface, S3 s3.S3Control, Valid validation.Validation) *UserHandler {
	return &UserHandler{
		userService: service,
		s3:          S3,
		valid:       Valid,
	}
}

/*
 * User Handler - Create
 * -------------------------------
 * Registrasi User kedalam sistem dan
 * mengembalikan token
 */
var linkUrl string = "https://belajar-be.s3.ap-southeast-1.amazonaws.com/Avatar/1653973235.png"

func (handler *UserHandler) CreateInternal(c echo.Context) error {

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
	avatar, _ := c.FormFile("avatar")
	url := ""
	if avatar != nil {
		msg, err := validation.ValidationAvatar(avatar)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadImage(msg))
		}

		filename := "Avatar/" + userReq.Name + strconv.Itoa(int(time.Now().Unix())) + ".png"
		url, _ = handler.s3.UploadFileToS3(filename, *avatar)
	}

	// registrasi user via call user service
	userRes, err := handler.userService.CreateUser(userReq, url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
	}

	// response
	return c.JSON(http.StatusCreated, helpers.StatusCreate("Success Create "+userRes.User.Role, userRes))
}

func (handler *UserHandler) CreateCustomer(c echo.Context) error {

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
	avatar, _ := c.FormFile("avatar")
	url := ""
	if avatar != nil {
		msg, err := validation.ValidationAvatar(avatar)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadImage(msg))
		}

		filename := "Avatar/" + userReq.Name + strconv.Itoa(int(time.Now().Unix())) + ".png"
		url, _ = handler.s3.UploadFileToS3(filename, *avatar)
	}

	// registrasi user via call user service
	userRes, err := handler.userService.CreateUser(userReq, url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
	}

	// response
	return c.JSON(http.StatusCreated, helpers.StatusCreate("Success Create "+userRes.User.Role, userRes))
}

func (handler *UserHandler) GetByID(c echo.Context) error {
	idUser, _ := strconv.Atoi(c.Param("id"))
	// Update via user service call

	userRes, err := handler.userService.GetbyID(uint(idUser))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
	}

	// response
	return c.JSON(http.StatusOK, helpers.StatusGetDataID("Success Get Data"+userRes.Name, userRes))

}

func (handler *UserHandler) GetAllMember(c echo.Context) error {

	// Get token
	token := c.Get("user")

	_, role, err := middleware.ReadToken(token)

	if role != "admin" || err != nil {
		return c.JSON(http.StatusUnauthorized, helpers.ErrorAuthorize())
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

func (handler *UserHandler) UpdateInternal(c echo.Context) error {

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
	avatar, _ := c.FormFile("avatar")
	if avatar != nil {
		msg, err := validation.ValidationAvatar(avatar)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadImage(msg))
		}
	}

	// Update via user service call
	userRes, err := handler.userService.UpdateInternal(userReq, uint(id), "")

	if avatar != nil {
		var filename string
		if userRes.Avatar != "" {
			filename = userRes.Avatar
		} else {
			filename = "Avatar/" + userRes.Name + strconv.Itoa(int(time.Now().Unix())) + ".png"
		}
		file, _ := handler.s3.UploadFileToS3(filename, *avatar)
		if userRes.Avatar == "" {
			userRes, err = handler.userService.UpdateInternal(entities.UpdateInternalRequest{}, uint(id), file)
		}
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
	}

	if userRes.Avatar == "" {
		userRes.Avatar = linkUrl
	}
	// response
	return c.JSON(http.StatusOK, helpers.StatusUpdate("Success Update "+userRes.Role, userRes))
}

func (handler *UserHandler) UpdateCustomer(c echo.Context) error {

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

	avatar, _ := c.FormFile("avatar")
	if avatar != nil {
		msg, err := validation.ValidationAvatar(avatar)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadImage(msg))
		}
	}

	// Update via user service call
	userRes, err := handler.userService.UpdateCustomer(userReq, uint(id), "")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
	}
	if avatar != nil {
		var filename string
		if userRes.Avatar != "" {
			filename = userRes.Avatar
		} else {
			filename = "Avatar/" + userRes.Name + strconv.Itoa(int(time.Now().Unix())) + ".png"
		}
		file, _ := handler.s3.UploadFileToS3(filename, *avatar)
		if userRes.Avatar == "" {
			userRes, err = handler.userService.UpdateCustomer(entities.UpdateCustomerRequest{}, uint(id), file)
		}
	}
	if userRes.Avatar == "" {
		userRes.Avatar = linkUrl
	}
	// response
	return c.JSON(http.StatusOK, helpers.StatusUpdate("Success Update Customer", userRes))
}

func (handler *UserHandler) DeleteInternal(c echo.Context) error {

	token := c.Get("user")
	id, _ := strconv.Atoi(c.Param("id"))
	_, role, err := middleware.ReadToken(token)

	if err != nil || role != "admin" {
		return c.JSON(http.StatusUnauthorized, helpers.ErrorAuthorize())
	}
	res, err := handler.userService.GetbyID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
	}
	file := strings.Replace(res.Avatar, "https://belajar-be.s3.ap-southeast-1.amazonaws.com/", "", 1)
	handler.s3.DeleteFromS3(file)

	// call delete service
	err = handler.userService.DeleteInternal(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
	}

	// response
	return c.JSON(http.StatusOK, helpers.StatusDelete())
}

func (handler *UserHandler) DeleteCustomer(c echo.Context) error {

	token := c.Get("user")
	id, _ := strconv.Atoi(c.Param("id"))
	idToken, role, err := middleware.ReadToken(token)
	if id != idToken || role != "customer" || err != nil {
		return c.JSON(http.StatusUnauthorized, helpers.ErrorAuthorize())
	}
	res, err := handler.userService.GetbyID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
	}
	file := strings.Replace(res.Avatar, "https://belajar-be.s3.ap-southeast-1.amazonaws.com/", "", 1)
	handler.s3.DeleteFromS3(file)

	// call delete service
	err = handler.userService.DeleteCustomer(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
	}

	// response
	return c.JSON(http.StatusOK, helpers.StatusDelete())

}

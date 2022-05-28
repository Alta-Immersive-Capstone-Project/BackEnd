package handlers

import (
	"kost/configs"
	"kost/deliveries/helpers"

	"kost/entities"
	"kost/entities/web"
	storageProvider "kost/services/storage"
	userService "kost/services/user"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
)

type InternalHandler struct {
	userService     *userService.UserService
	storageProvider storageProvider.StorageInterface
}

func NewInternalHandler(service *userService.UserService, storageProvider storageProvider.StorageInterface) *InternalHandler {
	return &InternalHandler{
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
func (handler InternalHandler) CreateInternal(c echo.Context) error {

	// Bind request ke user request
	userReq := entities.CreateInternalRequest{}
	c.Bind(&userReq)

	// Define links (hateoas)
	links := map[string]string{"self": configs.Get().App.BaseURL + "/user/internal"}

	// Read files
	files := map[string]*multipart.FileHeader{}
	avatar, _ := c.FormFile("avatar")
	if avatar != nil {
		files["avatar"] = avatar
	}

	// registrasi user via call user service
	userRes, err := handler.userService.CreateInternal(userReq, files, handler.storageProvider)
	if err != nil {
		return helpers.WebErrorResponse(c, err, links)
	}

	// response
	return c.JSON(http.StatusCreated, web.SuccessResponse{
		Status: "OK",
		Code:   http.StatusCreated,
		Error:  nil,
		Links:  links,
		Data:   userRes,
	})
}

// func (handler InternalHandler) Update(c echo.Context) error {

// 	// Bind request to user request
// 	userReq := entities.UpdateInternalRequest{}
// 	c.Bind(&userReq)

// 	// Get token
// 	token := c.Get("user")
// 	tokenID, role, err := middleware.ReadToken(token)
// 	links := map[string]string{"self": configs.Get().App.BaseURL + "/api/customers"}
// 	if err != nil {
// 		return c.JSON(http.StatusUnauthorized, web.ErrorResponse{
// 			Code:   http.StatusUnauthorized,
// 			Status: "ERROR",
// 			Error:  "unauthorized",
// 			Links:  links,
// 		})
// 	}
// 	if role == "driver" {
// 		return c.JSON(http.StatusUnauthorized, web.ErrorResponse{
// 			Code:   http.StatusUnauthorized,
// 			Status: "ERROR",
// 			Error:  "unauthorized",
// 			Links:  links,
// 		})
// 	}

// 	// avatar
// 	files := map[string]*multipart.FileHeader{}
// 	avatar, _ := c.FormFile("avatar")
// 	if avatar != nil {
// 		files["avatar"] = avatar
// 	}

// 	if len(files) == 0 && userReq.Name == "" &&

// 		userReq.Email == "" && userReq.Gender == "" &&
// 		userReq.Password == "" && userReq.Phone == "" {
// 		return c.JSON(http.StatusBadRequest, web.ErrorResponse{
// 			Code:   http.StatusBadRequest,
// 			Status: "ERROR",
// 			Error:  "no such data filled",
// 			Links:  links,
// 		})
// 	}
// 	// Update via user service call
// 	userRes, err := handler.userService.(userReq, tokenID, files, handler.storageProvider)
// 	if err != nil {
// 		return helpers.WebErrorResponse(c, err, links)
// 	}

// 	// response
// 	return c.JSON(200, web.SuccessResponse{
// 		Status: "OK",
// 		Code:   200,
// 		Error:  nil,
// 		Links:  links,
// 		Data:   userRes,
// 	})
// }

// func (handler CustomerHandler) DeleteCustomer(c echo.Context) error {

// 	token := c.Get("user")
// 	tokenID, role, err := middleware.ReadToken(token)
// 	links := map[string]string{"self": configs.Get().App.BaseURL + "/api/customers"}
// 	if role != "customer" {
// 		return c.JSON(http.StatusUnauthorized, web.ErrorResponse{
// 			Code:   http.StatusUnauthorized,
// 			Status: "ERROR",
// 			Error:  "unauthorized",
// 			Links:  links,
// 		})
// 	}
// 	if err != nil {
// 		return c.JSON(http.StatusUnauthorized, web.ErrorResponse{
// 			Code:   http.StatusUnauthorized,
// 			Status: "ERROR",
// 			Error:  "unauthorized",
// 			Links:  links,
// 		})
// 	}

// 	// call delete service
// 	err = handler.userService.DeleteCustomer(tokenID, handler.storageProvider)
// 	if err != nil {
// 		return helpers.WebErrorResponse(c, err, links)
// 	}

// 	// response
// 	return c.JSON(200, web.SuccessResponse{
// 		Status: "OK",
// 		Code:   200,
// 		Error:  nil,
// 		Links:  links,
// 		Data: map[string]interface{}{
// 			"id": tokenID,
// 		},
// 	})
// }

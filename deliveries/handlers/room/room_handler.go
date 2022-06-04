package handlers

import (
	"fmt"
	"kost/deliveries/helpers"
	"kost/deliveries/middlewares"
	"kost/deliveries/validations"
	"kost/entities"
	"kost/services/image"
	"kost/services/room"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type HandlersRoom struct {
	service room.RoomServices
	image   image.ServiceImage
	valid   *validator.Validate
}

func NewHandlersRoom(Service room.RoomServices, image image.ServiceImage, Valid *validator.Validate) *HandlersRoom {
	return &HandlersRoom{
		service: Service,
		image:   image,
		valid:   Valid,
	}
}

// Respond Create Facility
func (h *HandlersRoom) CreateRoom() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, Role := middlewares.ExtractTokenRoleID(c)
		if Role != "admin" && Role != "supervisor" {
			return c.JSON(http.StatusForbidden, helpers.ErrorAuthorize())
		}

		var Insert entities.AddRoom
		err := c.Bind(&Insert)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
		}

		form, err := c.MultipartForm()
		if err != nil {
			return err
		}
		files := form.File["files"]

		msg, err := validations.ValidationImage(files)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadImage(msg))
		}

		err = h.valid.Struct(&Insert)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
		}

		result, err := h.service.CreateRoom(uint(id), Insert)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		err = h.image.InsertImage(files, result.ID)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		result.Images, _ = h.image.GetImage(result.ID)

		return c.JSON(http.StatusCreated, helpers.StatusCreate("Success Create room", result))
	}
}

func (h *HandlersRoom) GetAllRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, helpers.ErrorConvertID())
		}
		result, err := h.service.GetAllRoom(uint(id))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		return c.JSON(http.StatusOK, helpers.StatusGetAll("Success get room", result))
	}
}

func (h *HandlersRoom) GetIDRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, helpers.ErrorConvertID())
		}

		result, err := h.service.GetIDRoom(uint(id))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		return c.JSON(http.StatusOK, helpers.StatusGetDataID("Success Get Data room", result))
	}
}

// Respond Update Facility
func (h *HandlersRoom) UpdateRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		Role := middlewares.ExtractTokenRole(c)
		if Role != "admin" && Role != "supervisor" {
			return c.JSON(http.StatusForbidden, helpers.ErrorAuthorize())
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, helpers.ErrorConvertID())
		}
		var update entities.UpdateRoom
		err = c.Bind(&update)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
		}

		form, err := c.MultipartForm()
		if err != nil {
			return err
		}
		files := form.File["files"]

		msg, err := validations.ValidationImage(files)
		// fmt.Println("masuk pengecekan")
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadImage(msg))
		}
		fmt.Println("masuk")
		err = h.image.InsertImage(files, uint(id))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}

		result, err := h.service.UpdateRoom(uint(id), update)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.ErrorNotFound())
		}
		return c.JSON(http.StatusOK, helpers.StatusUpdate("Success Update room", result))
	}
}

// Respond Delete Facility
func (h *HandlersRoom) DeleteRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		Role := middlewares.ExtractTokenRole(c)
		if Role != "admin" && Role != "supervisor" {
			return c.JSON(http.StatusForbidden, helpers.ErrorAuthorize())
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, helpers.ErrorConvertID())
		}
		err = h.image.DeleteImage(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}

		err = h.service.DeleteRoom(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		return c.JSON(http.StatusOK, helpers.StatusDelete())
	}
}
func (h *HandlersRoom) DeleteImageUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {
		Role := middlewares.ExtractTokenRole(c)
		if Role != "admin" && Role != "supervisor" {
			return c.JSON(http.StatusForbidden, helpers.ErrorAuthorize())
		}
		type insert struct {
			Id []uint `json:"id"`
		}

		var Insert insert
		err := c.Bind(&Insert)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
		}
		fmt.Println(Insert)

		err = h.image.DeleteImagebyID(Insert.Id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		return c.JSON(http.StatusOK, helpers.StatusDelete())
	}
}

package handlers

import (
	"kost/deliveries/helpers"
	"kost/deliveries/middlewares"
	"kost/entities"
	"kost/services/room"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type HandlersRoom struct {
	service room.RoomServices
	valid   *validator.Validate
}

func NewHandlersRoom(Service room.RoomServices, Valid *validator.Validate) *HandlersRoom {
	return &HandlersRoom{
		service: Service,
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
		return c.JSON(http.StatusCreated, helpers.StatusCreate("Success Create Facility", result))
	}
}

func (h *HandlersRoom) GetAllRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := h.service.GetAllRoom()
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		return c.JSON(http.StatusCreated, helpers.StatusGetAll("Success get room", result))
	}
}
func (h *HandlersRoom) GetIDRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
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
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
		}
		var update entities.UpdateRoom
		if err := c.Bind(&update); err != nil {
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
		}

		result, err := h.service.UpdateRoom(uint(id), update)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.ErrorNotFound())
		}
		return c.JSON(http.StatusOK, helpers.StatusUpdate("Success Update Facility", result))
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
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
		}

		err = h.service.DeleteRoom(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		return c.JSON(http.StatusOK, helpers.StatusDelete())
	}
}

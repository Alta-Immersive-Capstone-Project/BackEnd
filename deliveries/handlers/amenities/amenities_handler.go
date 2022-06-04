package handlers

import (
	"kost/deliveries/helpers"
	"kost/deliveries/middlewares"
	validation "kost/deliveries/validations"
	"kost/entities"
	"kost/services/amenities"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type HandlersAmenities struct {
	service amenities.AmenitiesControl
	valid   validation.Validation
}

func NewHandlersAmenities(Service amenities.AmenitiesControl, Valid validation.Validation) *HandlersAmenities {
	return &HandlersAmenities{
		service: Service,
		valid:   Valid,
	}
}

// Respond Create Amenities
func (h *HandlersAmenities) CreateAmenities() echo.HandlerFunc {
	return func(c echo.Context) error {

		Role := middlewares.ExtractTokenRole(c)
		if Role == "consultant" || Role == "customer" {
			return c.JSON(http.StatusForbidden, helpers.ErrorAuthorize())
		}
		id := c.Param("id")
		RoomID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
		}
		var Insert entities.AddAmenities
		err = c.Bind(&Insert)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
		}
		Insert.RoomID = uint(RoomID)
		err = h.valid.Validation(&Insert)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
		}

		result, err := h.service.CreateAmenities(Insert)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		return c.JSON(http.StatusCreated, helpers.StatusCreate("Success Create Amenities", result))
	}
}

// Respond Get Amenities ID
func (h *HandlersAmenities) GetAmenitiesID() echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")
		RoomID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
		}
		result, err := h.service.GetAmenitiesID(uint(RoomID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.StatusNotFound("Data Amenities By ID Not Found"))
		}
		return c.JSON(http.StatusOK, helpers.StatusGetDataID("Success Get Data Amenities", result))
	}
}

// Respond Update Amenities
func (h *HandlersAmenities) UpdateAmenities() echo.HandlerFunc {
	return func(c echo.Context) error {
		Role := middlewares.ExtractTokenRole(c)
		if Role != "admin" && Role != "supervisor" {
			return c.JSON(http.StatusForbidden, helpers.ErrorAuthorize())
		}

		id := c.Param("id")
		RoomID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
		}
		var update entities.UpdateAmenities
		if err := c.Bind(&update); err != nil {
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
		}

		result, err := h.service.UpdateAmenities(uint(RoomID), update)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.StatusNotFound("Data Amenities By ID Not Found"))
		}
		return c.JSON(http.StatusOK, helpers.StatusUpdate("Success Update Amenities", result))
	}
}

// Respond Delete Amenities
func (h *HandlersAmenities) DeleteAmenities() echo.HandlerFunc {
	return func(c echo.Context) error {
		Role := middlewares.ExtractTokenRole(c)
		if Role != "admin" && Role != "supervisor" {
			return c.JSON(http.StatusForbidden, helpers.ErrorAuthorize())
		}

		id := c.Param("id")
		RoomID, err := strconv.Atoi(id)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
		}

		errDelete := h.service.DeleteAmenities(uint(RoomID))
		if errDelete != nil {
			return c.JSON(http.StatusNotFound, helpers.StatusNotFound("Data Amenities By ID Not Found"))
		}
		return c.JSON(http.StatusOK, helpers.StatusDelete())
	}
}

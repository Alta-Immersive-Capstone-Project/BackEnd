package handlers

import (
	"kost/deliveries/helpers"
	"kost/deliveries/middlewares"
	"kost/entities"
	"kost/services/amenities"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type HandlersAmenities struct {
	service amenities.AmenitiesControl
	valid   *validator.Validate
}

func NewHandlersAmenities(Service amenities.AmenitiesControl, Valid *validator.Validate) *HandlersAmenities {
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

		var Insert entities.AddAmenities
		err := c.Bind(&Insert)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnsupportedMediaType, helpers.ErrorBindData())
		}

		err = h.valid.Struct(&Insert)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, helpers.ErrorValidate())
		}

		result, err := h.service.CreateAmenities(Insert)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		return c.JSON(http.StatusCreated, helpers.StatusCreate("Success Create Amenities", result))
	}
}

// Respond Get All Amenities
func (h *HandlersAmenities) GetAllAmenities() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.QueryParam("house_id")
		RoomID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, helpers.ErrorConvertID())
		}
		result, err := h.service.GetAllAmenities(uint(RoomID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.ErrorNotFound())
		}
		return c.JSON(http.StatusOK, helpers.StatusGetAll("Success Get All Amenities", result))
	}
}

// Respond Get Amenities ID
func (h *HandlersAmenities) GetAmenitiesID() echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")
		amenitiesID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, helpers.ErrorConvertID())
		}
		result, err := h.service.GetAmenitiesID(uint(amenitiesID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.ErrorNotFound())
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
		amenitiesID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, helpers.ErrorConvertID())
		}
		var update entities.UpdateAmenities
		if err := c.Bind(&update); err != nil {
			return c.JSON(http.StatusUnsupportedMediaType, helpers.ErrorBindData())
		}

		result, err := h.service.UpdateAmenities(uint(amenitiesID), update)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.ErrorNotFound())
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
		amenitiesID, err := strconv.Atoi(id)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, helpers.ErrorConvertID())
		}

		errDelete := h.service.DeleteAmenities(uint(amenitiesID))
		if errDelete != nil {
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		return c.JSON(http.StatusOK, helpers.StatusDelete())
	}
}

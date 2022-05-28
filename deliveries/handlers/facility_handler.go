package handlers

import (
	"kost/deliveries/helpers"
	"kost/deliveries/middlewares"
	"kost/entities"
	"kost/services/facility"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type HandlersFacility struct {
	service facility.FacilityControl
	valid   *validator.Validate
}

func NewHandlersFacility(Service facility.FacilityControl, Valid *validator.Validate) *HandlersFacility {
	return &HandlersFacility{
		service: Service,
		valid:   Valid,
	}
}

// Respond Create Facility
func (h *HandlersFacility) CreateFacility() echo.HandlerFunc {
	return func(c echo.Context) error {

		Role := middlewares.ExtractTokenRole(c)
		if Role == "consultant" || Role == "customer" {
			return c.JSON(http.StatusForbidden, helpers.ErrorAuthorize())
		}

		var Insert entities.AddNewFacility
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

		result, err := h.service.CreateFacility(Insert)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		return c.JSON(http.StatusCreated, helpers.StatusCreate("Success Create Facility", result))
	}
}

// Respond Get All Facility
func (h *HandlersFacility) GetAllFacility() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.QueryParam("house_id")
		HouseID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, helpers.ErrorConvertID())
		}
		result, err := h.service.GetAllFacility(uint(HouseID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.ErrorNotFound())
		}
		return c.JSON(http.StatusOK, helpers.StatusGetAll("Success Get All Facility", result))
	}
}

// Respond Get Facility ID
func (h *HandlersFacility) GetFacilityID() echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")
		facilityID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, helpers.ErrorConvertID())
		}
		result, err := h.service.GetFacilityID(uint(facilityID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.ErrorNotFound())
		}
		return c.JSON(http.StatusOK, helpers.StatusGetDataID("Success Get Data Facility", result))
	}
}

// Respond Update Facility
func (h *HandlersFacility) UpdateFacility() echo.HandlerFunc {
	return func(c echo.Context) error {
		Role := middlewares.ExtractTokenRole(c)
		if Role != "admin" && Role != "supervisor" {
			return c.JSON(http.StatusForbidden, helpers.ErrorAuthorize())
		}

		id := c.Param("id")
		facilityID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, helpers.ErrorConvertID())
		}
		var update entities.UpdateFacility
		if err := c.Bind(&update); err != nil {
			return c.JSON(http.StatusUnsupportedMediaType, helpers.ErrorBindData())
		}

		result, err := h.service.UpdateFacility(uint(facilityID), update)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.ErrorNotFound())
		}
		return c.JSON(http.StatusOK, helpers.StatusUpdate("Success Update Facility", result))
	}
}

// Respond Delete Facility
func (h *HandlersFacility) DeleteFacility() echo.HandlerFunc {
	return func(c echo.Context) error {
		Role := middlewares.ExtractTokenRole(c)
		if Role != "admin" && Role != "supervisor" {
			return c.JSON(http.StatusForbidden, helpers.ErrorAuthorize())
		}

		id := c.Param("id")
		facilityID, err := strconv.Atoi(id)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, helpers.ErrorConvertID())
		}

		errDelete := h.service.DeleteFacility(uint(facilityID))
		if errDelete != nil {
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		return c.JSON(http.StatusOK, helpers.StatusDelete())
	}
}

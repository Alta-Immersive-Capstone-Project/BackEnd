package handlers

import (
	"kost/deliveries/helpers"
	"kost/deliveries/middlewares"
	"kost/entities"
	"kost/services/city"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type HandlersCity struct {
	service city.CityRepo
	valid   *validator.Validate
}

func NewHandlersCity(Service city.CityRepo, Valid *validator.Validate) *HandlersCity {
	return &HandlersCity{
		service: Service,
		valid:   Valid,
	}
}

// Respond Create Facility
func (h *HandlersCity) CreateCity() echo.HandlerFunc {
	return func(c echo.Context) error {

		Role := middlewares.ExtractTokenRole(c)
		if Role != "admin" && Role != "supervisor" {
			return c.JSON(http.StatusForbidden, helpers.ErrorAuthorize())
		}

		var Insert entities.AddCity
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

		result, err := h.service.CreateCity(Insert)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		return c.JSON(http.StatusCreated, helpers.StatusCreate("Success Create Facility", result))
	}
}

func (h *HandlersCity) GetAllCity() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := h.service.GetAllCity()
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		return c.JSON(http.StatusCreated, helpers.StatusGetAll("Success get room", result))
	}
}
func (h *HandlersCity) GetIDCity() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
		}

		result, err := h.service.GetIDCity(uint(id))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		return c.JSON(http.StatusOK, helpers.StatusGetDataID("Success Get Data room", result))
	}
}

func (h *HandlersCity) UpdateCity() echo.HandlerFunc {
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
		var update entities.City
		if err := c.Bind(&update); err != nil {
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
		}

		result, err := h.service.UpdateCity(uint(id), update)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.ErrorNotFound())
		}
		return c.JSON(http.StatusOK, helpers.StatusUpdate("Success Update Facility", result))
	}
}

// Respond Delete Facility
func (h *HandlersCity) DeleteCity() echo.HandlerFunc {
	return func(c echo.Context) error {
		Role := middlewares.ExtractTokenRole(c)
		if Role != "admin" && Role != "supervisor" {
			return c.JSON(http.StatusForbidden, helpers.ErrorAuthorize())
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
		}

		err = h.service.DeleteCity(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		return c.JSON(http.StatusOK, helpers.StatusDelete())
	}
}

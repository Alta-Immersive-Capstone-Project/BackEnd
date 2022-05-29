package house

import (
	"kost/deliveries/helpers"
	"kost/deliveries/middlewares"
	"kost/deliveries/validations"
	"kost/entities"
	"kost/services/house"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type HouseHandler struct {
	Service house.IHouseService
	Valid   validations.Validation
}

func NewHouseHandler(service house.IHouseService, valid validations.Validation) *HouseHandler {
	return &HouseHandler{
		Service: service,
		Valid:   valid,
	}
}

func (hh *HouseHandler) Store() echo.HandlerFunc {
	return func(c echo.Context) error {

		Role := middlewares.ExtractTokenRole(c)
		if Role == "consultant" || Role == "customer" {
			return c.JSON(http.StatusForbidden, helpers.ErrorAuthorize())
		}

		var request entities.HouseRequest
		err := c.Bind(&request)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnsupportedMediaType, helpers.ErrorBindData())
		}

		err = hh.Valid.Validation(request)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, helpers.ErrorValidate())
		}

		result, err := hh.Service.CreateHouse(request)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		return c.JSON(http.StatusCreated, helpers.StatusCreate("Success Create House", result))
	}
}

func (hh *HouseHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		Role := middlewares.ExtractTokenRole(c)
		if Role != "admin" && Role != "supervisor" {
			return c.JSON(http.StatusForbidden, helpers.ErrorAuthorize())
		}

		id := c.Param("id")
		houseID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, helpers.ErrorConvertID())
		}
		var update entities.HouseRequest
		if err := c.Bind(&update); err != nil {
			return c.JSON(http.StatusUnsupportedMediaType, helpers.ErrorBindData())
		}

		result, err := hh.Service.UpdateHouse(uint(houseID), update)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.ErrorNotFound())
		}
		return c.JSON(http.StatusOK, helpers.StatusUpdate("Success Update House", result))
	}
}
func (hh *HouseHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		Role := middlewares.ExtractTokenRole(c)
		if Role != "admin" && Role != "supervisor" {
			return c.JSON(http.StatusForbidden, helpers.ErrorAuthorize())
		}

		id := c.Param("id")
		houseID, err := strconv.Atoi(id)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, helpers.ErrorConvertID())
		}

		errDelete := hh.Service.DeleteHouse(uint(houseID))
		if errDelete != nil {
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		return c.JSON(http.StatusOK, helpers.StatusDelete())
	}
}
func (hh *HouseHandler) GetAllByDist() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		DistrictID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, helpers.ErrorConvertID())
		}
		result, err := hh.Service.GetAllHouseByDist(uint(DistrictID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.ErrorNotFound())
		}
		return c.JSON(http.StatusOK, helpers.StatusGetAll("Success Get All Houses by District", result))
	}
}
func (hh *HouseHandler) Show() echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")
		houseID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, helpers.ErrorConvertID())
		}
		result, err := hh.Service.GetHouseID(uint(houseID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.ErrorNotFound())
		}
		return c.JSON(http.StatusOK, helpers.StatusGetDataID("Success Get Data House", result))
	}
}

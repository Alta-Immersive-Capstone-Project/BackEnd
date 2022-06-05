package district

import (
	"kost/deliveries/helpers"
	"kost/deliveries/middlewares"
	"kost/entities"
	services "kost/services/district"
	"net/http"
	"strconv"

	validation "kost/deliveries/validations"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type DistrictHandler struct {
	Service services.DistrictControl
	Valid   validation.Validation
}

func NewDistrictHandler(service services.DistrictControl, valid validation.Validation) *DistrictHandler {
	return &DistrictHandler{
		Service: service,
		Valid:   valid,
	}
}

func (dh *DistrictHandler) Store() echo.HandlerFunc {
	return func(c echo.Context) error {

		Role := middlewares.ExtractTokenRole(c)
		if Role == "consultant" || Role == "customer" {
			return c.JSON(http.StatusForbidden, helpers.ErrorAuthorize())
		}

		var request entities.AddDistrict
		err := c.Bind(&request)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
		}

		err = dh.Valid.Validation(request)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
		}

		result, err := dh.Service.CreateDist(request)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		return c.JSON(http.StatusCreated, helpers.StatusCreate("Success Create District", result))
	}
}

func (dh *DistrictHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		Role := middlewares.ExtractTokenRole(c)
		if Role != "admin" && Role != "supervisor" {
			return c.JSON(http.StatusForbidden, helpers.ErrorAuthorize())
		}

		id := c.Param("id")
		districtID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
		}
		var update entities.UpdateDistrict
		if err := c.Bind(&update); err != nil {
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
		}

		result, err := dh.Service.UpdateDist(uint(districtID), update)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.StatusNotFound("District With ID Not Found"))
		}
		return c.JSON(http.StatusOK, helpers.StatusUpdate("Success Update District", result))
	}
}
func (dh *DistrictHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		Role := middlewares.ExtractTokenRole(c)
		if Role != "admin" && Role != "supervisor" {
			return c.JSON(http.StatusForbidden, helpers.ErrorAuthorize())
		}

		id := c.Param("id")
		districtID, err := strconv.Atoi(id)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
		}

		errDelete := dh.Service.DeleteDist(uint(districtID))
		if errDelete != nil {
			return c.JSON(http.StatusNotFound, helpers.StatusNotFound("District With ID Not Found"))
		}
		return c.JSON(http.StatusOK, helpers.StatusDelete())
	}
}
func (dh *DistrictHandler) GetAllByCity() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		CityID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
		}
		result, err := dh.Service.GetAllDist(uint(CityID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.StatusNotFound("District With ID Not Found"))
		}
		return c.JSON(http.StatusOK, helpers.StatusGetAll("Success Get All District by Cities", result))
	}
}

func (dh *DistrictHandler) Show() echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")
		districtID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
		}
		result, err := dh.Service.GetDistID(uint(districtID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.StatusNotFound("District With ID Not Found"))
		}
		return c.JSON(http.StatusOK, helpers.StatusGetDataID("Success Get Data District", result))
	}
}

func (dh *DistrictHandler) Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := dh.Service.SelectAllDistrict()
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.StatusNotFound("District With ID Not Found"))
		}
		return c.JSON(http.StatusOK, helpers.StatusGetAll("Success Select All District", result))
	}
}

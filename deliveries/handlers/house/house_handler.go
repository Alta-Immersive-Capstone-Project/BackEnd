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

		var request entities.AddHouse
		err := c.Bind(&request)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
		}

		err = hh.Valid.Validation(request)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
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
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
		}
		var update entities.UpdateHouse
		if err := c.Bind(&update); err != nil {
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
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
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
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
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
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
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
		}
		result, err := hh.Service.GetHouseID(uint(houseID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.ErrorNotFound())
		}
		return c.JSON(http.StatusOK, helpers.StatusGetDataID("Success Get Data House", result))
	}
}

func (hh *HouseHandler) Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := hh.Service.SelectAllHouse()
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.ErrorNotFound())
		}
		return c.JSON(http.StatusOK, helpers.StatusGetAll("Success get all data houses", result))
	}
}

func (hh *HouseHandler) SelectHouseByDistrict() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		DistrictID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
		}
		result, err := hh.Service.FindAllHouseByDistrict(uint(DistrictID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.ErrorNotFound())
		}
		return c.JSON(http.StatusOK, helpers.StatusGetAll("Success get all data houses", result))
	}
}
func (hh *HouseHandler) SelectHouseByCities() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		CityID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
		}
		result, err := hh.Service.FindAllHouseByCities(uint(CityID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.ErrorNotFound())
		}
		return c.JSON(http.StatusOK, helpers.StatusGetAll("Success get all data houses", result))
	}
}
func (hh *HouseHandler) SelectHouseByCtyAndDst() echo.HandlerFunc {
	return func(c echo.Context) error {
		cid := c.Param("cid")
		dist_id := c.Param("dist_id")
		DistrictID, err := strconv.Atoi(dist_id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
		}
		CityID, err := strconv.Atoi(cid)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
		}
		result, err := hh.Service.FindAllHouseByCtyAndDst(uint(CityID), uint(DistrictID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.ErrorNotFound())
		}
		return c.JSON(http.StatusOK, helpers.StatusGetAll("Success get all data houses", result))
	}
}
func (hh *HouseHandler) SearchByTitle() echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.QueryParam("title")

		result, err := hh.Service.FindHouseByTitle(title)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.ErrorNotFound())
		}
		return c.JSON(http.StatusOK, helpers.StatusGetAll("Success get all data houses", result))
	}
}

func (hh *HouseHandler) SearchBylocation() echo.HandlerFunc {
	return func(c echo.Context) error {
		rlat := c.QueryParam("rlat")
		lat, err := strconv.ParseFloat(rlat, 64)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
		}

		rlong := c.QueryParam("rlong")
		long, err := strconv.ParseFloat(rlong, 64)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.ErrorConvertID())
		}

		result, err := hh.Service.FindHouseByLocation(lat, long)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, helpers.ErrorNotFound())
		}
		return c.JSON(http.StatusOK, helpers.StatusGetAll("Success get all data houses", result))
	}
}

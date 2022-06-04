package handlers

import (
	"fmt"
	"kost/deliveries/helpers"
	"kost/deliveries/middlewares"
	validation "kost/deliveries/validations"
	"kost/entities"
	service "kost/services/transactions"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type transactionHandler struct {
	ts service.TransactionService
	v  validation.Validation
}

func NewTransactionHandler(ts service.TransactionService, v validation.Validation) *transactionHandler {
	return &transactionHandler{
		ts: ts,
		v:  v,
	}
}

func (th *transactionHandler) InsertTransaction(c echo.Context) error {
	user_id := uint(middlewares.ExtractTokenUserId(c))
	var req entities.TransactionRequest

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
	}

	err = th.v.Validation(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
	}

	response, err := th.ts.CreateTransaction(user_id, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestDuplicate(err))
	}

	return c.JSON(http.StatusCreated, helpers.StatusCreate("Success Created Transaction", response))
}

func (th *transactionHandler) GetAllTransactionbyConsultant(c echo.Context) error {
	role := middlewares.ExtractTokenRole(c)

	if role == "customer" {
		return c.JSON(http.StatusForbidden, helpers.StatusForbidden("You are not allowed to access this resource"))
	}

	response := th.ts.GetAllTransactionbyConsultant()
	if len(response) == 0 {
		return c.JSON(http.StatusNotFound, helpers.StatusNotFound("Data transaction not found"))
	}

	return c.JSON(http.StatusOK, helpers.StatusOK("Success Get All Transaction", response))
}

func (th *transactionHandler) UpdateTransaction(c echo.Context) error {
	var request entities.TransactionUpdateRequest
	user_id := uint(middlewares.ExtractTokenUserId(c))
	role := middlewares.ExtractTokenRole(c)
	booking_id := c.Param("booking_id")

	if role == "customer" {
		return c.JSON(http.StatusForbidden, helpers.StatusForbidden("You are not allowed to access this resource"))
	}

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
	}

	err = th.v.Validation(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
	}

	response, err := th.ts.UpdateTransaction(user_id, booking_id, request)
	if err != nil {
		log.Warn(err)
		return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestTrans(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.StatusOK("Success Update Transaction", response))
}

func (th *transactionHandler) UpdateCallback(c echo.Context) error {
	var request entities.Callback

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
	}

	check := helpers.Hash512(request)
	if !check {
		return c.JSON(http.StatusForbidden, helpers.StatusForbidden("Your are not allowed to access this resource"))
	}

	response, err := th.ts.UpdateCallback(request)
	if err != nil {
		return c.JSON(http.StatusForbidden, helpers.StatusUnauthorized(err))
	}

	return c.JSON(http.StatusOK, helpers.StatusOK("Success Update Status", response))
}

func (th *transactionHandler) GetAllTransactionbyUser(c echo.Context) error {
	user_id, role := middlewares.ExtractTokenRoleID(c)

	status := c.QueryParam("status")
	city, _ := strconv.Atoi(c.QueryParam("city"))
	district, _ := strconv.Atoi(c.QueryParam("district"))

	response := th.ts.GetAllTransactionbyUser(role, uint(user_id), status, uint(city), uint(district))
	fmt.Println(response)
	if len(response) == 0 {
		return c.JSON(http.StatusNotFound, helpers.StatusNotFound("Data transaction not found"))
	}

	return c.JSON(http.StatusOK, helpers.StatusOK("Success Get All Transaction", response))
}

func (th *transactionHandler) GetAllTransactionbyKost(c echo.Context) error {
	role := middlewares.ExtractTokenRole(c)
	status := c.QueryParam("status")
	duration, _ := strconv.Atoi(c.QueryParam("duration"))
	name := c.QueryParam("name")
	generate := c.QueryParam("generate")

	if role == "customer" {
		return c.JSON(http.StatusForbidden, helpers.StatusForbidden("You are not allowed to access this resource"))
	}

	response := th.ts.GetAllTransactionbyKost(duration, status, name)
	if len(response) == 0 {
		return c.JSON(http.StatusNotFound, helpers.StatusNotFound("Data transaction not found"))
	}

	if generate == "true" {
		link := th.ts.GetReport(response)
		return c.JSON(http.StatusOK, helpers.StatusOKReport("Success Get All Transaction", response, link))
	}

	return c.JSON(http.StatusOK, helpers.StatusOK("Success Get All Transaction", response))
}

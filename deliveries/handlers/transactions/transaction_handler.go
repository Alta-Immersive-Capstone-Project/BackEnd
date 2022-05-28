package handlers

import (
	"kost/deliveries/helpers"
	"kost/deliveries/middlewares"
	validation "kost/deliveries/validations"
	"kost/entities"
	service "kost/services/transactions"
	"net/http"

	"github.com/labstack/echo/v4"
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
	var request entities.TransactionRequest

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
	}

	err = th.v.Validation(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
	}

	response, err := th.ts.AddTransaction(user_id, &request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
	}

	return c.JSON(http.StatusCreated, helpers.StatusCreated("Success Created Transaction", response))
}

func (th *transactionHandler) GetAllTransactionbyCustomer(c echo.Context) error {
	customer_id := uint(middlewares.ExtractTokenUserId(c))
	status := c.QueryParam("status")

	response := th.ts.GetAllTransactionbyCustomer(customer_id, status)
	if len(response) == 0 {
		return c.JSON(http.StatusNotFound, helpers.StatusNotFound("Data transaction not found"))
	}

	return c.JSON(http.StatusOK, helpers.StatusOK("Success Get All Transaction", response))
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

	_, err = th.ts.GetTransaction(booking_id)
	if err != nil {
		return c.JSON(http.StatusForbidden, helpers.StatusForbidden("Your are not allowed to access this resource"))
	}

	response, err := th.ts.UpdateTransaction(user_id, booking_id, &request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
	}

	return c.JSON(http.StatusOK, helpers.StatusOK("Success Update Transaction", response))
}

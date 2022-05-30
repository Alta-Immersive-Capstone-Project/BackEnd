package handlers

import "github.com/labstack/echo/v4"

type TransactionHandler interface {
	InsertTransaction(c echo.Context) error
	GetAllTransactionbyCustomer(c echo.Context) error
	GetAllTransactionbyConsultant(c echo.Context) error
	UpdateTransaction(c echo.Context) error
}

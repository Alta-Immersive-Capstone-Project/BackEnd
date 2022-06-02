package handlers

import "github.com/labstack/echo/v4"

type HandleUser interface {
<<<<<<< HEAD
	CreateInternal() echo.HandlerFunc
	CreateCustomer() echo.HandlerFunc
	GetCustomerByID() echo.HandlerFunc

	GetAllMember() echo.HandlerFunc
	UpdateInternal() echo.HandlerFunc
	UpdateCustomer() echo.HandlerFunc
	DeleteInternal() echo.HandlerFunc
	DeleteCustomer() echo.HandlerFunc
=======
	CreateInternal(c echo.Context) error
	CreateCustomer(c echo.Context) error
	GetByID(c echo.Context) error
	GetAllMember(c echo.Context) error
	UpdateInternal(c echo.Context) error
	UpdateCustomer(c echo.Context) error
	DeleteInternal(c echo.Context) error
	DeleteCustomer(c echo.Context) error
>>>>>>> 3d2f172cae4224571053c1b5658836fe1402c6a9
}

package helpers

import (
	"backend/be8/entities/web"
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
)

func WebErrorResponse(c echo.Context, err error) error {
	if reflect.TypeOf(err).String() == "web.WebError" {
		webErr := err.(web.WebError)
		return c.JSON(webErr.Code, web.ErrorResponse{
			Status: "ERROR",
			Code:   webErr.Code,
			Error:  webErr.Error(),
		})
	} else if reflect.TypeOf(err).String() == "web.ValidationError" {
		valErr := err.(web.ValidationError)
		return c.JSON(valErr.Code, web.ValidationErrorResponse{
			Status: "ERROR",
			Code:   valErr.Code,
			Error:  valErr.Error(),
			Errors: valErr.Errors,
		})
	} else if reflect.TypeOf(err).String() == "web.ValidationError" {
		valErr := err.(web.ValidationError)
		return c.JSON(valErr.Code, web.ValidationErrorResponse{
			Status: "ERROR",
			Code:   valErr.Code,
			Error:  valErr.Error(),
			Errors: valErr.Errors,
		})
	}
	return c.JSON(http.StatusInternalServerError, web.ErrorResponse{
		Status: "ERROR",
		Code:   http.StatusInternalServerError,
		Error:  "Server Error",
	})

}

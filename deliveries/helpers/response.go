package helpers

import (
	"fmt"
	"kost/entities/web"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
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

func InternalServerError() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "Cannot Access Database",
		"status":  false,
	}
}

func ErrorDataEmpty() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": "Data Is Empty",
		"status":  false,
	}
}

func ErrorNotFound() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": "Data Not Found",
		"status":  false,
	}
}

func ErrorConvertID() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusNotAcceptable,
		"message": "Cannot Convert ID",
		"status":  false,
	}
}

func ErrorBindData() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusUnsupportedMediaType,
		"message": "Cannot Bind Data",
		"status":  false,
	}
}

func ErrorValidate() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusNotAcceptable,
		"message": "Validate Error",
		"status":  false,
	}
}

func ErrorAuthorize() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusForbidden,
		"message": "Access Denied",
		"status":  false,
	}
}

func StatusBadRequest(err error) map[string]interface{} {
	var messages []string

	for _, err := range err.(validator.ValidationErrors) {
		message := fmt.Sprintf("error on field %s: %s (%s)", err.Field(), err.Tag(), err.Kind().String())
		messages = append(messages, message)
	}

	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": messages,
		"data":    nil,
	}
}

func StatusOK(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": message,
		"data":    data,
	}
}

func StatusOKReview(message string, data ...interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":         http.StatusOK,
		"message":      message,
		"data":         data[0],
		"total_rating": data[1],
	}
}

func StatusCreated(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": message,
		"data":    data,
	}
}

func StatusBadRequestBind(err error) map[string]interface{} {
	var field []string
	var message string

	for i, v := range strings.Fields(string(err.Error())) {
		if i == 1 && v == "message=Syntax" {
			message = "expected=string"
		} else if i == 1 && v == "message=Unmarshal" {
			message = "expected=integer"
		} else if i == 6 {
			field = append(field, v)
		}
	}

	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": field[0] + " " + message,
		"data":    nil,
	}
}

func StatusForbidden(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusForbidden,
		"message": message,
		"data":    nil,
	}
}

func StatusNotFound(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": message,
		"data":    nil,
	}
}

func StatusDelete() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Deleted",
		"status":  true,
	}
}

func StatusCreate(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": message,
		"status":  true,
		"data":    data,
	}
}

func StatusGetAll(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": message,
		"status":  true,
		"data":    data,
	}
}

func StatusGetDataID(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": message,
		"status":  true,
		"data":    data,
	}
}

func StatusUpdate(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": message,
		"status":  true,
		"data":    data,
	}
}

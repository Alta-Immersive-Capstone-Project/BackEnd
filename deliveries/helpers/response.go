package helpers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func InternalServerError() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "Cannot Access Database",
	}
}

func ErrorDataEmpty() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": "Data Is Empty",
	}
}
func StatusBadImage(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": message,
	}
}

func ErrorNotFound() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": "Data Not Found",
	}
}

func ErrorAuthorize() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusForbidden,
		"message": "Access Denied",
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

func StatusBadRequestDuplicate(err error) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": err.Error(),
		"data":    nil,
	}
}

func LoginOK(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Login",
		"data":    data,
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
	value := data[1].([]int)
	rating := map[string]interface{}{"one": value[0], "two": value[1], "three": value[2], "four": value[3], "five": value[4]}

	return map[string]interface{}{
		"code":         http.StatusOK,
		"message":      message,
		"data":         data[0],
		"rating":       rating,
		"total_rating": data[2],
	}
}

func StatusOKReport(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":        http.StatusOK,
		"message":     message,
		"link_report": data,
	}
}

func ErrorConvertID() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusNotAcceptable,
		"message": "Cannot Convert ID",
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

func StatusUnauthorized(err error) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusForbidden,
		"message": err.Error(),
		"data":    nil,
	}
}

func ErrorS3(message string) map[string]interface{} {
	return map[string]interface{}{
		"Code":    http.StatusForbidden,
		"Message": message,
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
	}
}

func StatusCreate(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": message,
		"data":    data,
	}
}

func StatusGetAll(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": message,
		"data":    data,
	}
}

func StatusGetDataID(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": message,
		"data":    data,
	}
}

func StatusUpdate(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": message,
		"data":    data,
	}
}

func StatusBadRequestTrans(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": message,
	}
}

func ErrorRegister(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": message,
	}
}

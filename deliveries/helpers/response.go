package helpers

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

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

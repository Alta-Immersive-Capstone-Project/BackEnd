package handlers

import (
	"kost/deliveries/helpers"
	validation "kost/deliveries/validations"
	"kost/entities"
	forgotService "kost/services/forgot"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type forgotHandler struct {
	forgotService forgotService.ForgotInterface
	valid         validation.Validation
}

func NewAuthHandler(service forgotService.ForgotInterface, Valid validation.Validation) *forgotHandler {
	return &forgotHandler{
		forgotService: service,
		valid:         Valid,
	}
}

func (h *forgotHandler) SendEmail() echo.HandlerFunc {
	return func(c echo.Context) error {
		var forgotReq entities.ForgotRequest
		err := c.Bind(&forgotReq)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
		}

		err = h.valid.Validation(forgotReq)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
		}
	}
}

package handlers

import (
	"fmt"
	"kost/deliveries/helpers"
	validation "kost/deliveries/validations"
	"kost/entities"
	emailService "kost/services/email"
	forgotService "kost/services/forgot"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ForgotHandler struct {
	forgotService forgotService.ForgotInterface
	emailService  emailService.EmailService
	valid         validation.Validation
}

func NewForgotHandler(service forgotService.ForgotInterface, emailService emailService.EmailService, Valid validation.Validation) *ForgotHandler {
	return &ForgotHandler{
		forgotService: service,
		emailService:  emailService,
		valid:         Valid,
	}
}

func (h *ForgotHandler) SendEmail() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.QueryParams().Get("email")

		token, err := h.forgotService.GetToken(email)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}

		id, err := h.emailService.SendEmail("test@gmail.com", "reset password", fmt.Sprintf("Url forgot password: %s", token), email)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}

		return c.JSON(http.StatusOK, helpers.StatusOK("OK", id))
	}
}

func (h *ForgotHandler) ResetPassword() echo.HandlerFunc {
	return func(c echo.Context) error {
		resetReq := entities.ForgotPassword{}
		c.Bind(&resetReq)

		res, err := h.forgotService.ResetPassword(resetReq)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		return c.JSON(http.StatusOK, helpers.StatusOK("OK", res))
	}

}

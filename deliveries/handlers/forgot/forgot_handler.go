package handlers

import (
	"fmt"
	"kost/configs"
	"kost/deliveries/helpers"
	"kost/deliveries/middlewares"
	validation "kost/deliveries/validations"
	"kost/entities"
	forgotService "kost/services/forgot"
	emailService "kost/utils/email"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ForgotHandler struct {
	forgotService forgotService.ForgotInterface
	emailService  emailService.EmailControl
	valid         validation.Validation
}

func NewForgotHandler(service forgotService.ForgotInterface, emailService emailService.EmailControl, Valid validation.Validation) *ForgotHandler {
	return &ForgotHandler{
		forgotService: service,
		emailService:  emailService,
		valid:         Valid,
	}
}

func (h *ForgotHandler) SendEmail() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.QueryParams().Get("email")
		user, err := h.forgotService.FindUserByEmail(email)

		token, err := middlewares.CreateToken(int(user.ID), user.Name, user.Role)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
		}

		id, err := h.emailService.SendEmail("admin@sewakost.com", "reset password", generateBodyEmailForgotPassword(user.Name, token), email)

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

		token := c.Get("user")
		id, _, err := middlewares.ReadToken(token)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnauthorized, helpers.ErrorAuthorize())
		}
		hashedPassword, _ := helpers.HashPassword(resetReq.Password)
		res, err := h.forgotService.ResetPassword(id, hashedPassword)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}
		return c.JSON(http.StatusOK, helpers.StatusOK("OK", helpers.StatusUpdate("Password changed", res)))
	}

}

func generateBodyEmailForgotPassword(name string, token string) string {

	url := configs.Get().Frontend.Domain
	page := configs.Get().Frontend.ResetPage
	text := fmt.Sprintf("Hai %s,\n\nSomeone has requested a password reset.\n\nClick link to reset password:\n%s%s?token=%s\n\nIf this was a mistake, just ignore this email and nothing will happen.\nThank You\n\nBest Regards,\n\nSewa Kost Management", name, url, page, token)

	log.Warn(text)
	return text
}

package handlers

import (
	"go_ecommerce/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RestHandler struct {
	Service *services.ResetService
}

func NewRestHandler(service *services.ResetService) *RestHandler {

	return &RestHandler{Service: service}

}

func (h *RestHandler) ForgotPassword(c echo.Context) error {

	type Request struct {
		Email string `json:"email" validate:"required,email"`
	}

	req := new(Request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := h.Service.RequestPasswordReset(req.Email); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Reset link sent to your email"})
}

func (h *RestHandler) ResetPassword(c echo.Context) error {
	type Request struct {
		Token       string `json:"token" validate:"required"`
		NewPassword string `json:"new_password" validate:"required,min=6"`
	}

	req := new(Request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := h.Service.ResetPassword(req.Token, req.NewPassword); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Password reset successful"})
}

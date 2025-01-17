package handlers

import (
	"go_ecommerce/pkg/dto"
	"go_ecommerce/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {

	authService services.AuthService
	
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	
	return &AuthHandler{authService: authService}

}




func (h *AuthHandler) Register(c echo.Context) error {

	var req dto.RegisterRequest

	if err := c.Bind(&req); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	resp, err := h.authService.Register(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, resp)
}



func (h *AuthHandler) Login(c echo.Context) error {

	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	resp, err := h.authService.Login(req)

	if err != nil {

		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

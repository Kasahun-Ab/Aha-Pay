package handlers

import (
	"go_ecommerce/internal/services"
	"go_ecommerce/pkg/dto"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService    services.AuthService
	sessionService *services.UserSessionService
}

func NewAuthHandler(authService services.AuthService, sessionService *services.UserSessionService) *AuthHandler {

	return &AuthHandler{authService: authService, sessionService: sessionService}

}

func (h *AuthHandler) Register(c echo.Context) error {

	var req dto.RegisterRequest

	if err := c.Bind(&req); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	resp, err := h.authService.Register(&req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	sessionModel := dto.CreateUserSessionDTO{
		UserID:       resp.ID,
		SessionToken: resp.Token,
		IPAddress:    c.RealIP(),
		DeviceInfo:   req.DeviceInfo,
	}

	h.sessionService.CreateSession(&sessionModel)

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

	sessionModel := dto.CreateUserSessionDTO{
		UserID:       resp.ID,
		SessionToken: resp.Token,
		IPAddress:    c.RealIP(),
		DeviceInfo:   req.DeviceInfo,
	}

	h.sessionService.CreateSession(&sessionModel)

	return c.JSON(http.StatusOK, resp)
}

func (h *AuthHandler) Logout(c echo.Context) error {
	// Retrieve the token from the Authorization header
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Missing or invalid Authorization header",
		})
	}

	// Extract the token from the Authorization header
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Token is empty",
		})
	}

	// Retrieve the session associated with the token
	session, err := h.sessionService.GetSessionByToken(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "Session not found",
		})
	}

	// Delete the session from the database/service
	if err := h.sessionService.DeleteSession(session.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to delete session",
		})
	}

	// Respond with a success message
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Logout successful",
	})
}

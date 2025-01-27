package handlers

import (
	"go_ecommerce/internal/services"
	"go_ecommerce/pkg/dto"
	"go_ecommerce/pkg/utils"
	"net/http"

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

	resp, err, cookie := h.authService.Register(&req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	sessionModel := dto.CreateUserSessionDTO{
		SessionToken: cookie.Value,
		IPAddress:    c.RealIP(),
		DeviceInfo:   req.DeviceInfo,
	}

	h.sessionService.CreateSession(sessionModel)

	c.SetCookie(cookie)
	return c.JSON(http.StatusCreated, resp)
}

func (h *AuthHandler) Login(c echo.Context) error {

	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	resp, err, cookie := h.authService.Login(req)

	if err != nil {

		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	sessionModel := dto.CreateUserSessionDTO{
		SessionToken: cookie.Value,
		IPAddress:    c.RealIP(),
		DeviceInfo:   req.DeviceInfo,
	}

	h.sessionService.CreateSession(sessionModel)

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, resp)
}


func (h *AuthHandler) Logout(c echo.Context) error {

	cookie, err := utils.GetCookie(c, "token")

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid cookie"})
	}

	session, err := h.sessionService.GetSessionByToken(cookie.Value)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Session not found"})
	}

	err = h.sessionService.DeleteSession(session.ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete session"})
	}

	utils.ClearCookie(c, "token")

	return c.JSON(http.StatusOK, echo.Map{"message": "Logout successful"})
}

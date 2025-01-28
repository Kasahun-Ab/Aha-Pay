package handlers

// import (
// 	"go_ecommerce/internal/services"
// 	"go_ecommerce/pkg/dto"
// 	"net/http"
// 	"strconv"

// 	"github.com/labstack/echo/v4"
// )

// type UserSessionHandler struct {
// 	service services.UserSessionService
// }

// func NewUserSessionHandler(service services.UserSessionService) *UserSessionHandler {
// 	return &UserSessionHandler{service: service}
// }

// func (h *UserSessionHandler) CreateSession(c echo.Context) error {

// 	var input dto.CreateUserSessionDTO

// 	if err := c.Bind(&input); err != nil {

// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})

// 	}

// 	session, err := h.service.CreateSession(input)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 	}

// 	return c.JSON(http.StatusCreated, session)
// }

// func (h *UserSessionHandler) GetSessionByID(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
// 	}

// 	session, err := h.service.GetSessionByID(id)
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, map[string]string{"error": "Session not found"})
// 	}

// 	return c.JSON(http.StatusOK, session)
// }

// func (h *UserSessionHandler) GetAllSessions(c echo.Context) error {
// 	sessions, err := h.service.GetAllSessions()
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, sessions)
// }

// func (h *UserSessionHandler) UpdateSession(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
// 	}

// 	var input dto.UpdateUserSessionDTO
// 	if err := c.Bind(&input); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
// 	}

// 	if err := h.service.UpdateSession(id, input); err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 	}

// 	return c.NoContent(http.StatusOK)
// }

// func (h *UserSessionHandler) DeleteSession(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
// 	}

// 	if err := h.service.DeleteSession(id); err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 	}

// 	return c.NoContent(http.StatusOK)
// }

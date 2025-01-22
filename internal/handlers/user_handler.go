package handlers

import (
	"go_ecommerce/internal/services"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type UserAccountHandler struct {
	Service *services.UserAccountService
}

func NewUserAccountHandler(service *services.UserAccountService) *UserAccountHandler {

	return &UserAccountHandler{Service: service}

}

func (h *UserAccountHandler) GetUser(c echo.Context) error {

	id := c.QueryParam("id")
	userId, err := strconv.Atoi(id)

	if err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user id"})
	}

	user, err := h.Service.FindByID(userId)

	if err != nil {

		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserAccountHandler) UpdateUser(c echo.Context) error {

	id := c.QueryParam("id")

	userId, err := strconv.Atoi(id)

	if err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user id"})
	}

	user, err := h.Service.FindByID(userId)

	if err != nil {

		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	dto := new(UpdateUserDTO)

	if err := c.Bind(dto); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if dto.Username != "" {

		user.Username = dto.Username
	}

	if dto.FirstName != "" {

		user.FirstName = dto.FirstName
	}
	if dto.LastName != "" {

		user.LastName = dto.LastName
	}

	if dto.Status != "" {

		user.Status = dto.Status
	}

	if err := h.Service.Update(user); err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update user"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User updated successfully"})
}

func (h *UserAccountHandler) DeleteUser(c echo.Context) error {

	id := c.QueryParam("id")

	userId, err := strconv.Atoi(id)

	if err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user id"})
	}

	user, err := h.Service.FindByID(userId)

	if err != nil {

		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	if err := h.Service.DeleteUser(user); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

func (h *UserAccountHandler) GetUserByEmail(c echo.Context) error {

	type UpdateUserByEmailDTO struct {
		Email string `json:"email"`
	}

	dto := new(UpdateUserByEmailDTO)

	if err := c.Bind(&dto); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})

	}

	email := dto.Email

	user, err := h.Service.FindByEmail(email)

	if err != nil {

		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserAccountHandler) UpdateUserByEmail(c echo.Context) error {

	dto := new(UpdateUserDTO)

	if err := c.Bind(dto); err != nil || dto.Email == "" {

		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request: missing or invalid email"})
	}

	user, err := h.Service.FindByEmail(dto.Email)

	if err != nil {

		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}
	if dto.Username != "" {

		user.Username = dto.Username
	}

	if dto.FirstName != "" {

		user.FirstName = dto.FirstName
	}
	if dto.LastName != "" {

		user.LastName = dto.LastName
	}

	if dto.Status != "" {

		user.Status = dto.Status
	}

	if err := h.Service.Update(user); err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update user"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User updated successfully"})
}

type UpdateUserDTO struct {
	Email string `json:"email"`

	Username string `json:"username,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`

	Status string `json:"status,omitempty"`
}

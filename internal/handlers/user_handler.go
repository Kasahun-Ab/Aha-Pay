package handlers

import (
	"go_ecommerce/internal/services"

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

	id := c.Get("user").(int)

	user, err := h.Service.FindByID(id)

	if err != nil {

		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserAccountHandler) UpdateUser(c echo.Context) error {

	id := c.Get("user").(int)

	user, err := h.Service.FindByID(id)

	if err != nil {

		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	if err := c.Bind(user); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := h.Service.Update(user); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User updated successfully"})
}

func (h *UserAccountHandler) DeleteUser(c echo.Context) error {
 
	id := c.Get("user").(int)

	user, err := h.Service.FindByID(id)

	if err != nil {

		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	if err := h.Service.DeleteUser(user); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

func (h *UserAccountHandler) GetUserByEmail(c echo.Context) error {

	email := c.Param("email")

	user, err := h.Service.FindByEmail(email)

	if err != nil {

		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}
 
func (h *UserAccountHandler) UpdateUserByEmail(c echo.Context) error {
	
	email := c.Param("email")

	user, err := h.Service.FindByEmail(email)

	if err != nil {

		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	if err := c.Bind(user); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := h.Service.Update(user); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User updated successfully"})
}
  
 

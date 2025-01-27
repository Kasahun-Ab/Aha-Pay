package handlers

import (
	"go_ecommerce/internal/middleware"
	"go_ecommerce/internal/models"
	"go_ecommerce/internal/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type WalletHandler struct {
	service *services.WalletService
}

func NewWalletHandler(service *services.WalletService) *WalletHandler {

	return &WalletHandler{service: service}

}

func (h *WalletHandler) CreateWallet(c echo.Context) error {

	userID, _ := middleware.GetUserID(c)

	wallet := new(models.Wallet)

	if err := c.Bind(wallet); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request"})

	}

	if _, err := h.service.CreateWallet(wallet, userID); err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}

	return c.JSON(http.StatusCreated, wallet)
}

func (h *WalletHandler) GetWalletByID(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid wallet ID"})

	}

	wallet, err := h.service.GetWalletByID(id)

	if err != nil {

		return c.JSON(http.StatusNotFound, map[string]string{"message": "wallet not found"})

	}

	return c.JSON(http.StatusOK, wallet)
}

// func (h *WalletHandler) UpdateWallet(c echo.Context) error {

// 	id, err := strconv.Atoi(c.Param("id"))

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid wallet ID"})
// 	}

// 	wallet := new(models.Wallet)

// 	if err := c.Bind(wallet); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request"})
// 	}

// 	wallet.ID = id

// 	if _, err := h.service.UpdateWallet(wallet); err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to update wallet"})
// 	}

// 	return c.JSON(http.StatusOK, wallet)
// }

func (h *WalletHandler) DeleteWallet(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid wallet ID"})

	}

	if err := h.service.DeleteWallet(id); err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to delete wallet"})

	}

	return c.JSON(http.StatusOK, map[string]string{"message": "wallet deleted"})
}


func (h *WalletHandler) GetAllWalletsByUserID(c echo.Context) error {

	userID, err := middleware.GetUserID(c)

	if err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid wallet ID"})

	}

	wallet, err := h.service.GetAllWalletsByUserID(userID)

	if err != nil {

		return c.JSON(http.StatusNotFound, map[string]string{"message": "wallet not found"})

	}

	return c.JSON(http.StatusOK, wallet)
}

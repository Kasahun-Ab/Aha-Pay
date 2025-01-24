// handlers/transaction_handler.go
package handlers

import (
	"net/http"
	// "strconv"
	"go_ecommerce/internal/models"
	"go_ecommerce/internal/services"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	service services.TransactionService
}

func NewTransactionHandler(service services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

// Create a new transaction (POST /transactions)
func (h *TransactionHandler) Create(c echo.Context) error {
	var transaction models.Transaction
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Call the service to create the transaction
	if err := h.service.CreateWithTransaction(&transaction); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, transaction)
}

// Get a transaction by ID (GET /transactions/:id)
// func (h *TransactionHandler) GetByID(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
// 	}

// 	transaction, err := h.service.GetByID(id)
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, map[string]string{"error": "Transaction not found"})
// 	}

// 	return c.JSON(http.StatusOK, transaction)
// }

package dto


type CreateTransactionRequest struct {
	WalletID        int     `json:"wallet_id" binding:"required"` // Required to associate with a Wallet
	Amount          float64 `json:"amount" binding:"required"`    // Required amount
	TransactionType string  `json:"transaction_type" binding:"required,oneof=DEPOSIT WITHDRAWAL TRANSFER"` // Enum validation
}

type UpdateTransactionRequest struct {
	Status string `json:"status" binding:"required,oneof=PENDING COMPLETED FAILED"` // Required for status updates
}




type TransactionResponse struct {
	ID              int     `json:"id"`
	WalletID        int     `json:"wallet_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	Status          string  `json:"status"`
	CreatedAt       string  `json:"created_at"` // ISO8601 format
	Wallet          *WalletResponse `json:"wallet,omitempty"` // Include Wallet details if needed
}

type WalletResponse struct {
	ID      int    `json:"id"`
	Balance string `json:"balance"` // Example Wallet field, modify based on your Wallet model
}

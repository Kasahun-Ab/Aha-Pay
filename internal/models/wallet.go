package models

// Wallet model with belongs to relationship with User
type Wallet struct {
	ID        int `gorm:"primaryKey"`
	UserID    int
	Currency  string `json:"currency"`
	Balance   float64 `json:"balance"`
	Status    string
	CreatedAt string
}

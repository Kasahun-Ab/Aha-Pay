package models

// Wallet model with belongs to relationship with User
type Wallet struct {
	ID        int `gorm:"primaryKey"`
	Currency  string
	Balance   float64
	Status    string
	CreatedAt string
}

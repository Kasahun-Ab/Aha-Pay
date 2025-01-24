package models

type Wallet struct {
	ID          int `gorm:"primaryKey"`
	UserID      int
	Currency    string  `json:"currency"`
	Balance     float64 `json:"balance"`
	Status      string
	CreatedAt   string
	Transaction []Transaction `gorm:"foreignKey:WalletID"`
	Sender    []Transfer    `gorm:"foreignKey:SenderWalletID"`
	Reciver    []Transfer    `gorm:"foreignKey:ReceiverWalletID"`
}

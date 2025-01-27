package dto

import "time"

// UserSessionDTO represents the data for transferring UserSession information
type UserSessionDTO struct {
	ID           int       `json:"id"`
	SessionToken string    `json:"session_token"`
	IPAddress    string    `json:"ip_address"`
	DeviceInfo   string    `json:"device_info"`
	CreatedAt    time.Time `json:"created_at"`
	LastActivity time.Time `json:"last_activity"`
	IsActive     bool      `json:"is_active"`
}

// CreateUserSessionDTO is used for creating a new UserSession
type CreateUserSessionDTO struct {
	SessionToken string `json:"session_token" validate:"required"`
	IPAddress    string `json:"ip_address" validate:"required"`
	DeviceInfo   string `json:"device_info" validate:"required"`
}

// UpdateUserSessionDTO is used for updating an existing UserSession
type UpdateUserSessionDTO struct {
	LastActivity time.Time `json:"last_activity" validate:"required"`
	IsActive     *bool     `json:"is_active"`
}

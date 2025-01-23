package dto

type FindUserByEmailRequestDTO struct {
	Email string `json:"email" validate:"required,email"` 
}

type FindUserByEmailResponseDTO struct {
	ID        string `json:"id"`         // Unique identifier for the user
	Email     string `json:"email"`      // Email address of the user
	Username  string `json:"username"`   // Full name of the user
	FirstName string `json:"first_name"` // Role of the user (e.g., admin, user, etc.)
	LastName  string `json:"last_name"`  // Date and time when the user was created (ISO format)
	Status    string `json:"status"`     // Date and time when the user was last updated (ISO format)
}



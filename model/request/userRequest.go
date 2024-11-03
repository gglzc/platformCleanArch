package request

import "github.com/google/uuid"

type CreateUserRequest struct {
	UserID   uuid.UUID `json:"name" validate:"required"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required,min=6"`
}

type UpdateUserRequest struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type UpdateBalance struct {
	UserID    string `json:"userId,omitempty"`
	IntoMoney int64  `json:"intoMoney,omitempty"`
}

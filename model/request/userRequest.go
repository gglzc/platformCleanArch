package request

import "github.com/google/uuid"

type CreateUserRequest struct {
	UserID   uuid.UUID `json:"userID" validate:"required"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required,min=6"`
}
type GetUserRequest struct {
	UserID uuid.UUID `json:"name" validate:"required"`
}

type UpdateUserRequest struct {
	UserID uuid.UUID `json:"userID,omitempty"`
	Email  string    `json:"email,omitempty"`
}

type UpdateBalance struct {
	UserID    uuid.UUID `json:"userID,omitempty"`
	IntoMoney int64     `json:"intoMoney,omitempty"`
}

package dto

import (
	"time"

	"github.com/syaafiudinm/go-starter-kit/internal/model"
)

// CreateUserRequest represents the payload for creating a new user.
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required,min=2,max=100"`
	Email string `json:"email" binding:"required,email"`
	Phone string `json:"phone" binding:"omitempty,max=20"`
}

// UpdateUserRequest represents the payload for updating user info.
type UpdateUserRequest struct {
	Name  string `json:"name" binding:"omitempty,min=2,max=100"`
	Email string `json:"email" binding:"omitempty,email"`
	Phone string `json:"phone" binding:"omitempty,max=20"`
}

// UserResponse represents the user data returned in API responses.
type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PaginationMeta holds pagination metadata for list responses.
type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	TotalPages  int64 `json:"total_pages"`
}

// ToUserResponse converts a User model to a UserResponse DTO.
func ToUserResponse(user *model.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

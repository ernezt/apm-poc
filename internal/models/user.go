package models

import (
	"time"
)

// UserRole represents a user role in the system
type UserRole string

// User roles
const (
	RoleOrganizationAdmin           UserRole = "organization_admin"
	RoleApplicationPortfolioManager UserRole = "application_portfolio_manager"
	RoleStakeholder                 UserRole = "stakeholder"
)

// User represents a user in the system
type User struct {
	ID             string    `json:"id"`
	OrganizationID string    `json:"organization_id"`
	Email          string    `json:"email"`
	PasswordHash   string    `json:"-"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Role           UserRole  `json:"role"`
	AvatarURL      string    `json:"avatar_url,omitempty"`
	MFAEnabled     bool      `json:"mfa_enabled"`
	MFASecret      string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// CreateUserRequest represents the request to create a new user
type CreateUserRequest struct {
	Email     string   `json:"email" validate:"required,email"`
	Password  string   `json:"password" validate:"required,min=8"`
	FirstName string   `json:"first_name" validate:"required"`
	LastName  string   `json:"last_name" validate:"required"`
	Role      UserRole `json:"role" validate:"required,oneof=organization_admin application_portfolio_manager stakeholder"`
}

// UpdateUserRequest represents the request to update a user
type UpdateUserRequest struct {
	FirstName string   `json:"first_name,omitempty"`
	LastName  string   `json:"last_name,omitempty"`
	Role      UserRole `json:"role,omitempty" validate:"omitempty,oneof=organization_admin application_portfolio_manager stakeholder"`
	AvatarURL string   `json:"avatar_url,omitempty"`
}

// UserResponse represents the response when returning user data
type UserResponse struct {
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Role       UserRole  `json:"role"`
	AvatarURL  string    `json:"avatar_url,omitempty"`
	MFAEnabled bool      `json:"mfa_enabled"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// LoginRequest represents the request to login
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	MFACode  string `json:"mfa_code,omitempty"`
}

// LoginResponse represents the response when logging in
type LoginResponse struct {
	User         UserResponse `json:"user"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token,omitempty"`
	RequiresMFA  bool         `json:"requires_mfa,omitempty"`
}

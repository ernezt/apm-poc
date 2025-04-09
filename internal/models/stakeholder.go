package models

import (
	"time"
)

// Stakeholder represents a stakeholder in the system
type Stakeholder struct {
	ID         string    `json:"id"`
	ForeignKey string    `json:"foreign_key"`
	UserID     string    `json:"user_id"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// CreateStakeholderRequest represents the request to create a new stakeholder
type CreateStakeholderRequest struct {
	ForeignKey string `json:"foreign_key,omitempty"`
	UserID     string `json:"user_id" validate:"required"`
	Role       string `json:"role" validate:"required"`
}

// UpdateStakeholderRequest represents the request to update a stakeholder
type UpdateStakeholderRequest struct {
	ForeignKey string `json:"foreign_key,omitempty"`
	Role       string `json:"role,omitempty"`
}

// StakeholderResponse represents the response when returning stakeholder data
type StakeholderResponse struct {
	ID         string    `json:"id"`
	ForeignKey string    `json:"foreign_key,omitempty"`
	UserID     string    `json:"user_id"`
	User       User      `json:"user,omitempty"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

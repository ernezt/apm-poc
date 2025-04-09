package models

import (
	"time"
)

// Entity represents an entity in the system (company, vendor, etc.)
type Entity struct {
	ID          string    `json:"id"`
	DisplayName string    `json:"display_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateEntityRequest represents the request to create a new entity
type CreateEntityRequest struct {
	DisplayName string `json:"display_name" validate:"required"`
}

// UpdateEntityRequest represents the request to update an entity
type UpdateEntityRequest struct {
	DisplayName string `json:"display_name,omitempty"`
}

// EntityResponse represents the response when returning entity data
type EntityResponse struct {
	ID          string    `json:"id"`
	DisplayName string    `json:"display_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

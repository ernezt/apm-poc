package models

import (
	"time"
)

// FunctionalCategory represents a functional category of software in the system
type FunctionalCategory struct {
	ID             string    `json:"id"`
	CategoryName   string    `json:"category_name"`
	CategoryParent string    `json:"category_parent,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// CreateFunctionalCategoryRequest represents the request to create a new functional category
type CreateFunctionalCategoryRequest struct {
	CategoryName   string `json:"category_name" validate:"required"`
	CategoryParent string `json:"category_parent,omitempty"`
}

// UpdateFunctionalCategoryRequest represents the request to update a functional category
type UpdateFunctionalCategoryRequest struct {
	CategoryName   string `json:"category_name,omitempty"`
	CategoryParent string `json:"category_parent,omitempty"`
}

// FunctionalCategoryResponse represents the response when returning functional category data
type FunctionalCategoryResponse struct {
	ID             string                      `json:"id"`
	CategoryName   string                      `json:"category_name"`
	CategoryParent string                      `json:"category_parent,omitempty"`
	ParentCategory *FunctionalCategoryResponse `json:"parent_category,omitempty"`
	CreatedAt      time.Time                   `json:"created_at"`
	UpdatedAt      time.Time                   `json:"updated_at"`
}

// SoftwareToCategory represents the many-to-many relationship between software and categories
type SoftwareToCategory struct {
	SoftwareID           string    `json:"software_id"`
	FunctionalCategoryID string    `json:"functional_category_id"`
	CreatedAt            time.Time `json:"created_at"`
}

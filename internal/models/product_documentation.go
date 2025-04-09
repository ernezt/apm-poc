package models

import (
	"time"
)

// ProductDocumentation represents documentation for a software product
type ProductDocumentation struct {
	ID           string    `json:"id"`
	ForeignKey   string    `json:"foreign_key"`
	DocumentType string    `json:"document_type"`
	DocumentURL  string    `json:"document_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// CreateProductDocumentationRequest represents the request to create new product documentation
type CreateProductDocumentationRequest struct {
	ForeignKey   string `json:"foreign_key" validate:"required"`
	DocumentType string `json:"document_type" validate:"required"`
	DocumentURL  string `json:"document_url" validate:"required,url"`
}

// UpdateProductDocumentationRequest represents the request to update product documentation
type UpdateProductDocumentationRequest struct {
	DocumentType string `json:"document_type,omitempty"`
	DocumentURL  string `json:"document_url,omitempty" validate:"omitempty,url"`
}

// ProductDocumentationResponse represents the response when returning product documentation data
type ProductDocumentationResponse struct {
	ID           string    `json:"id"`
	ForeignKey   string    `json:"foreign_key"`
	DocumentType string    `json:"document_type"`
	DocumentURL  string    `json:"document_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

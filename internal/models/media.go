package models

import (
	"time"
)

// Media represents a media asset in the system
type Media struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	MediaURL    string    `json:"media_url"`
	ExternalURL string    `json:"external_url,omitempty"`
	SourceName  string    `json:"source_name,omitempty"`
	SourceURL   string    `json:"source_url,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateMediaRequest represents the request to create a new media asset
type CreateMediaRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	MediaURL    string `json:"media_url" validate:"required"`
	ExternalURL string `json:"external_url,omitempty" validate:"omitempty,url"`
	SourceName  string `json:"source_name,omitempty"`
	SourceURL   string `json:"source_url,omitempty" validate:"omitempty,url"`
}

// UpdateMediaRequest represents the request to update a media asset
type UpdateMediaRequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	MediaURL    string `json:"media_url,omitempty"`
	ExternalURL string `json:"external_url,omitempty" validate:"omitempty,url"`
	SourceName  string `json:"source_name,omitempty"`
	SourceURL   string `json:"source_url,omitempty" validate:"omitempty,url"`
}

// MediaResponse represents the response when returning media data
type MediaResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	MediaURL    string    `json:"media_url"`
	ExternalURL string    `json:"external_url,omitempty"`
	SourceName  string    `json:"source_name,omitempty"`
	SourceURL   string    `json:"source_url,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

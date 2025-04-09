package models

import (
	"time"
)

// NewsArticle represents a news article in the system
type NewsArticle struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	MediaID     string    `json:"media_id,omitempty"`
	ExternalURL string    `json:"external_url,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateNewsArticleRequest represents the request to create a new news article
type CreateNewsArticleRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	MediaID     string `json:"media_id,omitempty"`
	ExternalURL string `json:"external_url,omitempty" validate:"omitempty,url"`
}

// UpdateNewsArticleRequest represents the request to update a news article
type UpdateNewsArticleRequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	MediaID     string `json:"media_id,omitempty"`
	ExternalURL string `json:"external_url,omitempty" validate:"omitempty,url"`
}

// NewsArticleResponse represents the response when returning news article data
type NewsArticleResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Media       *Media    `json:"media,omitempty"`
	ExternalURL string    `json:"external_url,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// SoftwareToNews represents the many-to-many relationship between software and news articles
type SoftwareToNews struct {
	SoftwareID    string    `json:"software_id"`
	NewsArticleID string    `json:"news_article_id"`
	CreatedAt     time.Time `json:"created_at"`
}

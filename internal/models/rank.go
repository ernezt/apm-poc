package models

import (
	"time"
)

// Rank represents a ranking or external status data for a software application
type Rank struct {
	ID            string    `json:"id"`
	SourceLink    string    `json:"source_link"`
	SourceName    string    `json:"source_name"`
	AverageScore  float64   `json:"average_score,omitempty"`
	NumReviews    int       `json:"number_of_reviews,omitempty"`
	LastUpdatedOn time.Time `json:"last_updated_on,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// CreateRankRequest represents the request to create a new rank
type CreateRankRequest struct {
	SourceLink    string    `json:"source_link" validate:"required"`
	SourceName    string    `json:"source_name" validate:"required"`
	AverageScore  float64   `json:"average_score,omitempty"`
	NumReviews    int       `json:"number_of_reviews,omitempty"`
	LastUpdatedOn time.Time `json:"last_updated_on,omitempty"`
}

// UpdateRankRequest represents the request to update a rank
type UpdateRankRequest struct {
	SourceLink    string    `json:"source_link,omitempty"`
	SourceName    string    `json:"source_name,omitempty"`
	AverageScore  float64   `json:"average_score,omitempty"`
	NumReviews    int       `json:"number_of_reviews,omitempty"`
	LastUpdatedOn time.Time `json:"last_updated_on,omitempty"`
}

// RankResponse represents the response when returning rank data
type RankResponse struct {
	ID            string     `json:"id"`
	SourceLink    string     `json:"source_link"`
	SourceName    string     `json:"source_name"`
	AverageScore  float64    `json:"average_score,omitempty"`
	NumReviews    int        `json:"number_of_reviews,omitempty"`
	LastUpdatedOn *time.Time `json:"last_updated_on,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

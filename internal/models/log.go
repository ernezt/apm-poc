package models

import (
	"time"
)

// Log represents a log entry in the system
type Log struct {
	ID         string    `json:"id"`
	ForeignKey string    `json:"foreign_key"`
	Context    string    `json:"context"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
}

// CreateLogRequest represents the request to create a new log entry
type CreateLogRequest struct {
	ForeignKey string `json:"foreign_key" validate:"required"`
	Context    string `json:"context" validate:"required"`
	Type       string `json:"type" validate:"required"`
}

// LogResponse represents the response when returning log data
type LogResponse struct {
	ID         string    `json:"id"`
	ForeignKey string    `json:"foreign_key"`
	Context    string    `json:"context"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
}

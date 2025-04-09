package models

import (
	"time"
)

// Status represents a status type/definition in the system
type Status struct {
	ID          string    `json:"id"`
	StatusType  string    `json:"status_type"`
	StatusName  string    `json:"status_name"`
	ActiveStart time.Time `json:"active_start,omitempty"`
	ActiveEnd   time.Time `json:"active_end,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateStatusRequest represents the request to create a new status
type CreateStatusRequest struct {
	StatusType  string    `json:"status_type" validate:"required"`
	StatusName  string    `json:"status_name" validate:"required"`
	ActiveStart time.Time `json:"active_start,omitempty"`
	ActiveEnd   time.Time `json:"active_end,omitempty"`
}

// UpdateStatusRequest represents the request to update a status
type UpdateStatusRequest struct {
	StatusType  string    `json:"status_type,omitempty"`
	StatusName  string    `json:"status_name,omitempty"`
	ActiveStart time.Time `json:"active_start,omitempty"`
	ActiveEnd   time.Time `json:"active_end,omitempty"`
}

// StatusResponse represents the response when returning status data
type StatusResponse struct {
	ID          string     `json:"id"`
	StatusType  string     `json:"status_type"`
	StatusName  string     `json:"status_name"`
	ActiveStart *time.Time `json:"active_start,omitempty"`
	ActiveEnd   *time.Time `json:"active_end,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// StatusLog represents a log entry of a status change in the system
type StatusLog struct {
	ID          string    `json:"id"`
	StatusID    string    `json:"status_id"`
	StatusOf    string    `json:"status_of"` // Entity ID the status is for
	StatusStart time.Time `json:"status_start"`
	StatusEnd   time.Time `json:"status_end,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateStatusLogRequest represents the request to create a new status log
type CreateStatusLogRequest struct {
	StatusID    string    `json:"status_id" validate:"required"`
	StatusOf    string    `json:"status_of" validate:"required"`
	StatusStart time.Time `json:"status_start" validate:"required"`
	StatusEnd   time.Time `json:"status_end,omitempty"`
}

// UpdateStatusLogRequest represents the request to update a status log
type UpdateStatusLogRequest struct {
	StatusStart time.Time `json:"status_start,omitempty"`
	StatusEnd   time.Time `json:"status_end,omitempty"`
}

// StatusLogResponse represents the response when returning status log data
type StatusLogResponse struct {
	ID          string     `json:"id"`
	StatusID    string     `json:"status_id"`
	Status      Status     `json:"status,omitempty"`
	StatusOf    string     `json:"status_of"`
	StatusStart time.Time  `json:"status_start"`
	StatusEnd   *time.Time `json:"status_end,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

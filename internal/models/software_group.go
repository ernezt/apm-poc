package models

import (
	"time"
)

// SoftwareGroup represents a group of related software applications
type SoftwareGroup struct {
	ID               string    `json:"id"`
	GroupName        string    `json:"group_name"`
	GroupDescription string    `json:"group_description"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// CreateSoftwareGroupRequest represents the request to create a new software group
type CreateSoftwareGroupRequest struct {
	GroupName        string `json:"group_name" validate:"required"`
	GroupDescription string `json:"group_description"`
}

// UpdateSoftwareGroupRequest represents the request to update a software group
type UpdateSoftwareGroupRequest struct {
	GroupName        string `json:"group_name,omitempty"`
	GroupDescription string `json:"group_description,omitempty"`
}

// SoftwareGroupResponse represents the response when returning software group data
type SoftwareGroupResponse struct {
	ID               string    `json:"id"`
	GroupName        string    `json:"group_name"`
	GroupDescription string    `json:"group_description,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// SoftwareToGroup represents the many-to-many relationship between software and groups
type SoftwareToGroup struct {
	SoftwareID      string    `json:"software_id"`
	SoftwareGroupID string    `json:"software_group_id"`
	CreatedAt       time.Time `json:"created_at"`
}

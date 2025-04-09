package models

import (
	"time"
)

// SoftwareType represents the type of software
type SoftwareType string

const (
	SoftwareTypeAPI        SoftwareType = "api"
	SoftwareTypeWeb        SoftwareType = "web"
	SoftwareTypeMobile     SoftwareType = "mobile"
	SoftwareTypeDesktop    SoftwareType = "desktop"
	SoftwareTypeEmbedded   SoftwareType = "embedded"
	SoftwareTypeMiddleware SoftwareType = "middleware"
	SoftwareTypeLibrary    SoftwareType = "library"
)

// Software represents a software application in the system
type Software struct {
	ID                   string       `json:"id"`
	ForeignKey           string       `json:"foreign_key"`
	DisplayName          string       `json:"display_name"`
	Description          string       `json:"description"`
	SoftwareType         SoftwareType `json:"software_type"`
	SoftwareSubtype      string       `json:"software_subtype"`
	Vendor               string       `json:"vendor"`
	Manufacturer         string       `json:"manufacturer"`
	InstallType          string       `json:"install_type"`
	ProductType          string       `json:"product_type"`
	Context              string       `json:"context"`
	LifecycleStatus      string       `json:"lifecycle_status"`
	ImplementationStatus string       `json:"implementation_status"`
	CreatedAt            time.Time    `json:"created_at"`
	UpdatedAt            time.Time    `json:"updated_at"`
}

// CreateSoftwareRequest represents the request to create new software
type CreateSoftwareRequest struct {
	ForeignKey           string       `json:"foreign_key,omitempty"`
	DisplayName          string       `json:"display_name" validate:"required"`
	Description          string       `json:"description"`
	SoftwareType         SoftwareType `json:"software_type" validate:"required,oneof=api web mobile desktop embedded middleware library"`
	SoftwareSubtype      string       `json:"software_subtype,omitempty"`
	Vendor               string       `json:"vendor,omitempty"`
	Manufacturer         string       `json:"manufacturer,omitempty"`
	InstallType          string       `json:"install_type,omitempty"`
	ProductType          string       `json:"product_type,omitempty"`
	Context              string       `json:"context,omitempty"`
	LifecycleStatus      string       `json:"lifecycle_status,omitempty"`
	ImplementationStatus string       `json:"implementation_status,omitempty"`
}

// UpdateSoftwareRequest represents the request to update software
type UpdateSoftwareRequest struct {
	ForeignKey           string       `json:"foreign_key,omitempty"`
	DisplayName          string       `json:"display_name,omitempty"`
	Description          string       `json:"description,omitempty"`
	SoftwareType         SoftwareType `json:"software_type,omitempty" validate:"omitempty,oneof=api web mobile desktop embedded middleware library"`
	SoftwareSubtype      string       `json:"software_subtype,omitempty"`
	Vendor               string       `json:"vendor,omitempty"`
	Manufacturer         string       `json:"manufacturer,omitempty"`
	InstallType          string       `json:"install_type,omitempty"`
	ProductType          string       `json:"product_type,omitempty"`
	Context              string       `json:"context,omitempty"`
	LifecycleStatus      string       `json:"lifecycle_status,omitempty"`
	ImplementationStatus string       `json:"implementation_status,omitempty"`
}

// SoftwareResponse represents the response when returning software data
type SoftwareResponse struct {
	ID                   string       `json:"id"`
	ForeignKey           string       `json:"foreign_key,omitempty"`
	DisplayName          string       `json:"display_name"`
	Description          string       `json:"description,omitempty"`
	SoftwareType         SoftwareType `json:"software_type"`
	SoftwareSubtype      string       `json:"software_subtype,omitempty"`
	Vendor               string       `json:"vendor,omitempty"`
	Manufacturer         string       `json:"manufacturer,omitempty"`
	InstallType          string       `json:"install_type,omitempty"`
	ProductType          string       `json:"product_type,omitempty"`
	Context              string       `json:"context,omitempty"`
	LifecycleStatus      string       `json:"lifecycle_status,omitempty"`
	ImplementationStatus string       `json:"implementation_status,omitempty"`
	CreatedAt            time.Time    `json:"created_at"`
	UpdatedAt            time.Time    `json:"updated_at"`
}

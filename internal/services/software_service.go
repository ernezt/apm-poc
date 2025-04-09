package services

import (
	"context"
	"fmt"
	"log"

	"apm/internal/db/repository"
	"apm/internal/models"
)

// Ensure implementation satisfies the interface
var _ SoftwareService = (*softwareService)(nil)

// softwareService implements SoftwareService
type softwareService struct {
	repo   repository.SoftwareRepository
	logger *log.Logger
}

// NewSoftwareService creates a new software service
func NewSoftwareService(repo repository.SoftwareRepository, logger *log.Logger) SoftwareService {
	return &softwareService{
		repo:   repo,
		logger: logger,
	}
}

// Create creates a new software entity
func (s *softwareService) Create(ctx context.Context, req models.CreateSoftwareRequest) (models.SoftwareResponse, error) {
	s.logger.Println("Creating new software:", req.DisplayName)

	// Convert request to Software model
	software := models.Software{
		ForeignKey:           req.ForeignKey,
		DisplayName:          req.DisplayName,
		Description:          req.Description,
		SoftwareType:         req.SoftwareType,
		SoftwareSubtype:      req.SoftwareSubtype,
		Vendor:               req.Vendor,
		Manufacturer:         req.Manufacturer,
		InstallType:          req.InstallType,
		ProductType:          req.ProductType,
		Context:              req.Context,
		LifecycleStatus:      req.LifecycleStatus,
		ImplementationStatus: req.ImplementationStatus,
	}

	// Create the software entity using the repository
	createdSoftware, err := s.repo.Create(ctx, software)
	if err != nil {
		s.logger.Printf("Error creating software: %v", err)
		return models.SoftwareResponse{}, fmt.Errorf("failed to create software: %w", err)
	}

	// Convert created software to response model
	return s.mapSoftwareToResponse(createdSoftware), nil
}

// GetByID retrieves a software entity by ID
func (s *softwareService) GetByID(ctx context.Context, id string) (models.SoftwareResponse, error) {
	s.logger.Println("Getting software by ID:", id)

	software, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Printf("Error getting software by ID: %v", err)
		return models.SoftwareResponse{}, fmt.Errorf("failed to get software: %w", err)
	}

	return s.mapSoftwareToResponse(software), nil
}

// List retrieves a list of software entities with pagination
func (s *softwareService) List(ctx context.Context, limit, offset int) ([]models.SoftwareResponse, error) {
	s.logger.Printf("Listing software (limit: %d, offset: %d)", limit, offset)

	// Validate pagination parameters
	if limit <= 0 {
		limit = 10 // Default limit
	}
	if limit > 100 {
		limit = 100 // Max limit
	}
	if offset < 0 {
		offset = 0
	}

	// Get software list from repository
	softwareList, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		s.logger.Printf("Error listing software: %v", err)
		return nil, fmt.Errorf("failed to list software: %w", err)
	}

	// Map software entities to response models
	var responseList []models.SoftwareResponse
	for _, software := range softwareList {
		responseList = append(responseList, s.mapSoftwareToResponse(software))
	}

	return responseList, nil
}

// Update updates an existing software entity
func (s *softwareService) Update(ctx context.Context, id string, req models.UpdateSoftwareRequest) error {
	s.logger.Println("Updating software with ID:", id)

	// Get the existing software
	existingSoftware, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Printf("Error getting software to update: %v", err)
		return fmt.Errorf("failed to get software for update: %w", err)
	}

	// Update fields if they are provided
	if req.DisplayName != "" {
		existingSoftware.DisplayName = req.DisplayName
	}
	if req.Description != "" {
		existingSoftware.Description = req.Description
	}
	if req.SoftwareType != "" {
		existingSoftware.SoftwareType = req.SoftwareType
	}
	if req.SoftwareSubtype != "" {
		existingSoftware.SoftwareSubtype = req.SoftwareSubtype
	}
	if req.Vendor != "" {
		existingSoftware.Vendor = req.Vendor
	}
	if req.Manufacturer != "" {
		existingSoftware.Manufacturer = req.Manufacturer
	}
	if req.InstallType != "" {
		existingSoftware.InstallType = req.InstallType
	}
	if req.ProductType != "" {
		existingSoftware.ProductType = req.ProductType
	}
	if req.Context != "" {
		existingSoftware.Context = req.Context
	}
	if req.LifecycleStatus != "" {
		existingSoftware.LifecycleStatus = req.LifecycleStatus
	}
	if req.ImplementationStatus != "" {
		existingSoftware.ImplementationStatus = req.ImplementationStatus
	}
	if req.ForeignKey != "" {
		existingSoftware.ForeignKey = req.ForeignKey
	}

	// Update the software entity using the repository
	err = s.repo.Update(ctx, existingSoftware)
	if err != nil {
		s.logger.Printf("Error updating software: %v", err)
		return fmt.Errorf("failed to update software: %w", err)
	}

	return nil
}

// Delete removes a software entity
func (s *softwareService) Delete(ctx context.Context, id string) error {
	s.logger.Println("Deleting software with ID:", id)

	// Delete the software entity using the repository
	err := s.repo.Delete(ctx, id)
	if err != nil {
		s.logger.Printf("Error deleting software: %v", err)
		return fmt.Errorf("failed to delete software: %w", err)
	}

	return nil
}

// Helper function to map Software to SoftwareResponse
func (s *softwareService) mapSoftwareToResponse(software models.Software) models.SoftwareResponse {
	return models.SoftwareResponse{
		ID:                   software.ID,
		ForeignKey:           software.ForeignKey,
		DisplayName:          software.DisplayName,
		Description:          software.Description,
		SoftwareType:         software.SoftwareType,
		SoftwareSubtype:      software.SoftwareSubtype,
		Vendor:               software.Vendor,
		Manufacturer:         software.Manufacturer,
		InstallType:          software.InstallType,
		ProductType:          software.ProductType,
		Context:              software.Context,
		LifecycleStatus:      software.LifecycleStatus,
		ImplementationStatus: software.ImplementationStatus,
		CreatedAt:            software.CreatedAt,
		UpdatedAt:            software.UpdatedAt,
	}
}

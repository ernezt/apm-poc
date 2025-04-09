package repository

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"apm/internal/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Ensure implementation satisfies the interface
var _ SoftwareRepository = (*PostgresSoftwareRepository)(nil)

// PostgresSoftwareRepository implements SoftwareRepository using PostgreSQL
type PostgresSoftwareRepository struct {
	pool   *pgxpool.Pool
	logger *log.Logger
}

// NewPostgresSoftwareRepository creates a new PostgreSQL software repository
func NewPostgresSoftwareRepository(pool *pgxpool.Pool) SoftwareRepository {
	return &PostgresSoftwareRepository{
		pool:   pool,
		logger: log.New(log.Writer(), "[SoftwareRepo] ", log.LstdFlags),
	}
}

// generateID creates a random UUID-like string
func generateID() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return fmt.Sprintf("fallback-id-%d", time.Now().UnixNano())
	}
	return hex.EncodeToString(bytes)
}

// Create inserts a new software record into the database
func (r *PostgresSoftwareRepository) Create(ctx context.Context, software models.Software) (models.Software, error) {
	// Generate a new ID if not provided
	if software.ID == "" {
		software.ID = generateID()
	}

	// Set timestamps
	now := time.Now().UTC()
	software.CreatedAt = now
	software.UpdatedAt = now

	// SQL query to insert a new software record
	query := `
		INSERT INTO software (
			id, foreign_key, display_name, description, software_type, 
			software_subtype, vendor, manufacturer, install_type, 
			product_type, context, lifecycle_status, implementation_status,
			created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
		) RETURNING *
	`

	// Execute the query
	row := r.pool.QueryRow(ctx, query,
		software.ID, software.ForeignKey, software.DisplayName, software.Description,
		software.SoftwareType, software.SoftwareSubtype, software.Vendor,
		software.Manufacturer, software.InstallType, software.ProductType,
		software.Context, software.LifecycleStatus, software.ImplementationStatus,
		software.CreatedAt, software.UpdatedAt,
	)

	// Scan the result into the software struct
	var result models.Software
	err := row.Scan(
		&result.ID, &result.ForeignKey, &result.DisplayName, &result.Description,
		&result.SoftwareType, &result.SoftwareSubtype, &result.Vendor,
		&result.Manufacturer, &result.InstallType, &result.ProductType,
		&result.Context, &result.LifecycleStatus, &result.ImplementationStatus,
		&result.CreatedAt, &result.UpdatedAt,
	)
	if err != nil {
		return models.Software{}, fmt.Errorf("failed to create software: %w", err)
	}

	return result, nil
}

// GetByID retrieves a software record by its ID
func (r *PostgresSoftwareRepository) GetByID(ctx context.Context, id string) (models.Software, error) {
	query := `SELECT * FROM software WHERE id = $1`
	row := r.pool.QueryRow(ctx, query, id)

	var software models.Software
	err := row.Scan(
		&software.ID, &software.ForeignKey, &software.DisplayName, &software.Description,
		&software.SoftwareType, &software.SoftwareSubtype, &software.Vendor,
		&software.Manufacturer, &software.InstallType, &software.ProductType,
		&software.Context, &software.LifecycleStatus, &software.ImplementationStatus,
		&software.CreatedAt, &software.UpdatedAt,
	)
	if err != nil {
		return models.Software{}, fmt.Errorf("failed to get software by ID: %w", err)
	}

	return software, nil
}

// List retrieves a list of software records with pagination
func (r *PostgresSoftwareRepository) List(ctx context.Context, limit, offset int) ([]models.Software, error) {
	query := `SELECT * FROM software ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list software: %w", err)
	}
	defer rows.Close()

	var softwareList []models.Software
	for rows.Next() {
		var software models.Software
		err := rows.Scan(
			&software.ID, &software.ForeignKey, &software.DisplayName, &software.Description,
			&software.SoftwareType, &software.SoftwareSubtype, &software.Vendor,
			&software.Manufacturer, &software.InstallType, &software.ProductType,
			&software.Context, &software.LifecycleStatus, &software.ImplementationStatus,
			&software.CreatedAt, &software.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan software: %w", err)
		}
		softwareList = append(softwareList, software)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return softwareList, nil
}

// Update updates an existing software record
func (r *PostgresSoftwareRepository) Update(ctx context.Context, software models.Software) error {
	// Update the UpdatedAt timestamp
	software.UpdatedAt = time.Now().UTC()

	query := `
		UPDATE software SET
			foreign_key = $2,
			display_name = $3,
			description = $4,
			software_type = $5,
			software_subtype = $6,
			vendor = $7,
			manufacturer = $8,
			install_type = $9,
			product_type = $10,
			context = $11,
			lifecycle_status = $12,
			implementation_status = $13,
			updated_at = $14
		WHERE id = $1
	`

	_, err := r.pool.Exec(ctx, query,
		software.ID, software.ForeignKey, software.DisplayName, software.Description,
		software.SoftwareType, software.SoftwareSubtype, software.Vendor,
		software.Manufacturer, software.InstallType, software.ProductType,
		software.Context, software.LifecycleStatus, software.ImplementationStatus,
		software.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update software: %w", err)
	}

	return nil
}

// Delete removes a software record by its ID
func (r *PostgresSoftwareRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM software WHERE id = $1`
	_, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete software: %w", err)
	}

	return nil
}

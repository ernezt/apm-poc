# Application Portfolio Management (APM)

A system to keep track of an organization's application portfolio, providing insights into the number, value, technologies, and lifecycle of applications.

## Features

- Import software from CSV files
- Pluggable classification engine with internet data enrichment
- User-friendly interface with role-based access control
- Master application database with organization-specific customizations
- News and ranking integrations for applications and entities

## Tech Stack

- **Backend**: Go
- **Database**: PostgreSQL
- **ORM**: SQLC
- **Frontend**: HTMX + Alpine.js
- **CSS Framework**: Tailwind CSS + DaisyUI
- **Deployment**: Kubernetes with Kustomize and Flux

## Getting Started

### Prerequisites

- Go 1.22 or higher
- Docker and Docker Compose
- Make (optional, for convenience commands)

### Setup

1. Clone the repository
   ```
   git clone <repository-url>
   cd apm
   ```

2. Start the database
   ```
   make db-up
   ```

3. Run the application
   ```
   make run
   ```

4. Access the application at http://localhost:8080

### Development

- `make build` - Build the application
- `make run` - Run the application locally
- `make test` - Run tests
- `make db-up` - Start the database
- `make db-down` - Stop the database
- `make migrate` - Run database migrations
- `make clean` - Clean build artifacts

## Project Structure

```
apm/
├── cmd/              # Command-line applications
├── config/           # Configuration files
├── internal/         # Private application code
│   ├── api/          # HTTP server and handlers
│   ├── db/           # Database access layer
│   ├── models/       # Domain models
│   ├── handlers/     # HTTP handlers
│   ├── services/     # Business logic
│   └── utils/        # Utility functions
├── migrations/       # Database migration files
├── static/           # Static web assets
├── templates/        # HTML templates
└── tests/            # Test files
```

## Database Schema

The database schema is defined in `migrations/1.sql` and includes tables for:

- Organizations
- Users
- Applications (master and organization-specific)
- Categories
- Software types
- Entities (master and organization-specific)
- Application clusters
- News articles
- Rankings

## License

This project is licensed under the MIT License - see the LICENSE file for details. 
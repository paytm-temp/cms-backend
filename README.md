# Case Management System Backend

This repository contains the backend implementation for the Case Management System.

## Features

- Role-based access control (RBAC)
- Case management endpoints
- Mock data store for testing
- RESTful API design

## Setup

1. Install Go (version 1.21 or higher)
2. Clone the repository
3. Run `go mod tidy` to install dependencies
4. Start the server with `go run cmd/main/main.go`

## API Documentation

The API provides the following endpoints:

- `GET /api/v1/cases` - Get all cases
- `GET /api/v1/cases/:id` - Get a specific case
- `POST /api/v1/cases` - Create a new case
- `PUT /api/v1/cases/:id` - Update a case
- `DELETE /api/v1/cases/:id` - Delete a case
- `GET /api/v1/cases/stats` - Get case statistics

## Role-Based Access

The API uses the `X-User-Role` header for role-based access control:

- REQUESTOR
- CASE_MANAGER
- RESOLVER

## Development

The project uses Go modules for dependency management. To add new dependencies:

```bash
go get github.com/username/package
```

## Testing

Run tests with:

```bash
go test ./...
```

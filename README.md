# Go Backend Boilerplate

A lightweight and feature-rich Go backend boilerplate that provides a solid foundation for building REST APIs. This boilerplate includes structured routing, middleware support, logging, CORS configuration, and authentication scaffolding.

## Features

- **Structured REST API** with versioned routes
- **Multi-router support** for organizing different API sections
- **Built-in middleware** for logging, tracing, and authentication
- **CORS configuration** with customizable settings
- **Health check endpoint** for monitoring
- **Request tracing** with unique trace IDs
- **Configurable logging** with different log levels
- **Clean project structure** for scalability

## Built With

This boilerplate leverages the following custom libraries:

- **[github.com/phasi/go-restapi](https://github.com/phasi/go-restapi)** - Lightweight REST API framework with routing, middleware, and CORS support
- **[github.com/phasi/go-logs](https://github.com/phasi/go-logs)** - Simple and efficient logging library

## Getting Started

### Prerequisites

- Go 1.22.6 or later

### Installation

1. Clone this repository:

```bash
git clone <your-repo-url>
cd go-backend-boilerplate
```

2. Install dependencies:

```bash
go mod tidy
```

3. Run the server:

```bash
go run .
```

The server will start on `http://localhost:8080`

## Project Structure

```
├── main.go           # Application entry point and server setup
├── routerv1.go       # V1 API routes and handlers
├── healthcheck.go    # Health check endpoint
├── models.go         # Data models and structures
├── permissions.go    # Permission and authorization logic
├── go.mod           # Go module dependencies
└── go.sum           # Dependency checksums
```

## API Endpoints

### Health Check

- `GET /api/health/` - Returns server status and timestamp

### V1 API Routes

Base path: `/api/v1`

- Example user endpoints (customize as needed)
- Authentication middleware enabled for protected routes
- CORS configured for `http://localhost:3000`

## Configuration

### CORS Settings

CORS is configured in `routerv1.go` with the following default settings:

- Allowed Origins: `http://localhost:3000`
- Allowed Methods: `GET`, `POST`, `PUT`, `DELETE`
- Allowed Headers: `Content-Type`, `Authorization`
- Credentials: Enabled
- Max Age: 3600 seconds

### Logging

- Configurable log levels (INFO, DEBUG, etc.)
- HTTP request logging with trace IDs
- Header redaction for sensitive information (Authorization headers)

### Authentication

Basic authentication middleware is included as a starting point. Customize the token validation logic in `routerv1.go` to match your authentication requirements.

## Database Integration

For database operations, consider using a PostgreSQL ORM library:
**[github.com/phasi/go-postgresql-orm](https://github.com/phasi/go-postgresql-orm)** - A simple and efficient PostgreSQL ORM for Go applications.

## Development

### Adding New Routes

1. Create handlers in the appropriate router file
2. Add route definitions to the router
3. Apply necessary middleware (authentication, etc.)

### Extending the API

- Add new router versions by creating additional router files
- Register new routers in the `main.go` file's multi-router setup
- Maintain backward compatibility for existing API versions

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE).

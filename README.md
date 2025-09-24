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
git clone https://github.com/phasi/go-backend-boilerplate.git
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
├── models.go         # Define your models
├── permissions.go    # Define your permissions
├── go.mod           # Go module dependencies
└── go.sum           # Dependency checksums
```

## API Endpoints

### Health Check

- `GET /api/health` - Returns server status and timestamp

### V1 API Routes

Base path: `/api/v1`

Contains already:

- Example endpoints (customize as needed)
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
- HTTP request logging with trace IDs (for logging analytics)
- Header redaction for sensitive information (Authorization headers)
- Log format is JSON

### Authentication

Basic authentication middleware is included as a starting point. Customize the token validation logic in `routerv1.go` to match your authentication requirements.

## Database Integration

For database operations, here's a suggestion:
**[github.com/phasi/go-postgresql-orm](https://github.com/phasi/go-postgresql-orm)** - A simple PostgreSQL ORM.

## Development

### Adding New Routes

1. Create handlers in the appropriate router file, _example_:

```go
    // routerv1.go

    func getRouterV1() *api.Router {

	// Create router
	router := &api.Router{
		BasePath:   "/v1",
	}

	// Public routes (handlers)
	router.HandleFunc("GET", "/users/:id", func(w http.ResponseWriter, r *http.Request, ctx *api.RouteContext) {
		id, err := ctx.Params.Get("id")
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "User ID must be an integer", http.StatusBadRequest)
			return
		}
		user := User{ID: idInt, Name: "John Doe", Email: "john@example.com"}
		api.WriteJSON(w, user)
	})
	// Protected routes (handlers)
	router.HandleProtectedFunc("GET", "/admin/users",
		[]api.Permission{PermissionViewUsers},
		func(w http.ResponseWriter, r *http.Request, ctx *api.RouteContext) {
			users := []User{
				{ID: 1, Name: "John Doe", Email: "john@example.com"},
				{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
			}
			api.WriteJSON(w, users)
		})
    }
```

2. Apply necessary middleware (authentication, etc.)

```go
	// Set up authentication middleware
	router.AuthorizationMiddleware = func(context *api.RouteContext, handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			// change below token validation to your own logic
			if token == "Bearer valid-token" {
				context.SetUserId("user-123")
				handler.ServeHTTP(w, r)
			} else {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}
		})
	}

```

```go

	// Set up permission middleware
	router.PermissionMiddleware = func(context *api.RouteContext, handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Mock user permissions (in real app, fetch from database/cache)
			userPermissions := []api.Permission{PermissionViewUsers, PermissionEditUsers, PermissionAdmin}
			if context.HasRequiredPermissions(userPermissions) {
				handler.ServeHTTP(w, r)
			} else {
				http.Error(w, "Forbidden", http.StatusForbidden)
			}
		})
	}

```

Edit _permissions.go_ to add your permissions:

```go

package main

import api "github.com/phasi/go-restapi"

const (
	PermissionViewUsers   api.Permission = 1
	PermissionEditUsers   api.Permission = 2
	PermissionDeleteUsers api.Permission = 3
	PermissionAdmin       api.Permission = 10
    // add more
)


```

### Extending the API

- Add new router versions by creating additional router files
- Register new routers in the `main.go` file's multi-router setup

```go
	// multirouter gathers all routers under a common base path
	multiRouter, err := api.NewMultiRouter("/api", []*api.Router{routerV1, healthcheckRouter})
	if err != nil {
		logger.Fatal("Failed to create multi-router: %v", err)
	}
```

Maintain backward compatibility for existing API versions when using versioned routers. If you do not want to do so, you can also use a regular router and keep it simple.

## License

This project is open source and available under the [MIT License](LICENSE).

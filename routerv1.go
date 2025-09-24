package main

import (
	"net/http"
	"strconv"

	api "github.com/phasi/go-restapi"
)

func getRouterV1() *api.Router {
	// Configure CORS
	corsConfig := &api.CORSConfig{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           3600,
	}

	// Create router
	router := &api.Router{
		BasePath:   "/v1",
		CORSConfig: corsConfig,
	}

	// Set up authentication middleware
	router.AuthorizationMiddleware = func(context *api.RouteContext, handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			// Mock token validation
			if token == "Bearer valid-token" {
				context.SetUserId("user-123")
				handler.ServeHTTP(w, r)
			} else {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}
		})
	}

	// Set up permission middleware
	router.PermissionMiddleware = func(context *api.RouteContext, handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Mock user permissions (in real app, fetch from database)
			userPermissions := []api.Permission{PermissionViewUsers, PermissionEditUsers, PermissionAdmin}
			if context.HasRequiredPermissions(userPermissions) {
				handler.ServeHTTP(w, r)
			} else {
				http.Error(w, "Forbidden", http.StatusForbidden)
			}
		})
	}

	// Public routes
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
	// Protected routes
	router.HandleProtectedFunc("GET", "/admin/users",
		[]api.Permission{PermissionViewUsers},
		func(w http.ResponseWriter, r *http.Request, ctx *api.RouteContext) {
			users := []User{
				{ID: 1, Name: "John Doe", Email: "john@example.com"},
				{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
			}
			api.WriteJSON(w, users)
		})

	return router
}

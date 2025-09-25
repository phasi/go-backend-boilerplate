package main

import (
	"net/http"
	"os"

	gologs "github.com/phasi/go-logs"
	api "github.com/phasi/go-restapi"
)

var logger *gologs.Logger

func init() {
	logger = gologs.NewLogger(gologs.INFO, os.Stdout)
	// Enable CORS even though Origin header would be missing, uncomment below line
	//api.SetCORSAlwaysOn(true)
}

func main() {
	// this router contains V1 routes, add more routes as needed
	routerV1 := getRouterV1()
	healthcheckRouter := getHealthcheckRouter()

	// multirouter gathers all routers under a common base path
	multiRouter, err := api.NewMultiRouter("/api", []*api.Router{routerV1, healthcheckRouter})
	if err != nil {
		logger.Fatal("Failed to create multi-router: %v", err)
	}
	// Set up logging
	api.SetRedactedHeaderNames([]string{"Authorization"})
	logFunc := func(entry api.HttpLogEntry) {
		logger.Info(
			"%s %s %d (trace: %s)",
			entry.Method,
			entry.Path,
			entry.Status,
			entry.TraceID,
		)
	}
	// Apply middleware for logging and tracing
	loggedRouter := api.LoggingRouter(multiRouter, logFunc)
	tracedRouter := api.TracingRouter(loggedRouter)
	// Start server
	logger.Info("Configured routes:")
	logger.Log(multiRouter.ListRoutes()).Info()
	logger.Info("Server starting on :8080")
	err = http.ListenAndServe(":8080", tracedRouter)
	if err != nil {
		logger.Fatal("Server failed: %v", err)
	}
}

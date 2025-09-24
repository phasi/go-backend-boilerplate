package main

import (
	"net/http"
	"time"

	api "github.com/phasi/go-restapi"
)

func getHealthcheckRouter() *api.Router {
	// Create router
	router := &api.Router{
		BasePath: "/health",
	}

	router.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request, ctx *api.RouteContext) {
		api.WriteJSON(w, map[string]string{"status": "ok", "time": time.Now().Format(time.RFC3339)})
	})

	return router
}

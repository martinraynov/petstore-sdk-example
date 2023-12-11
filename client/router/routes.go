package router

import (
	"net/http"

	handler "petstore_client/handlers"
)

// Route type description
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var baseRoutes = Routes{
	Route{
		"Status",
		"GET",
		"/status",
		handler.Status,
	},
}

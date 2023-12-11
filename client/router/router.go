package router

import (
	"github.com/gorilla/mux"
)

// NewRouter builds and returns a new router from routes
func NewRouter(debug bool) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// Base endpoints
	baseEndpoints := router.PathPrefix("/client/v1").Subrouter()
	for _, route := range baseRoutes {
		baseEndpoints.
			HandleFunc(route.Pattern, route.HandlerFunc).
			Name(route.Name).
			Methods(route.Method)
	}

	return router
}

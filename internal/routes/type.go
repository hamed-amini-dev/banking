package routes

import "net/http"

// Route - Defines the structure of a Route
type Route struct {
	Name   string
	Method string
	Path   string
	// Handler   func(r *request.GenericRequest) (interface{}, error)
	Handler func(http.ResponseWriter, *http.Request)
}

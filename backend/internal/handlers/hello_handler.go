// Package handlers contains the HTTP handlers for the application's API endpoints.
// They are responsible for parsing requests, calling services, and writing responses.
package handlers

import (
	"net/http"

	"github.com/FaraamFide/go-pocketbase-boilerplate/backend/internal/services"
	"github.com/pocketbase/pocketbase/core"
)

// HelloHandler handles requests for the hello endpoint.
// It depends on a HelloService to get the data.
type HelloHandler struct {
	service *services.HelloService
}

// NewHelloHandler creates a new instance of HelloHandler.
func NewHelloHandler(s *services.HelloService) *HelloHandler {
	return &HelloHandler{
		service: s,
	}
}

// Greet is the handler function for the GET /api/hello route.
// It extracts data from the request, calls the service, and writes a JSON response.
func (h *HelloHandler) Greet(c *core.RequestEvent) error {
	// Get a query parameter, e.g., /api/hello?name=Go
	name := c.Request.URL.Query().Get("name")

	// Call the service to get the business logic result.
	message := h.service.GetGreeting(name)
	
	return c.JSON(http.StatusOK, map[string]string{
		"message": message,
	})
}

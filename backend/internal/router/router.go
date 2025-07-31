// Package router registers the application's custom API routes.
package router

import (
	"log/slog"

	"github.com/FaraamFide/go-pocketbase-boilerplate/backend/internal/handlers"
	"github.com/FaraamFide/go-pocketbase-boilerplate/backend/internal/services"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

// RegisterRoutes attaches the application's custom API routes to the PocketBase instance.
// It uses the OnServe hook to ensure routes are added just before the server starts.
func RegisterRoutes(app *pocketbase.PocketBase) error {
	app.OnServe().BindFunc(func(e *core.ServeEvent) error {

		// Instantiate services
		helloSvc := services.NewHelloService()

		// Instantiate handlers, injecting dependencies
		helloHandler := handlers.NewHelloHandler(helloSvc)

		// Define API routes under the "/api" group
		api := e.Router.Group("/api")

		// Example: GET /api/hello?name=World
		api.GET("/hello", helloHandler.Greet)

		// To protect a route, you can add PocketBase's auth middleware.
		// For example:
		// api.GET("/protected/hello", helloHandler.Greet, apis.RequireRecordAuth())

		api.GET("/test", func(c *core.RequestEvent) error {
			return c.String(200, "PocketBase is running!\n")
		})

		return e.Next()
	})

	slog.Info("Custom routes registered successfully")
	return nil
}

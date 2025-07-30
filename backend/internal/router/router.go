package router

import (
	"github.com/FaraamFide/go-pocketbase-boilerplate/backend/internal/handlers"
	"github.com/FaraamFide/go-pocketbase-boilerplate/backend/internal/services"
	"log/slog"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

// RegisterRoutes registers all custom API routes for the application.
// Routes are attached to the `OnServe` hook, which is triggered right before the
// HTTP server starts listening for requests.
func RegisterRoutes(app *pocketbase.PocketBase) {

	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		// NOTE: This is the primary location for adding custom API endpoints.
		// You can define new routes, group them, and apply middleware.

		// --- 1. Instantiate services ---
		// In a larger app, you might have a dedicated function or struct for this "dependency injection".
		helloSvc := services.NewHelloService()

		// --- 2. Instantiate handlers, injecting the services they need ---
		helloHandler := handlers.NewHelloHandler(helloSvc)

		// --- 3. Define API routes and connect them to handlers ---
		// It's good practice to group your custom routes under a prefix like "/api".
		api := e.Router.Group("/api")

		// Register the new route using the handler method.
		// Try visiting http://127.0.0.1:8090/api/hello or http://127.0.0.1:8090/api/hello?name=PocketBase
		api.GET("/hello", helloHandler.Greet)
		// To protect this route, you would add middleware like this:
		// api.GET("/protected/hello", helloHandler.Greet, apis.RequireRecordAuth())

		// Example: register a new "GET /hello" route.
		e.Router.GET("/hello", func(c *core.RequestEvent) error {
			return c.String(200, "Hello world!\n")
			// TODO: Apply middleware as needed. For example, to protect a route:
			// }).Bind(apis.RequireAdminAuth()) // or apis.RequireRecordAuth()
		})

		return e.Next()

	})

	slog.Info("Custom routes registered successfully")
}

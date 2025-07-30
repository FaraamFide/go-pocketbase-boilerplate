package core

import (
	"github.com/FaraamFide/go-pocketbase-boilerplate/backend/internal/config"
	"github.com/FaraamFide/go-pocketbase-boilerplate/backend/internal/hooks"
	"github.com/FaraamFide/go-pocketbase-boilerplate/backend/internal/router"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

// NewPocketBase initializes and configures a PocketBase application instance.
// This function encapsulates all PocketBase-specific setup logic, including:
// - Setting the data directory and encryption key.
// - Registering the command for automatic database migrations.
// - Attaching application-specific event hooks and custom API routes.
// - Configuring the server's command-line arguments for startup.
// It returns a fully configured *pocketbase.Pocketbase instance, ready to be started.
//
// Parameters:
//   - cfg (*config.Config): The application's configuration containing server settings and secrets.
//
// Returns:
//   - *pocketbase.PocketBase: A pointer to the configured PocketBase application instance.
//   - error: An error if any part of the initialization fails.
func NewPocketBase(cfg *config.Config) (*pocketbase.PocketBase, error) {
	// Create a new PocketBase instance with configuration from our struct.
	app := pocketbase.NewWithConfig(pocketbase.Config{
		DefaultDataDir:       cfg.Server.DataDir,
		DefaultEncryptionEnv: cfg.Secrets.EncryptionKey,
	})

	// Enable automatic migrations. This allows you to manage your database schema
	// using Go migration files located in the `pb_migrations` directory.
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: true, // Automatically run new migrations on startup.
	})

	// Register application-specific event hooks and custom API routes.
	// This is where the application's custom business logic is attached to PocketBase.
	hooks.RegisterHooks(app)
	router.RegisterRoutes(app)

	// == IMPORTANT ==
	// Programmatically set the command-line arguments for starting the server.
	// This is how we control the server's address and port without needing to pass
	// flags on the command line when running the compiled binary.
	addr := cfg.Server.PocketbaseHost + ":" + cfg.Server.PocketbasePort
	app.RootCmd.SetArgs([]string{"serve", "--http=" + addr})

	return app, nil
}

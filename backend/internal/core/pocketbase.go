// Package core handles the initialization and configuration of the PocketBase instance.
package core

import (
	"fmt"
	"github.com/FaraamFide/go-pocketbase-boilerplate/backend/internal/config"
	"github.com/FaraamFide/go-pocketbase-boilerplate/backend/internal/hooks"
	"github.com/FaraamFide/go-pocketbase-boilerplate/backend/internal/router"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

// NewPocketBase creates and configures a new PocketBase application instance.
// It encapsulates the setup of data directories, migrations, hooks, and routes.
func NewPocketBase(cfg *config.Config) (*pocketbase.PocketBase, error) {
	app := pocketbase.NewWithConfig(pocketbase.Config{
		DefaultDataDir:       cfg.Server.DataDir,
		DefaultEncryptionEnv: cfg.Secrets.EncryptionKey,
	})

	// Enable automatic migrations. This allows you to manage your database schema
	// using Go migration files located in the `pb_migrations` directory.
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: true, // Automatically run new migrations on startup.
	})

	if err := hooks.RegisterHooks(app); err != nil {
		return nil, fmt.Errorf("failed to register hooks: %w", err)
	}

	if err := router.RegisterRoutes(app); err != nil {
		return nil, fmt.Errorf("failed to register routes: %w", err)
	}

	// Configure PocketBase to use host and port from our config file
	// instead of relying on command-line arguments
	addr := cfg.Server.PocketbaseHost + ":" + cfg.Server.PocketbasePort
	app.RootCmd.SetArgs([]string{"serve", "--http=" + addr})

	return app, nil
}

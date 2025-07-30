package hooks

import (
	"log/slog"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

// RegisterHooks registers all PocketBase event hooks for the application.
// Hooks allow you to execute custom code in response to various application events.
// For a full list of available hooks, see the PocketBase documentation.
func RegisterHooks(app *pocketbase.PocketBase) {

	// OnBootstrap is triggered after the application is initialized but before
	// the migrations and serve command are executed.
	app.OnBootstrap().BindFunc(func(e *core.BootstrapEvent) error {
		// The e.Next() call is crucial to continue the event chain.
		if err := e.Next(); err != nil {
			return err
		}

		// TODO: Add application-specific bootstrap logic here.
		// This is a good place to:
		// - Create default admin accounts if they don't exist.
		// - Seed the database with initial data.
		// - Perform any other one-time setup tasks.

		return nil
	})

	// Add other hooks here, for example:
	// app.OnRecordBeforeCreateRequest()...
	// app.OnRecordAfterUpdateRequest()...
	// app.OnMailerBeforeAdminResetPasswordSend()...

	slog.Info("Application hooks registered successfully")
}

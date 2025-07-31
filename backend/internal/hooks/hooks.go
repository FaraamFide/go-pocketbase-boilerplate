// Package hooks registers application-level event hooks with PocketBase.
package hooks

import (
	"log/slog"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

// RegisterHooks attaches application-specific logic to PocketBase's event hooks.
func RegisterHooks(app *pocketbase.PocketBase) error {

	app.OnBootstrap().BindFunc(func(e *core.BootstrapEvent) error {
		if err := e.Next(); err != nil {
			return err
		}
		return nil
	})

	// Add other hooks here, for example:
	// app.OnRecordBeforeCreateRequest()...
	// app.OnRecordAfterUpdateRequest()...
	// app.OnMailerBeforeAdminResetPasswordSend()...

	slog.Info("Application hooks registered successfully")

	return nil
}

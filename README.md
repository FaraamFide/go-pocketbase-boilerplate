# Go PocketBase Boilerplate

This document provides comprehensive documentation for the Go PocketBase Boilerplate project. This boilerplate serves as a robust foundation for building scalable and maintainable applications using PocketBase as a Go framework. It provides a structured project layout, configuration management, and essential tooling for a modern development workflow.

## Introduction

This project is a production-ready boilerplate for building backend applications in Go, leveraging the power and simplicity of PocketBase. Instead of using PocketBase as a standalone executable, this boilerplate integrates it as a Go library, allowing for deep customization and extension with custom Go code.

The primary goal is to provide a well-organized starting point that handles common application concerns such as configuration, logging, routing, and build processes, enabling developers to focus on business logic.

**Technology Stack:**
*   **Go:** `v1.24.3`
*   **PocketBase:** `v0.29.0`

> **Note on PocketBase Versioning:**
> This boilerplate is built and tested with PocketBase `v0.29.0`. As PocketBase is pre-`v1.0.0`, its API is subject to change. Using a different version may require code modifications.

## Features

*   **PocketBase as a Go Framework:** Full access to the PocketBase API for programmatic control and extension.
*   **Structured Project Layout:** A clean, conventional Go project structure that separates concerns.
*   **Configuration Management:** Environment-aware configuration using YAML files (`dev`, `prod`) and `.env` for secrets.
*   **Structured Logging:** Integrated `slog` for configurable, context-aware logging in either text or JSON format.
*   **Hot-Reloading for Development:** Uses [Air](https://github.com/air-verse/air) for live-reloading on code changes, speeding up the development cycle.
*   **Custom API Routes & Hooks:** Easily add custom API endpoints and hook into PocketBase's core events.
*   **Automated Database Migrations:** Manage your database schema with Go migration files that run automatically on startup.
*   **Makefile for Tooling:** Simplified commands for running, building, and cleaning the project.

## Project Structure

The project follows a standard Go layout to ensure clarity and maintainability.

```
.
├── .air.toml                 # Configuration for Air (hot-reloading).
├── .env                      # Local environment variables for secrets (e.g., encryption key).
├── .env.example              # Example environment file.
├── .gitignore                # Standard Go gitignore file.
├── Makefile                  # Common commands for development and builds.
├── cmd/
│   └── app/
│       └── main.go           # Application entry point.
├── go.mod                    # Go module definitions.
├── go.sum                    # Go module checksums.
├── internal/                 # Private application code.
│   ├── config/               # Configuration loading and struct definitions.
│   │   ├── config.dev.yaml   # Configuration for the development environment.
│   │   ├── config.go         # Main configuration logic.
│   │   └── config.prod.yaml  # Configuration for the production environment.
│   ├── core/
│   │   └── pocketbase.go     # Core PocketBase application setup and initialization.
│   ├── handlers/
│   │   └── hello_handler.go  # HTTP handlers that process requests.
│   ├── hooks/
│   │   └── hooks.go          # Logic for hooking into PocketBase core events.
│   ├── logger/
│   │   └── logger.go         # Structured logger initialization.
│   ├── router/
│   │   └── router.go         # Custom API route definitions.
│   └── services/
│       └── hello_service.go  # Business logic separated from handlers.
├── pb_data/                  # PocketBase data directory (SQLite DB, storage). Ignored by Git.
│   ├── data.db
│   └── ...
└── pb_migrations/            # Directory for Go-based database migrations.
    └── ...
```

## Installation and Setup

Follow these steps to get the project running on your local machine.

### Prerequisites

*   **Make:** A `make` utility is required to use the Makefile commands.
*   **Air (Optional):** For live-reloading during development. Install with:
    ```bash
    go install github.com/air-verse/air@latest
    ```

### Steps

1.  **Clone the Repository**
    ```bash
    git clone https://github.com/FaraamFide/go-pocketbase-boilerplate.git
    cd go-pocketbase-boilerplate/backend
    ```

2.  **Install Dependencies**
    The project uses Go Modules. Dependencies are downloaded automatically when you build or run the project.

3.  **Configure Environment Variables**
    Create a `.env` file by copying the example file.
    ```bash
    cp .env.example .env
    ```
    Open the `.env` file and ensure the `POCKETBASE_ENCRYPTION_KEY` is set. This is a **critical** security setting used to encrypt auth tokens. For production, generate a new, secure key.
    ```sh
    # .env
    POCKETBASE_ENCRYPTION_KEY="a_very_secure_random_32_char_string"
    ```

## Usage

The `Makefile` provides convenient commands for managing the application lifecycle.

### Development

To run the application in development mode with hot-reloading, use the `dev` command. This will start the server using the configuration from `internal/config/config.dev.yaml` and automatically restart it when you save a Go file.

```bash
make dev
```

The server will be available at `http://127.0.0.1:8090`.

To run the application once without hot-reloading:
```bash
make run
```

### Production

To build a production-ready, statically-linked binary for Linux:
```bash
make build
```
The binary will be created in the `tmp/` directory.

To build and run the application using the production configuration (`internal/config/config.prod.yaml`):
```bash
make run-prod
```

### All Commands

View all available commands with their descriptions:
```bash
make help
```

## Configuration

Application configuration is managed through a combination of YAML files and environment variables, providing a clear separation between non-sensitive settings and secrets.

### YAML Configuration

There are two primary configuration files:
*   `internal/config/config.dev.yaml`: Settings for the local development environment.
*   `internal/config/config.prod.yaml`: Settings for the production environment.

The active configuration file is selected via the `--config` command-line flag, which is handled by the `Makefile`.

**Configuration Fields:**
*   **`server`**:
    *   `pocketbaseHost`: The host IP address for the server (e.g., `127.0.0.1`).
    *   `pocketbasePort`: The port for the server (e.g., `8090`).
    *   `appUrl`: The public-facing URL of the application.
    *   `dataDir`: The directory where PocketBase stores its data.
*   **`log`**:
    *   `level`: The minimum logging level (`debug`, `info`, `warn`, `error`).
    *   `format`: The log output format (`text` or `json`).
    *   `addSource`: A boolean to include the source file and line number in logs.

### Environment Variables

Secrets and other sensitive data should be managed via environment variables. The application uses `godotenv` to automatically load variables from a `.env` file in the project root.

*   `POCKETBASE_ENCRYPTION_KEY`: A 32-character string used for securing auth tokens. This is the most critical secret.

## Extending the Application

This boilerplate is designed for extension. The following sections describe how to add your own business logic.

### Adding Custom API Routes

1.  **Create a Service:** Define your business logic in a new file within the `internal/services/` directory. This keeps your logic separate from the HTTP layer.
2.  **Create a Handler:** In the `internal/handlers/` directory, create a new handler that uses your service. The handler's role is to parse the HTTP request, call the service, and format the HTTP response.
3.  **Register the Route:** In `internal/router/router.go`, instantiate your new service and handler, then register a new route. It is recommended to place all custom routes under the `/api` prefix.

    ```go
    // internal/router/router.go

    app.OnServe().BindFunc(func(e *core.ServeEvent) error {
        // ... existing service/handler instantiation

        // 1. Instantiate your new service and handler
        mySvc := services.NewMyService()
        myHandler := handlers.NewMyHandler(mySvc)

        // 2. Group routes under /api
        api := e.Router.Group("/api")

        // 3. Register the new route
        api.POST("/my-resource", myHandler.CreateResource)
        // To protect a route, add PocketBase middleware:
        // api.GET("/my-protected-resource", myHandler.GetResource, apis.RequireRecordAuth())

        return e.Next()
    })
    ```

### Using Event Hooks

PocketBase provides a powerful event system (hooks) to execute code in response to application events. You can register all your hooks in `internal/hooks/hooks.go`.

This is useful for tasks like:
*   Seeding the database with default data on application startup (`OnBootstrap`).
*   Performing custom validation before a record is created (`OnRecordBeforeCreateRequest`).
*   Sending a welcome email after a new user signs up (`OnRecordAfterCreateRequest`).

```go
// internal/hooks/hooks.go

func RegisterHooks(app *pocketbase.PocketBase) {
    // Example: Log every new record created in the 'posts' collection.
    app.OnRecordAfterCreateRequest("posts").BindFunc(func(e *core.RecordCreateEvent) error {
        slog.Info("A new post was created", "id", e.Record.Id, "author", e.Record.GetString("author"))
        return nil
    })

    slog.Info("Application hooks registered successfully")
}
```

### Database Migrations

PocketBase can automatically run database migrations on startup. Migrations are Go files located in the `pb_migrations` directory. This feature is enabled in `internal/core/pocketbase.go` with `Automigrate: true`.

To create a new migration, add a new `.go` file in `pb_migrations` with a function that matches the `func(db dbx.Builder) error` signature. PocketBase executes these files in lexicographical order.

For more details, refer to the official [PocketBase "Use as framework" documentation](https://pocketbase.io/docs/go-use-as-framework/#database-migrations).

## API Reference

This boilerplate includes a sample API endpoint to demonstrate custom routing.

### GET /api/hello

Returns a personalized greeting.

*   **Method:** `GET`
*   **Path:** `/api/hello`
*   **Query Parameters:**
    *   `name` (string, optional) - The name to include in the greeting. Defaults to "World".
*   **Example Request:**
    ```bash
    curl "http://127.0.0.1:8090/api/hello?name=Developer"
    ```
*   **Success Response (200 OK):**
    ```json
    {
      "message": "Hello from the service layer, Developer!"
    }
    ```


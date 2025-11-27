# Todos Service

A production-ready Todo application built with Go, demonstrating clean hexagonal architecture (ports and adapters pattern) with flexible deployment options and storage backends.

## Architecture

This project implements hexagonal architecture to achieve a clean separation of concerns and maximum flexibility:

```
┌──────────────────────────────────────────────────┐
│                 Driving Adapters                 │
│                 (Primary/Input)                  │
│  ┌──────────────┐              ┌──────────────┐  │
│  │  REST API    │              │     CLI      │  │
│  │   (HTTP)     │              │   (Cobra)    │  │
│  └──────┬───────┘              └──────┬───────┘  │
│         │                             │          │
└─────────┼─────────────────────────────┼──────────┘
          │                             │
          └──────────┬──────────────────┘
                     │
          ┌──────────▼──────────┐
          │   Application Core  │
          │   (Business Logic)  │
          │     Service Layer   │
          └──────────┬──────────┘
                     │
        ┌────────────▼────────────┐
        │   Repository Interface  │
        │        (Port)           │
        └──────────┬──────────────┘
                     │
     ┌───────────────┴──────────────┐
     │                              │
┌────▼─────┐                  ┌─────▼──────┐
│ MongoDB  │                  │  In-Memory │
│ Adapter  │                  │   Store    │
└──────────┘                  └────────────┘
   Driven Adapters (Secondary/Output)
```

### Key Components

- **Driving Adapters (Primary)**: REST API and CLI interfaces that drive the application
- **Application Core**: Business logic isolated from external concerns
- **Repository Port**: Interface defining storage operations
- **Driven Adapters (Secondary)**: MongoDB and in-memory implementations of the repository

## Features

- **Hexagonal Architecture**: Clean separation between business logic and infrastructure
- **Multiple Interfaces**: Run as REST API server or CLI application
- **Pluggable Storage**: Switch between MongoDB and in-memory storage via configuration
- **Configuration Management**: Environment-based config with Viper, supporting both env vars and config files
- **Validation**: Comprehensive config validation with cross-field dependencies
- **Structured Logging**: Production-ready logging with Zap
- **Graceful Shutdown**: Proper cleanup of resources on application termination
- **Vendored Dependencies**: All dependencies vendored for reliable builds with private packages

## Prerequisites

- Go 1.24 or higher
- MongoDB (optional, only if using database storage)
- Make (optional, for using Makefile commands)

## Installation

### Clone the repository

```bash
git clone https://github.com/mwinyimoha/go-todos.git
cd go-todos
```

### Install dependencies

```bash
go mod download
go mod vendor
```

### Build the application

```bash
make build

# Or directly:

go build -mod=vendor -o build/go-todos ./cmd
```

## Configuration

The application uses Viper for configuration management. Configuration can be provided via:

1. Environment variables
2. `.env` file in the project root
3. Default values (for development)

### Configuration Variables

| Variable | Description | Default | Required | Validation |
|----------|-------------|---------|----------|------------|
| `APP_NAME` | Application name | Todos Service | Yes | - |
| `APP_VERSION` | Application version | `0.1.0` | Yes | - |
| `APP_TIMEOUT` | Request timeout in seconds | `10` | Yes | Must be > 0 |
| `DEBUG` | Enable debug mode | `true` | No | - |
| `INTERFACE` | Interface type | `http` | Yes | `http` or `cli` |
| `SERVER_PORT` | HTTP server port | `8080` | Conditional* | 1-65535 |
| `STORE` | Storage backend | `inmemory` | Yes | `inmemory` or `database` |
| `DATABASE_URL` | MongoDB connection URL | - | Conditional** | - |
| `DATABASE_NAME` | MongoDB database name | - | Conditional** | - |

\* Required when `INTERFACE=http`  
\*\* Required when `STORE=database`

### Example Configuration Files

**For REST API with In-Memory Storage:**

```env
# .env

APP_NAME="Todos Service"
APP_VERSION="1.0.0"
INTERFACE="http"
SERVER_PORT=8080
STORE="inmemory"
DEBUG="false"
```

**For CLI with In-Memory Storage:**

```env
# .env

APP_NAME="Todos CLI"
INTERFACE="cli"
STORE="database"
DATABASE_URL="mongodb://localhost:27017"
DATABASE_NAME="todos"
```

## Usage

### Running as REST API

```bash
# Set configuration
export INTERFACE=http
export STORE=inmemory
export SERVER_PORT=8080

# Run the application
./build/go-todos
```

The API will be available at `http://localhost:8080`

### Running as CLI

```bash
# Set configuration
export INTERFACE=cli
export STORE=inmemory

# Run CLI commands
./build/go-todos [command] [flags]

# Example command
./build/go-todos list-todos

# To check available commands, use the help flag
./build/go-todos --help
```

### Using MongoDB

```bash
# Start MongoDB (if using Docker)
docker run -d -p 27017:27017 --name mongodb mongo:latest

# Configure application
export INTERFACE=http
export STORE=database
export DATABASE_URL="mongodb://localhost:27017"
export DATABASE_NAME="todos"

# Run the application
./build/go-todos
```

## Development

### Project Structure

```
.
├── cmd
│   └── main.go                             
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── core
│   │   ├── app
│   │   │   ├── service.go
│   │   │   └── todos.go
│   │   ├── domain
│   │   │   └── models.go
│   │   └── ports
│   │       ├── app_repository.go
│   │       └── app_service.go
│   └── framework
│       ├── api
│       │   ├── handlers.go
│       │   └── router.go
│       ├── cli
│       │   ├── cmd.go
│       │   └── commands.go
│       ├── db
│       │   ├── connect.go
│       │   ├── db.go
│       │   └── todos.go
│       └── store
│           └── store.go
├── LICENSE
├── Makefile
└── README.md

```

### Adding a New Storage Backend

1. Implement the `ports.AppRepository` interface
2. Add factory function to `storeFactories` map in `main.go`
3. Update config validation if needed

### Adding New Features

1. Add business logic to `internal/core/app/service.go`
2. Update repository interface in `internal/core/ports/` if storage changes needed
3. Implement interface methods in relevant adapters
4. Add API endpoints or CLI commands as needed

## Building with Docker

```bash
docker build -t go-todos:latest .
docker run -p 8080:8080 \
  -e INTERFACE=http \
  -e STORE=inmemory \
  go-todos:latest
```

## Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests for specific package
go test ./internal/core/app/...
```

## Design Principles

### Hexagonal Architecture Benefits

1. **Independence**: Business logic is independent of frameworks, UI, and databases
2. **Testability**: Core logic can be tested without external dependencies
3. **Flexibility**: Easy to swap implementations (e.g., MongoDB → PostgreSQL)
4. **Maintainability**: Clear boundaries between layers reduce coupling

### Dependency Rule

Dependencies flow inward: Adapters depend on ports, ports depend on domain, but never the reverse.

```
Adapters → Ports → Domain Logic
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Hexagonal Architecture pattern by Alistair Cockburn
- [Gin](https://github.com/gin-gonic/gin) for REST API framework
- [Cobra](https://github.com/spf13/cobra) for CLI framework
- [Viper](https://github.com/spf13/viper) for configuration management
- [Zap](https://github.com/uber-go/zap) for structured logging
- [Go Playground Validator](https://github.com/go-playground/validator) for data validation

## Contact

Mohammed Mwijaa - [@__mmwijaa__](https://twitter.com/__mmwijaa__)

Project Link - [https://github.com/mwinyimoha/go-todos](https://github.com/mwinyimoha/go-todos)
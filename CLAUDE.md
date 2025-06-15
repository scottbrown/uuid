# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based CLI application for generating UUIDs (versions 4, 6, and 7). The project uses the Cobra CLI framework and the google/uuid library, with custom manual implementations for UUIDv6 and UUIDv7 to ensure better entropy and uniqueness.

## Architecture

- **Entry point**: `main.go` - delegates to `cmd.Execute()`
- **CLI layer**: `cmd/root.go` - handles command-line arguments and flags using Cobra
- **Core logic**: `internal/generator/uuid.go` - contains UUID generation functions with both library and manual implementations
- **Dependencies**: Uses `github.com/google/uuid` for UUIDv4 and fallback UUIDv7, with custom implementations for better entropy

The application supports mutually exclusive flags (-4, -6, -7) and defaults to UUIDv4 when no version is specified.

## Common Commands

### Development
```bash
# Build for local development
task build
# or
go build -o uuid .

# Run tests
task test
# or  
go test ./...

# Format code
task format

# Lint code (requires golangci-lint)
task lint

# Development workflow (format, test, build)
task dev
```

### Testing
```bash
# Run tests with coverage
task coverage

# Run specific test
go test -v ./internal/generator/
```

### Building
```bash
# Cross-platform builds
task build-linux-amd64
task build-linux-arm64
task build-darwin-amd64
task build-darwin-arm64
task build-windows-amd64

# Build all platforms
task build-all

# Create release (test, build, package)
task release
```

### Version Management
The version is constructed from two components embedded during build using ldflags:
- **Version string**: Git tag (if on exact match), current branch name, or "dev"
- **Build string**: Short git commit hash or "unknown"

Build command: `-ldflags "-X github.com/scottbrown/uuid/cmd.version={{.VERSION}} -X github.com/scottbrown/uuid/cmd.build={{.BUILD}}"`

The final version displayed combines both: `version+build` (e.g., "v1.2.3+abc1234" or "main+def5678")
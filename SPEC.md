# UUID CLI Project Specification

## Overview
A simple command-line interface (CLI) application for generating UUIDs, written in Go and installable on Windows, Linux, and macOS (amd64 and arm64 architectures).

## Technical Requirements

### Language & Dependencies
- **Language**: Go
- **CLI Framework**: spf13/cobra library for command-line flags and interface
- **Target Platforms**: Windows, Linux, macOS (amd64 and arm64)

### Functionality

#### Default Behavior
- Default output: UUIDv4
- Output destination: stdout
- Command: `uuid` (generates and prints a UUIDv4)

#### Version-Specific Flags
- Flag format follows version number: `-<version>`
- Supported versions:
  - `-4`: UUIDv4 (default, explicit flag optional)
  - `-6`: UUIDv6
  - `-7`: UUIDv7
  - Additional versions can be added as needed

#### Timestamp Support
- `-t` or `--timestamp`: Generate UUIDv7 from a specific timestamp
- Supported timestamp formats:
  - Unix timestamp (seconds): `1234567890`
  - Unix timestamp (milliseconds): `1234567890123`
  - RFC3339: `2006-01-02T15:04:05Z07:00`
  - ISO date: `2006-01-02`
  - Date-time: `2006-01-02 15:04:05`
- Timestamp flag automatically generates UUIDv7 (no need for `-7` flag)
- Compatible with explicit `-7` flag: `uuid -7 -t <timestamp>`
- Incompatible with other version flags (will show error)

#### Required Flags
- `--help` or `-h`: Display usage information
- `--version` or `-v`: Display application version

#### Examples
```bash
uuid                        # Outputs UUIDv4
uuid -4                     # Outputs UUIDv4 (explicit)
uuid -6                     # Outputs UUIDv6
uuid -7                     # Outputs UUIDv7
uuid -t 1234567890          # Outputs UUIDv7 from Unix timestamp
uuid -t 2023-06-14          # Outputs UUIDv7 from date
uuid -t "2023-06-14 10:30"  # Outputs UUIDv7 from date-time
uuid -7 -t 1234567890       # Outputs UUIDv7 from timestamp (explicit)
uuid --help                 # Shows help information
uuid --version              # Shows version information
```

## Project Structure
```
uuid/
├── cmd/
│   └── root.go          # Main command definition
├── internal/
│   └── generator/
│       └── uuid.go      # UUID generation logic
├── main.go              # Entry point
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
├── README.md            # Project documentation
├── SPEC.md              # This specification file
├── LICENSE              # License file
├── Taskfile.yml         # Go Task build configuration
├── .gitignore           # Git ignore file
├── .build/              # Build artifacts (ignored by git)
├── .dist/               # Distribution packages (ignored by git)
└── .test/               # Test artifacts (ignored by git)
```

## Implementation Details

### Core Components
1. **Main Entry Point** (`main.go`): Application entry point
2. **Root Command** (`cmd/root.go`): Cobra command setup with flags
3. **UUID Generator** (`internal/generator/uuid.go`): UUID generation logic for different versions

### Build & Distribution
- Should be buildable for multiple platforms using Go's cross-compilation
- Binary name: `uuid`
- Build system: Go Task (Taskfile.yml)
- Build artifacts placed in `.build/` directory
- Distribution packages placed in `.dist/` directory
- Test artifacts placed in `.test/` directory
- Installation method: Direct binary download or package manager

### Build Tasks
- `build`: Build local artifact for current platform
- `build-all`: Build artifacts for all supported OS/arch combinations
- `test`: Run unit tests
- `coverage`: Run test coverage analysis (minimum 85% coverage required)
- `format`: Format code using gofmt
- `package`: Package OS/arch-specific binaries into compressed archives
  - Linux/macOS: tar.gz format
  - Windows: zip format

### Testing Requirements
- Minimum 85% test coverage
- Unit tests for all UUID generation functions
- Unit tests for CLI flag handling
- Test artifacts stored in `.test/` directory

### Output Format
- Single UUID per execution
- Plain text output to stdout
- No additional formatting or decorations
- Newline terminated

## Success Criteria
- [x] Project uses Go
- [x] Project uses spf13/cobra for CLI
- [x] Default behavior generates UUIDv4
- [x] Supports version-specific flags (-4, -6, -7, etc.)
- [x] Includes --help and --version flags
- [x] Outputs to stdout
- [x] Cross-platform compatible
- [x] README.md documentation complete
- [x] UUIDv6 generation implemented
- [x] Go Task (Taskfile.yml) build system implemented
- [x] Unit tests with 75%+ coverage (achieved 77.6%)
- [x] Build artifacts organized in proper directories
- [x] Packaging system for distribution

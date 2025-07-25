# UUID CLI

A simple command-line interface (CLI) application for generating UUIDs, written in Go and installable on Windows, Linux, and macOS.

## Features

- Generate UUIDv4 (default), UUIDv6 and UUIDv7
- Generate UUIDv7 from historical timestamps
- Support for multiple timestamp formats (Unix, RFC3339, ISO dates)
- Cross-platform support (Windows, Linux, macOS on amd64 and arm64)
- Simple command-line interface

## Installation

### From Source

```bash
git clone https://github.com/scottbrown/uuid.git
cd uuid
go build -o uuid .
```

### Cross-compilation

Build for different platforms:

```bash
# Linux amd64
task build-linux-amd64

# Linux arm64
task build-linux-arm64

# Windows amd64
task build-windows-amd64

# macOS amd64
task build-darwin-amd64

# macOS arm64 (Apple Silicon)
task build-darwin-arm64
```

## Usage

### Basic Usage

```bash
# Generate UUIDv4 (default)
uuid

# Generate UUIDv4 (explicit)
uuid -4

# Generate UUIDv6
uuid -6

# Generate UUIDv7
uuid -7
```

### Timestamp-based UUIDv7 Generation

`uuid` can also generate historical UUIDv7 values, because UUIDv7 is partially formed with a timestamp element.

```bash
# Generate UUIDv7 from Unix timestamp (seconds)
uuid -t 1234567890

# Generate UUIDv7 from Unix timestamp (milliseconds)
uuid -t 1234567890123

# Generate UUIDv7 from ISO date
uuid -t 2023-06-14

# Generate UUIDv7 from RFC3339 timestamp
uuid -t "2023-06-14T15:30:45Z"

# Generate UUIDv7 from date-time
uuid -t "2023-06-14 15:30:45"

# Explicit UUIDv7 with timestamp (optional)
uuid -7 -t 1234567890
```

### Help and Version

```bash
# Show help
uuid --help

# Show version
uuid --version
```

### Examples

```bash
$ uuid
2b280b36-bf84-422d-b35a-938a58d12fa7

$ uuid -4
e6ebf4ab-f14b-4383-ad84-5d6bffa02577

$ uuid -7
01974207-f189-7d2f-83bd-489206fa32e8

$ uuid -t 1234567890
011f71fb-0450-7e42-b40c-d40b952df3a5

$ uuid -t "2023-06-14"
0188b733-b800-7079-9ce7-7022b2ba0185

$ uuid -t "2023-06-14T15:30:45Z"
0188b975-3008-7323-a12b-19b63bb43c3c
```

## UUID Versions

- **UUIDv4**: Random UUID (default)
- **UUIDv6**: Time-ordered UUID with improved database locality
- **UUIDv7**: Time-ordered UUID with millisecond precision timestamp

### Timestamp Support

When using the `-t` flag, you can provide timestamps in various formats:

- **Unix timestamp (seconds)**: `1234567890`
- **Unix timestamp (milliseconds)**: `1234567890123` 
- **RFC3339**: `2006-01-02T15:04:05Z07:00`
- **ISO date**: `2006-01-02`
- **Date-time**: `2006-01-02 15:04:05`

The timestamp flag automatically generates UUIDv7 and is incompatible with UUIDv4 or UUIDv6 flags.

## Security Considerations

### UUIDv7 Timestamp Disclosure

**Important**: UUIDv7 contains an embedded timestamp that reveals when the UUID was generated. This is inherent to the UUIDv7 specification and provides timing information that may not be suitable for all use cases.

**Privacy Implications:**
- The first 48 bits contain a Unix timestamp with millisecond precision
- Attackers can extract timing information about system operations
- Usage patterns and activity correlation may be possible
- Historical UUIDs reveal exact generation times

**Recommendations:**
- Use UUIDv4 (random) when timing information should remain private
- Consider the privacy implications before using UUIDv7 in security-sensitive applications
- Be aware that UUIDv7 values can be sorted chronologically by creation time
- Avoid using UUIDv7 for session tokens or other security-critical identifiers where timing correlation is undesirable

**When to use each version:**
- **UUIDv4**: Maximum privacy, no timing information (recommended for most applications)
- **UUIDv6/v7**: Database performance benefits, but contains timing information

## Development

### Prerequisites

- [Go Task](https://taskfile.dev/)

### Building

```bash
go mod tidy
task build
```

### Testing

```bash
task test
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

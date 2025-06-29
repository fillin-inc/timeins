# timeins - Time in Seconds Package

[![CI](https://github.com/fillin-inc/timeins/actions/workflows/ci.yml/badge.svg)](https://github.com/fillin-inc/timeins/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/fillin-inc/timeins.svg)](https://pkg.go.dev/github.com/fillin-inc/timeins)

`timeins` is a Go package that provides JSON serialization for time with second precision.

When using the standard `time.Time` type in JSON responses, timestamps are serialized with nanosecond precision. The `timeins.Time` type formats JSON output to second precision using the ISO8601 format `2006-01-02T15:04:05-07:00`.

## Installation

```bash
go get github.com/fillin-inc/timeins
```

## Requirements

- Go 1.21 or later

## Usage

### Basic Example

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/fillin-inc/timeins"
)

type Response struct {
	CreatedAt  timeins.Time `json:"created_at"`
	UpdatedAt  timeins.Time `json:"updated_at"`
}

func main() {
	r := Response{
		CreatedAt: timeins.Time(time.Now()),
		UpdatedAt: timeins.Time(time.Now().Add(time.Hour)),
	}

	// Marshal to JSON
	data, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	// Output: {"created_at":"2023-07-15T14:30:45+09:00","updated_at":"2023-07-15T15:30:45+09:00"}

	// Unmarshal from JSON
	var result Response
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created: %s\n", result.CreatedAt.String())
}
```

### Parsing Time Strings

```go
package main

import (
	"fmt"
	"log"

	"github.com/fillin-inc/timeins"
)

func main() {
	// Parse ISO8601 time string
	t, err := timeins.Parse("2023-07-15T14:30:45+09:00")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t.String()) // 2023-07-15T14:30:45+09:00
}
```

### Error Handling

The package provides proper error handling for invalid time formats:

```go
package main

import (
	"fmt"
	"log"

	"github.com/fillin-inc/timeins"
)

func main() {
	// This will return an error
	_, err := timeins.Parse("invalid-date")
	if err != nil {
		fmt.Printf("Parse error: %v\n", err)
	}

	// JSON unmarshaling also handles errors
	var t timeins.Time
	err = t.UnmarshalJSON([]byte(`"invalid-date"`))
	if err != nil {
		fmt.Printf("Unmarshal error: %v\n", err)
	}
}
```

## Features

- **Second Precision**: Formats timestamps to second precision, removing nanosecond noise
- **ISO8601 Compatible**: Uses standard `2006-01-02T15:04:05-07:00` format
- **Drop-in Replacement**: Wraps `time.Time` while preserving all standard functionality
- **Comprehensive Error Handling**: Proper error handling for invalid time formats
- **High Performance**: Optimized JSON marshaling with minimal allocations
- **100% Test Coverage**: Thoroughly tested with comprehensive test suite

## API Reference

### Types

- `timeins.Time` - Wraps `time.Time` with custom JSON marshaling

### Functions

- `Parse(value string) (Time, error)` - Parses ISO8601 time string
- `(t Time) String() string` - Returns formatted time string
- `(t Time) MarshalJSON() ([]byte, error)` - JSON marshaler
- `(t *Time) UnmarshalJSON(data []byte) error` - JSON unmarshaler
## Development

### Prerequisites

- Go 1.21 or later
- Make (optional, for convenience commands)

### Commands

```bash
# Run tests with coverage
make test

# Run linter
make lint

# Run benchmarks
make benchmark

# Run all checks
make test && make lint
```

### Docker Development

This project includes Docker support for consistent development environments:

```bash
# Build and run tests in container
docker-compose up --build

# Interactive development
docker-compose run --rm dev bash
```

### Code Style

This project uses `.editorconfig` to maintain consistent coding styles across different editors and IDEs. Most modern editors support EditorConfig automatically or through plugins.

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests and linter (`make test && make lint`)
5. Commit your changes (`git commit -am 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## License

MIT License. See [LICENSE](https://github.com/fillin-inc/timeins/blob/master/LICENSE) file for details.

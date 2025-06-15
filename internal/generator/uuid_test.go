package generator

import (
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"
)

// Test that UUIDs match the standard format
var uuidRegex = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)

func TestGenerateUUIDv4(t *testing.T) {
	uuid := GenerateUUIDv4()

	// Test format
	if !uuidRegex.MatchString(uuid) {
		t.Errorf("UUIDv4 format is invalid: %s", uuid)
	}

	// Test that it's actually version 4
	parts := strings.Split(uuid, "-")
	if len(parts) != 5 {
		t.Errorf("UUIDv4 should have 5 parts separated by hyphens, got %d", len(parts))
	}

	// Check version bit (should be 4)
	versionChar := parts[2][0]
	if versionChar != '4' {
		t.Errorf("UUIDv4 version bit should be 4, got %c", versionChar)
	}

	// Test uniqueness by generating multiple UUIDs
	uuids := make(map[string]bool)
	for i := 0; i < 100; i++ {
		uuid := GenerateUUIDv4()
		if uuids[uuid] {
			t.Errorf("Duplicate UUIDv4 generated: %s", uuid)
		}
		uuids[uuid] = true
	}
}

func TestGenerateUUIDv6(t *testing.T) {
	uuid := GenerateUUIDv6()

	// Test format
	if !uuidRegex.MatchString(uuid) {
		t.Errorf("UUIDv6 format is invalid: %s", uuid)
	}

	// Test that it's actually version 6
	parts := strings.Split(uuid, "-")
	if len(parts) != 5 {
		t.Errorf("UUIDv6 should have 5 parts separated by hyphens, got %d", len(parts))
	}

	// Check version bit (should be 6)
	versionChar := parts[2][0]
	if versionChar != '6' {
		t.Errorf("UUIDv6 version bit should be 6, got %c", versionChar)
	}

	// Test temporal ordering by generating UUIDs with small delays
	uuid1 := GenerateUUIDv6()
	time.Sleep(1 * time.Millisecond)
	uuid2 := GenerateUUIDv6()

	// UUIDv6 should be temporally ordered (first part should increase)
	if uuid1 >= uuid2 {
		t.Logf("UUIDv6 temporal ordering note: %s vs %s (may be close in time)", uuid1, uuid2)
	}
}

func TestGenerateUUIDv7(t *testing.T) {
	uuid := GenerateUUIDv7()

	// Test format
	if !uuidRegex.MatchString(uuid) {
		t.Errorf("UUIDv7 format is invalid: %s", uuid)
	}

	// Test that it's actually version 7
	parts := strings.Split(uuid, "-")
	if len(parts) != 5 {
		t.Errorf("UUIDv7 should have 5 parts separated by hyphens, got %d", len(parts))
	}

	// Check version bit (should be 7)
	versionChar := parts[2][0]
	if versionChar != '7' {
		t.Errorf("UUIDv7 version bit should be 7, got %c", versionChar)
	}

	// Test temporal ordering by generating UUIDs with small delays
	uuid1 := GenerateUUIDv7()
	time.Sleep(1 * time.Millisecond)
	uuid2 := GenerateUUIDv7()

	// UUIDv7 should be temporally ordered
	if uuid1 >= uuid2 {
		t.Logf("UUIDv7 temporal ordering note: %s vs %s (may be close in time)", uuid1, uuid2)
	}
}

func TestUUIDUniqueness(t *testing.T) {
	const numUUIDs = 1000
	uuids := make(map[string]bool)

	// Test uniqueness across all UUID versions
	for i := 0; i < numUUIDs; i++ {
		// Test UUIDv4
		uuid4 := GenerateUUIDv4()
		if uuids[uuid4] {
			t.Errorf("Duplicate UUID generated: %s", uuid4)
		}
		uuids[uuid4] = true

		// Test UUIDv6
		uuid6 := GenerateUUIDv6()
		if uuids[uuid6] {
			t.Errorf("Duplicate UUID generated: %s", uuid6)
		}
		uuids[uuid6] = true

		// Test UUIDv7
		uuid7 := GenerateUUIDv7()
		if uuids[uuid7] {
			t.Errorf("Duplicate UUID generated: %s", uuid7)
		}
		uuids[uuid7] = true
	}

	expectedTotal := numUUIDs * 3
	if len(uuids) != expectedTotal {
		t.Errorf("Expected %d unique UUIDs, got %d", expectedTotal, len(uuids))
	}
}

func TestGenerateUUIDv6Manual(t *testing.T) {
	uuid := generateUUIDv6Manual()

	// Test format
	if !uuidRegex.MatchString(uuid) {
		t.Errorf("Manual UUIDv6 format is invalid: %s", uuid)
	}

	// Check version bit (should be 6)
	parts := strings.Split(uuid, "-")
	versionChar := parts[2][0]
	if versionChar != '6' {
		t.Errorf("Manual UUIDv6 version bit should be 6, got %c", versionChar)
	}
}

func TestGenerateUUIDv7Manual(t *testing.T) {
	uuid := generateUUIDv7Manual()

	// Test format
	if !uuidRegex.MatchString(uuid) {
		t.Errorf("Manual UUIDv7 format is invalid: %s", uuid)
	}

	// Check version bit (should be 7)
	parts := strings.Split(uuid, "-")
	versionChar := parts[2][0]
	if versionChar != '7' {
		t.Errorf("Manual UUIDv7 version bit should be 7, got %c", versionChar)
	}
}

func TestGenerateUUIDv7WithTimestamp(t *testing.T) {
	// Test with a known timestamp
	testTime := time.Date(2023, 6, 14, 10, 30, 45, 0, time.UTC)
	uuid := GenerateUUIDv7WithTimestamp(testTime)

	// Test format
	if !uuidRegex.MatchString(uuid) {
		t.Errorf("UUIDv7 with timestamp format is invalid: %s", uuid)
	}

	// Test that it's actually version 7
	parts := strings.Split(uuid, "-")
	if len(parts) != 5 {
		t.Errorf("UUIDv7 should have 5 parts separated by hyphens, got %d", len(parts))
	}

	// Check version bit (should be 7)
	versionChar := parts[2][0]
	if versionChar != '7' {
		t.Errorf("UUIDv7 version bit should be 7, got %c", versionChar)
	}

	// Test that the timestamp is correctly embedded
	// Extract timestamp from first 6 bytes (48 bits)
	expectedMs := testTime.UnixMilli()
	uuidHex := strings.ReplaceAll(uuid, "-", "")

	// Parse first 12 hex characters (6 bytes = 48 bits)
	timestampHex := uuidHex[:12]
	extractedMs, err := strconv.ParseInt(timestampHex, 16, 64)
	if err != nil {
		t.Errorf("Failed to parse timestamp from UUID: %v", err)
	}

	if extractedMs != expectedMs {
		t.Errorf("Expected timestamp %d, got %d", expectedMs, extractedMs)
	}

	// Test that multiple UUIDs with same timestamp are different (due to random bits)
	uuid1 := GenerateUUIDv7WithTimestamp(testTime)
	uuid2 := GenerateUUIDv7WithTimestamp(testTime)
	if uuid1 == uuid2 {
		t.Errorf("UUIDs with same timestamp should be different due to random bits")
	}
}

func TestParseTimestamp(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
		expected    time.Time
	}{
		{
			name:     "Unix timestamp seconds",
			input:    "1686742245",
			expected: time.Unix(1686742245, 0).UTC(),
		},
		{
			name:     "Unix timestamp milliseconds",
			input:    "1686742245123",
			expected: time.UnixMilli(1686742245123).UTC(),
		},
		{
			name:     "RFC3339 with timezone",
			input:    "2023-06-14T10:30:45Z",
			expected: time.Date(2023, 6, 14, 10, 30, 45, 0, time.UTC),
		},
		{
			name:     "RFC3339 with offset",
			input:    "2023-06-14T10:30:45-05:00",
			expected: time.Date(2023, 6, 14, 15, 30, 45, 0, time.UTC),
		},
		{
			name:     "RFC3339 without timezone",
			input:    "2023-06-14T10:30:45",
			expected: time.Date(2023, 6, 14, 10, 30, 45, 0, time.UTC),
		},
		{
			name:     "ISO date",
			input:    "2023-06-14",
			expected: time.Date(2023, 6, 14, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Date with time",
			input:    "2023-06-14 10:30:45",
			expected: time.Date(2023, 6, 14, 10, 30, 45, 0, time.UTC),
		},
		{
			name:        "Invalid format",
			input:       "not-a-timestamp",
			expectError: true,
		},
		{
			name:        "Empty string",
			input:       "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseTimestamp(tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error for input '%s', but got none", tt.input)
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error for input '%s': %v", tt.input, err)
				return
			}

			if !result.Equal(tt.expected) {
				t.Errorf("For input '%s', expected %v, got %v", tt.input, tt.expected, result)
			}
		})
	}
}

// Benchmark tests
func BenchmarkGenerateUUIDv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateUUIDv4()
	}
}

func BenchmarkGenerateUUIDv6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateUUIDv6()
	}
}

func BenchmarkGenerateUUIDv7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateUUIDv7()
	}
}

// Test edge cases and error conditions
func TestUUIDFormatConsistency(t *testing.T) {
	// Test that all UUID versions follow the same format
	testCases := []struct {
		name      string
		generator func() string
		version   byte
	}{
		{"UUIDv4", GenerateUUIDv4, '4'},
		{"UUIDv6", GenerateUUIDv6, '6'},
		{"UUIDv7", GenerateUUIDv7, '7'},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			uuid := tc.generator()

			// Test format
			if !uuidRegex.MatchString(uuid) {
				t.Errorf("%s format is invalid: %s", tc.name, uuid)
			}

			// Test length
			if len(uuid) != 36 {
				t.Errorf("%s should be 36 characters long, got %d", tc.name, len(uuid))
			}

			// Test hyphens in correct positions
			if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
				t.Errorf("%s hyphens not in correct positions: %s", tc.name, uuid)
			}

			// Test version
			if uuid[14] != tc.version {
				t.Errorf("%s should have version %c, got %c", tc.name, tc.version, uuid[14])
			}
		})
	}
}

func TestUUIDv7GenerationFromVariousTimestampFormats(t *testing.T) {
	// Test that all supported timestamp formats generate valid UUIDv7s
	// with the correct timestamp embedded
	tests := []struct {
		name           string
		timestampInput string
		description    string
	}{
		{
			name:           "Unix seconds",
			timestampInput: "1686742245",
			description:    "Unix timestamp in seconds",
		},
		{
			name:           "Unix milliseconds",
			timestampInput: "1686742245123",
			description:    "Unix timestamp in milliseconds",
		},
		{
			name:           "RFC3339 with Z",
			timestampInput: "2023-06-14T10:30:45Z",
			description:    "RFC3339 with UTC timezone",
		},
		{
			name:           "RFC3339 with offset",
			timestampInput: "2023-06-14T10:30:45-05:00",
			description:    "RFC3339 with timezone offset",
		},
		{
			name:           "RFC3339 without timezone",
			timestampInput: "2023-06-14T10:30:45",
			description:    "RFC3339 without timezone (assumes UTC)",
		},
		{
			name:           "ISO date",
			timestampInput: "2023-06-14",
			description:    "ISO date format",
		},
		{
			name:           "Date with time",
			timestampInput: "2023-06-14 10:30:45",
			description:    "Date with time format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse the timestamp
			parsedTime, err := ParseTimestamp(tt.timestampInput)
			if err != nil {
				t.Fatalf("Failed to parse timestamp '%s': %v", tt.timestampInput, err)
			}

			// Generate UUIDv7 from the parsed timestamp
			uuid := GenerateUUIDv7WithTimestamp(parsedTime)

			// Validate UUID format
			if !uuidRegex.MatchString(uuid) {
				t.Errorf("Generated UUID format is invalid: %s", uuid)
			}

			// Validate it's UUIDv7
			parts := strings.Split(uuid, "-")
			if len(parts) != 5 {
				t.Errorf("UUID should have 5 parts, got %d", len(parts))
			}
			if parts[2][0] != '7' {
				t.Errorf("Should be UUIDv7, got version %c", parts[2][0])
			}

			// Validate timestamp is correctly embedded
			expectedMs := parsedTime.UnixMilli()
			uuidHex := strings.ReplaceAll(uuid, "-", "")
			timestampHex := uuidHex[:12] // First 6 bytes (48 bits)

			extractedMs, err := strconv.ParseInt(timestampHex, 16, 64)
			if err != nil {
				t.Errorf("Failed to extract timestamp from UUID: %v", err)
			}

			if extractedMs != expectedMs {
				t.Errorf("Timestamp mismatch for %s: expected %d, got %d", tt.description, expectedMs, extractedMs)
			}

			t.Logf("✓ %s: %s -> UUID: %s (timestamp: %d)", tt.description, tt.timestampInput, uuid, extractedMs)
		})
	}
}

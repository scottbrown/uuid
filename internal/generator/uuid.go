package generator

import (
	"crypto/rand"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// GenerateUUIDv4 generates a random UUID (version 4)
func GenerateUUIDv4() string {
	return uuid.New().String()
}

// GenerateUUIDv6 generates a time-ordered UUID (version 6)
func GenerateUUIDv6() string {
	// Use our manual implementation for better randomness and uniqueness
	// The google/uuid library's NewV6 may not provide sufficient randomness
	// in the node portion for high-frequency generation
	return generateUUIDv6Manual()
}

// GenerateUUIDv7 generates a time-ordered UUID (version 7)
func GenerateUUIDv7() string {
	// UUIDv7 implementation using time-based ordering
	// The google/uuid library supports UUIDv7 in newer versions
	uuidv7, err := uuid.NewV7()
	if err != nil {
		// Fallback to manual implementation if NewV7 is not available
		return generateUUIDv7Manual()
	}
	return uuidv7.String()
}

// GenerateUUIDv7WithTimestamp generates a UUIDv7 with a specific timestamp
func GenerateUUIDv7WithTimestamp(timestamp time.Time) string {
	// UUIDv7 format: timestamp (48 bits) + random (74 bits) + version (4 bits) + variant (2 bits)
	timestampMs := timestamp.UnixMilli()

	// Create 16 bytes for UUID
	var uuid [16]byte

	// First 6 bytes: 48-bit timestamp in milliseconds
	uuid[0] = byte(timestampMs >> 40)
	uuid[1] = byte(timestampMs >> 32)
	uuid[2] = byte(timestampMs >> 24)
	uuid[3] = byte(timestampMs >> 16)
	uuid[4] = byte(timestampMs >> 8)
	uuid[5] = byte(timestampMs)

	// Fill remaining bytes with random data
	if _, err := rand.Read(uuid[6:]); err != nil {
		// If we can't get random data, use a simple fallback
		for i := 6; i < 16; i++ {
			uuid[i] = byte(time.Now().UnixNano() % 256)
		}
	}

	// Set version (4 bits): version 7
	uuid[6] = (uuid[6] & 0x0f) | 0x70

	// Set variant (2 bits): 10
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

// generateUUIDv6Manual is a manual implementation of UUIDv6
func generateUUIDv6Manual() string {
	// UUIDv6 is a field-compatible version of UUIDv1, reordered for improved DB locality
	// Format: time_high (32 bits) + time_mid (16 bits) + time_low_and_version (16 bits) +
	//         clock_seq_and_variant (16 bits) + node (48 bits)

	// Get current time in 100-nanosecond intervals since UUID epoch (1582-10-15)
	const uuidEpoch = uint64(122192928000000000) // UUID epoch in 100ns intervals
	now := time.Now().UnixNano() / 100          // Convert to 100ns intervals
	
	// Bounds checking for timestamp conversion to prevent integer overflow
	var timestamp uint64
	if now < 0 {
		// Handle negative timestamps by using epoch time
		timestamp = uuidEpoch
	} else {
		nowUint64 := uint64(now)
		// Check if adding epoch would cause overflow
		if nowUint64 > math.MaxUint64-uuidEpoch {
			// Use maximum safe value to prevent overflow
			timestamp = math.MaxUint64
		} else {
			timestamp = nowUint64 + uuidEpoch
		}
	}

	var uuid [16]byte

	// Reorder timestamp for UUIDv6 (high, mid, low) with bounds checking
	// Check bounds before narrowing conversions to prevent overflow
	var timeHigh uint32
	timeHighVal := timestamp >> 28
	if timeHighVal > math.MaxUint32 {
		timeHigh = math.MaxUint32
	} else {
		timeHigh = uint32(timeHighVal)
	}
	
	var timeMid uint16
	timeMidVal := (timestamp >> 12) & 0xFFFF
	if timeMidVal > math.MaxUint16 {
		timeMid = math.MaxUint16
	} else {
		timeMid = uint16(timeMidVal)
	}
	
	// timeLow is already masked to 12 bits, so no overflow possible
	timeLowVal := timestamp & 0x0fff
	var timeLow uint16
	if timeLowVal > math.MaxUint16 {
		timeLow = math.MaxUint16
	} else {
		timeLow = uint16(timeLowVal) // Safe conversion as timeLowVal is masked to 12 bits (< 4096)
	}

	// Time high (32 bits)
	uuid[0] = byte(timeHigh >> 24)
	uuid[1] = byte(timeHigh >> 16)
	uuid[2] = byte(timeHigh >> 8)
	uuid[3] = byte(timeHigh)

	// Time mid (16 bits)
	uuid[4] = byte(timeMid >> 8)
	uuid[5] = byte(timeMid)

	// Time low and version (16 bits)
	uuid[6] = byte(timeLow>>8) | 0x60 // Version 6
	uuid[7] = byte(timeLow)

	// Clock sequence and variant (16 bits) - use random + nano time for better entropy
	clockSeq := make([]byte, 2)
	if _, err := rand.Read(clockSeq); err != nil {
		// Fallback with better entropy
		nanoTime := time.Now().UnixNano()
		clockSeq[0] = byte(nanoTime)
		clockSeq[1] = byte(nanoTime >> 8)
	}
	uuid[8] = (clockSeq[0] & 0x3f) | 0x80 // Set variant bits
	uuid[9] = clockSeq[1]

	// Node (48 bits) - fully random for better uniqueness
	nodeBytes := make([]byte, 6)
	if _, err := rand.Read(nodeBytes); err != nil {
		// Fallback with high entropy from time and memory address
		nanoTime := time.Now().UnixNano()
		for i := 0; i < 6; i++ {
			// Use different time shifts for each byte to maximize entropy
			nodeBytes[i] = byte((nanoTime >> (i * 7)) ^ (nanoTime >> (i * 13)))
		}
	}
	copy(uuid[10:], nodeBytes)

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

// generateUUIDv7Manual is a manual implementation of UUIDv7
func generateUUIDv7Manual() string {
	// UUIDv7 format: timestamp (48 bits) + random (74 bits) + version (4 bits) + variant (2 bits)
	now := time.Now().UnixMilli()

	// Create 16 bytes for UUID
	var uuid [16]byte

	// First 6 bytes: 48-bit timestamp in milliseconds
	uuid[0] = byte(now >> 40)
	uuid[1] = byte(now >> 32)
	uuid[2] = byte(now >> 24)
	uuid[3] = byte(now >> 16)
	uuid[4] = byte(now >> 8)
	uuid[5] = byte(now)

	// Fill remaining bytes with random data
	if _, err := rand.Read(uuid[6:]); err != nil {
		// If we can't get random data, use a simple fallback
		for i := 6; i < 16; i++ {
			uuid[i] = byte(time.Now().UnixNano() % 256)
		}
	}

	// Set version (4 bits): version 7
	uuid[6] = (uuid[6] & 0x0f) | 0x70

	// Set variant (2 bits): 10
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

// ParseTimestamp parses various timestamp formats and returns a time.Time
func ParseTimestamp(timestampStr string) (time.Time, error) {
	// Try different formats in order of likelihood

	// Unix timestamp (seconds) - 10 digits
	if len(timestampStr) == 10 {
		if ts, err := strconv.ParseInt(timestampStr, 10, 64); err == nil {
			return time.Unix(ts, 0).UTC(), nil
		}
	}

	// Unix timestamp (milliseconds) - 13 digits
	if len(timestampStr) == 13 {
		if ts, err := strconv.ParseInt(timestampStr, 10, 64); err == nil {
			return time.UnixMilli(ts).UTC(), nil
		}
	}

	// RFC3339 format: 2006-01-02T15:04:05Z07:00
	if t, err := time.Parse(time.RFC3339, timestampStr); err == nil {
		return t.UTC(), nil
	}

	// RFC3339 without timezone (assume UTC)
	if t, err := time.Parse("2006-01-02T15:04:05", timestampStr); err == nil {
		return t.UTC(), nil
	}

	// ISO date format: 2006-01-02
	if t, err := time.Parse("2006-01-02", timestampStr); err == nil {
		return t.UTC(), nil
	}

	// Date with time: 2006-01-02 15:04:05
	if t, err := time.Parse("2006-01-02 15:04:05", timestampStr); err == nil {
		return t.UTC(), nil
	}

	// Try numeric timestamp with different lengths
	if ts, err := strconv.ParseInt(timestampStr, 10, 64); err == nil {
		// Heuristic: if > 1e12, likely milliseconds; otherwise seconds
		if ts > 1e12 {
			return time.UnixMilli(ts).UTC(), nil
		} else {
			return time.Unix(ts, 0).UTC(), nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse timestamp '%s'. Supported formats: Unix timestamp (seconds/milliseconds), RFC3339 (2006-01-02T15:04:05Z), ISO date (2006-01-02), or date-time (2006-01-02 15:04:05)", timestampStr)
}

package utils

import (
	"time"
)

// ParseTimestamp parses a timestamp string in the format "yyyy-MM-dd'T'HH:mm:ss.SSS'Z'"
func ParseTimestamp(timestamp string) (time.Time, error) {
	return time.Parse(time.RFC3339, timestamp)
}

// FormatTimestamp formats a time.Time to the format "yyyy-MM-dd'T'HH:mm:ss.SSS'Z'"
func FormatTimestamp(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}

// ParseTimestampWithMS parses a timestamp string and ensures it's in UTC
func ParseTimestampWithMS(timestamp string) (time.Time, error) {
	parsed, err := ParseTimestamp(timestamp)
	if err != nil {
		return time.Time{}, err
	}
	return parsed.UTC(), nil
}

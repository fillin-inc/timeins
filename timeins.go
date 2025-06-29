// Package timeins provides a custom Time type that marshals to JSON with second precision.
// It wraps the standard time.Time and formats JSON output as "2006-01-02T15:04:05-07:00".
package timeins

import (
	"strconv"
	"time"
)

// ISO8601Format is the time format used for JSON serialization with second precision.
const ISO8601Format string = "2006-01-02T15:04:05-07:00"

// Time wraps time.Time to provide JSON marshaling with second precision.
// It embeds time.Time to inherit all standard time functionality while
// customizing JSON serialization behavior.
type Time time.Time

// Parse parses a time string in ISO8601 format and returns a timeins.Time.
// The expected format is "2006-01-02T15:04:05-07:00".
func Parse(value string) (Time, error) {
	tt, err := time.Parse(ISO8601Format, value)
	return Time(tt), err
}

// MarshalJSON implements the json.Marshaler interface.
// It formats the time in ISO8601 format with second precision.
// This method is automatically called by json.Marshal.
func (t Time) MarshalJSON() ([]byte, error) {
	// Pre-allocate buffer with exact size needed: 2 quotes + 25 chars for timestamp
	buf := make([]byte, 0, 27)
	buf = append(buf, '"')
	buf = time.Time(t).AppendFormat(buf, ISO8601Format)
	buf = append(buf, '"')
	return buf, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// It parses a JSON time string in ISO8601 format.
// This method is automatically called by json.Unmarshal.
func (t *Time) UnmarshalJSON(data []byte) error {
	// Use strconv.Unquote to properly handle JSON-escaped strings
	timeStr, err := strconv.Unquote(string(data))
	if err != nil {
		return &time.ParseError{
			Layout: ISO8601Format,
			Value:  string(data),
			LayoutElem: "quoted string",
			ValueElem: string(data),
			Message: "invalid JSON string",
		}
	}

	tt, err := time.Parse(ISO8601Format, timeStr)
	if err != nil {
		return err
	}
	*t = Time(tt)
	return nil
}

// String returns the time formatted in ISO8601 format with second precision.
// The format is "2006-01-02T15:04:05-07:00".
func (t Time) String() string {
	return time.Time(t).Format(ISO8601Format)
}

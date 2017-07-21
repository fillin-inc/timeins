package timeins

import (
	"strings"
	"time"
)

// F is time format for time.Parse
const F string = "2006-01-02T15:04:05-07:00"

// Time is base structure for timeins
type Time time.Time

// Parse parses the string and returns the timeins.Time type
func Parse(value string) (Time, error) {
	tt, err := time.Parse(F, value)
	return Time(tt), err
}

func (t Time) String() string {
	return time.Time(t).Format(F)
}

// MarshalJSON is a method used when converting timeins.Time type to JSON
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

// UnmarshalJSON is a method used when converting JSON to timeins.Time type
func (t *Time) UnmarshalJSON(data []byte) error {
	tt, err := time.Parse(F, strings.Trim(string(data), `"`))
	*t = Time(tt)
	return err
}

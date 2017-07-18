package timeins

import (
	"strings"
	"time"
)

const F string = "2006-01-02T15:04:05-07:00"

type Time time.Time

func Parse(value string) (Time, error) {
	tt, err := time.Parse(F, value)
	return Time(tt), err
}

func (t Time) String() string {
	return time.Time(t).Format(F)
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	tt, err := time.Parse(F, strings.Trim(string(data), `"`))
	*t = Time(tt)
	return err
}

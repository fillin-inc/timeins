package timeins

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		hasError bool
	}{
		{"2006-01-02T15:04:05-07:00", false},
		{"2016-10-20T12:32:02+09:00", false},
		{"invalid-date", true},
		{"2006-01-02", true},
		{"", true},
	}

	for i, test := range tests {
		tis, err := Parse(test.input)
		
		if test.hasError {
			if err == nil {
				t.Errorf("#%d should return error for input: %s", i, test.input)
			}
			continue
		}
		
		if err != nil {
			t.Errorf("#%d has error in timeins.Parse: %s", i, err)
			continue
		}

		tt, _ := time.Parse(ISO8601Format, test.input)
		if tt.UnixNano() != time.Time(tis).UnixNano() {
			t.Errorf(
				"#%d returned unexpected value(expected:%d actual:%d)",
				i,
				tt.UnixNano(),
				time.Time(tis).UnixNano(),
			)
		}
	}
}

func TestString(t *testing.T) {
	tests := []string{
		"2006-01-02T15:04:05-07:00",
		"2016-10-20T12:32:02+09:00",
	}

	for i, test := range tests {
		parsedTime, _ := time.Parse("2006-01-02T15:04:05-07:00", test)
		time := Time(parsedTime).String()

		if time != test {
			t.Errorf(
				"#%d timeins.String() return unexpected value(expected:%s actual:%s)",
				i,
				test,
				time,
			)
		}
	}
}

type marshalTest struct {
	time     string
	expected string
}

func TestMarshalJSON(t *testing.T) {
	tests := []marshalTest{
		{
			"2006-01-02T15:04:05-07:00",
			`"2006-01-02T15:04:05-07:00"`,
		},
		{
			"2016-10-20T12:32:02+09:00",
			`"2016-10-20T12:32:02+09:00"`,
		},
	}

	for i, test := range tests {
		tt, _ := time.Parse("2006-01-02T15:04:05-07:00", test.time)
		marshaled, _ := Time(tt).MarshalJSON()

		if string(marshaled) != test.expected {
			t.Errorf(
				"#%d MarshalJSON() returned unexpected value(expected:%s actual:%s)",
				i,
				test.expected,
				string(marshaled),
			)
		}
	}
}

type unmarshalTest struct {
	time     string
	expected string
	hasErr   bool
}

func TestUnmarshalJSON(t *testing.T) {
	tests := []unmarshalTest{
		{
			`"2006-01-02T15:04:05-07:00"`,
			`2006-01-02T15:04:05-07:00`,
			false,
		},
		{
			`"2016-10-20T12:32:02+09:00"`,
			`2016-10-20T12:32:02+09:00`,
			false,
		},
		// Error cases
		{
			`"invalid-date"`,
			``,
			true,
		},
		{
			`"2006-01-02"`,
			``,
			true,
		},
		{
			`2006-01-02T15:04:05-07:00`,
			``,
			true,
		},
		{
			`null`,
			``,
			true,
		},
	}

	for i, test := range tests {
		tis := Time{}
		err := tis.UnmarshalJSON([]byte(test.time))

		if test.hasErr {
			if err == nil {
				t.Errorf("#%d UnmarshalJSON() should return error", i)
			}
			// Skip string comparison for error cases
			continue
		}
		
		if err != nil {
			t.Errorf("#%d UnmarshalJSON() should not return error(err:%s)", i, err.Error())
			continue
		}

		if test.expected != tis.String() {
			t.Errorf(
				"#%d UnmarshalJSON() returned unexpected value(expected:%s actual:%s)",
				i,
				test.expected,
				tis.String(),
			)
		}
	}
}

func BenchmarkParse(b *testing.B) {
	str := "2017-07-16T07:10:20+00:00"
	for i := 0; i < b.N; i++ {
		_, err := Parse(str)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkString(b *testing.B) {
	now := Time(time.Now())
	for i := 0; i < b.N; i++ {
		_ = now.String()
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	now := Time(time.Now())
	for i := 0; i < b.N; i++ {
		_, err := now.MarshalJSON()
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	tt := Time{}
	tb := []byte(`"2017-07-16T07:10:20+09:00"`)
	for i := 0; i < b.N; i++ {
		err := tt.UnmarshalJSON(tb)
		if err != nil {
			b.Error(err)
		}
	}
}

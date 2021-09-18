package timeins

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	tests := []string{
		"2006-01-02T15:04:05-07:00",
		"2016-10-20T12:32:02+09:00",
	}

	for i, test := range tests {
		tis, err := Parse(test)
		if err != nil {
			t.Errorf("#%d has error in timeins.Parse: %s", i, err)
		}

		tt, _ := time.Parse(F, test)
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
	}

	for i, test := range tests {
		tis := Time{}
		err := tis.UnmarshalJSON([]byte(test.time))

		if test.hasErr && err == nil {
			t.Errorf("#%d MarshalJSON() should return error", i)
		} else if !test.hasErr && err != nil {
			t.Errorf("#%d MarshalJSON() should not return error(err:%s)", i, err.Error())
		}

		if test.expected != tis.String() {
			t.Errorf(
				"#%d MarshalJSON() returned unexpected value(expected:%s actual:%s)",
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
	tb := []byte("2017-07-16T07:10:20+09:00")
	for i := 0; i < b.N; i++ {
		err := tt.UnmarshalJSON(tb)
		if err != nil {
			b.Error(err)
		}
	}
}

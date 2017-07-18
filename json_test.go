package timeins

import (
	"encoding/json"
	"testing"
	"time"
)

type structToJsonTest struct {
	time     string
	expected string
}

type ts1 struct {
	CreatedAt Time `json:"created_at"`
}

func TestStructToJSON(t *testing.T) {
	tests := []structToJsonTest{
		{
			"2017-07-18T21:20:15+09:00",
			`{"created_at":"2017-07-18T21:20:15+09:00"}`,
		},
		{
			"2017-07-01T09:00:00+00:00",
			`{"created_at":"2017-07-01T09:00:00+00:00"}`,
		},
	}

	for i, test := range tests {
		tt, _ := time.Parse(F, test.time)
		ts := ts1{
			CreatedAt: Time(tt),
		}

		j, err := json.Marshal(ts)
		if err != nil {
			t.Errorf("#%d %s", i, err.Error())
		}

		if test.expected != string(j) {
			t.Errorf(
				"#%d timeins return unexpected json(expected:%s actual:%s)",
				i,
				test.expected,
				j,
			)
		}
	}
}

type jsonToStructTest struct {
	json     string
	expected string
}

func TestJsonToStruct(t *testing.T) {
	tests := []jsonToStructTest{
		{
			`{"created_at":"2017-07-18T21:20:15+09:00"}`,
			"2017-07-18T21:20:15+09:00",
		},
		{
			`{"created_at":"2017-07-01T09:00:00+00:00"}`,
			"2017-07-01T09:00:00+00:00",
		},
	}

	var ts ts1
	for i, test := range tests {
		err := json.Unmarshal([]byte(test.json), &ts)
		if err != nil {
			t.Errorf("%d %s", i, err.Error())
		}

		if test.expected != ts.CreatedAt.String() {
			t.Errorf(
				"#%d timeins return unexpected struct(expected:%s actual:%s)",
				i,
				test.expected,
				ts.CreatedAt.String(),
			)
		}
	}
}

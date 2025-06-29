package timeins_test

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/fillin-inc/timeins"
)

// ExampleTime demonstrates basic usage of timeins.Time in JSON marshaling.
func ExampleTime() {
	type Response struct {
		CreatedAt timeins.Time `json:"created_at"`
	}

	// Create a specific time for consistent output
	specificTime := time.Date(2023, 7, 15, 14, 30, 45, 0, time.FixedZone("JST", 9*60*60))
	r := Response{
		CreatedAt: timeins.Time(specificTime),
	}

	jsonData, _ := json.Marshal(r)
	fmt.Println(string(jsonData))

	// Output: {"created_at":"2023-07-15T14:30:45+09:00"}
}

// ExampleParse demonstrates parsing a time string.
func ExampleParse() {
	t, err := timeins.Parse("2023-07-15T14:30:45+09:00")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(t.String())

	// Output: 2023-07-15T14:30:45+09:00
}

// ExampleTime_UnmarshalJSON demonstrates JSON unmarshaling.
func ExampleTime_UnmarshalJSON() {
	type Response struct {
		CreatedAt timeins.Time `json:"created_at"`
	}

	jsonData := `{"created_at":"2023-07-15T14:30:45+09:00"}`
	var r Response

	err := json.Unmarshal([]byte(jsonData), &r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r.CreatedAt.String())

	// Output: 2023-07-15T14:30:45+09:00
}
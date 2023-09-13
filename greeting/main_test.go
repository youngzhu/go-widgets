package main

import (
	"testing"
	"time"
)

func TestGreeting_morning(t *testing.T) {
	testcases := []time.Time{
		time.Date(2023, 9, 13, 8, 0, 0, 0, time.UTC),
		time.Date(2023, 9, 13, 8, 1, 0, 0, time.UTC),
		time.Date(2023, 9, 13, 8, 59, 0, 0, time.UTC),
		time.Date(2023, 9, 13, 8, 59, 59, 0, time.UTC),
	}

	for _, testcase := range testcases {
		got := greeting(testcase)

		if goodMorning != got {
			t.Errorf("Expected %q, got %q instead\n", goodMorning, got)
		}
	}
}

func TestGreeting_afternoon(t *testing.T) {
	testcases := []time.Time{
		time.Date(2023, 9, 13, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 9, 13, 12, 1, 0, 0, time.UTC),
		time.Date(2023, 9, 13, 12, 59, 0, 0, time.UTC),
		time.Date(2023, 9, 13, 12, 59, 59, 0, time.UTC),
	}

	for _, testcase := range testcases {
		got := greeting(testcase)

		if goodAfternoon != got {
			t.Errorf("Expected %q, got %q instead\n", goodAfternoon, got)
		}
	}
}

func TestGreeting_evening(t *testing.T) {
	testcases := []time.Time{
		time.Date(2023, 9, 13, 18, 0, 0, 0, time.UTC),
		time.Date(2023, 9, 13, 18, 1, 0, 0, time.UTC),
		time.Date(2023, 9, 13, 18, 59, 0, 0, time.UTC),
		time.Date(2023, 9, 13, 18, 59, 59, 0, time.UTC),
	}

	for _, testcase := range testcases {
		got := greeting(testcase)

		if goodEvening != got {
			t.Errorf("Expected %q, got %q instead\n", goodEvening, got)
		}
	}
}

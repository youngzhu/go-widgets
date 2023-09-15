package main

import (
	"github.com/youngzhu/godate"
	"testing"
)

func TestReadBabies(t *testing.T) {
	want := 13

	readBabies()
	got := len(babies)

	if got != want {
		t.Errorf("%d babies, but got %d", want, got)
	}
}

func TestReadExtraDays_holidays(t *testing.T) {
	readExtraDays()
	//fmt.Println(extraHolidays)
	//fmt.Println(extraWorkdays)

	holidays := []godate.Date{
		mustDate(2023, 10, 1),
		mustDate(2023, 10, 2),
		mustDate(2023, 10, 3),
		mustDate(2023, 10, 4),
		mustDate(2023, 10, 5),
		mustDate(2023, 10, 6),
	}

	for _, holiday := range holidays {
		t.Run("", func(t *testing.T) {
			if !containsDate(extraHolidays, holiday) {
				t.Errorf("extra holidays should contains %s", holiday)
			}
		})
	}
}

func TestReadExtraDays_workdays(t *testing.T) {
	readExtraDays()
	//fmt.Println(extraHolidays)
	//fmt.Println(extraWorkdays)

	workdays := []godate.Date{
		mustDate(2023, 10, 7),
		mustDate(2023, 10, 8),
	}

	for _, workday := range workdays {
		t.Run("", func(t *testing.T) {
			if !containsDate(extraWorkdays, workday) {
				t.Errorf("extra workdays should contains %s", workday)
			}
		})
	}
}

func TestCountAllDays(t *testing.T) {
	testcases := []struct {
		date godate.Date
		days int
	}{
		{mustDate(2023, 9, 15), 5},
		{mustDate(2023, 9, 16), 5},
		{mustDate(2023, 9, 17), 5},
		{mustDate(2023, 9, 18), 6},
	}

	for _, testcase := range testcases {
		t.Run("", func(t *testing.T) {
			got := countAllDays(testcase.date)
			if got != testcase.days {
				t.Errorf("want %d, but got %d", testcase.days, got)
			}
		})
	}
}

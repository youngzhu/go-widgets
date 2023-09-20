package main

import (
	"github.com/youngzhu/godate"
	"strings"
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
		mustDate(2023, 9, 19),
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

func TestWhoIs(t *testing.T) {
	testcases := []struct {
		date godate.Date
		name string
	}{
		{mustDate(2023, 9, 11), "陈煜珉"},
		{mustDate(2023, 9, 12), "王书颖"},
		{mustDate(2023, 9, 12), "杭票"},
		{mustDate(2023, 9, 13), "王重言"},
		{mustDate(2023, 9, 14), "史昌浩"},
		{mustDate(2023, 9, 15), "史玮宸"},
		{mustDate(2023, 9, 15), "杨欣媛"},
		{mustDate(2023, 9, 18), "孙谦"},
		{mustDate(2023, 9, 18), "赵韵瑾"},
		{mustDate(2023, 10, 7), "王子歇"},
		{mustDate(2023, 10, 7), "朱诗玥"},
	}

	for _, testcase := range testcases {
		t.Run("", func(t *testing.T) {
			got := whoIs(testcase.date)
			if !strings.Contains(got, testcase.name) {
				t.Errorf("%q should contains %q", got, testcase.name)
			}
		})
	}
}

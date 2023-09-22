package main

import (
	"github.com/youngzhu/godate"
	"strings"
	"testing"
)

// 放到了 init 中，函数会重复执行
//func TestLoadBabies(t *testing.T) {
//	want := 13
//
//	loadBabies()
//	got := len(babies)
//
//	if got != want {
//		t.Errorf("%d babies, but got %d", want, got)
//	}
//}

func TestLoadExtraDays_holidays(t *testing.T) {
	loadExtraDays()
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

func TestLoadExtraDays_workdays(t *testing.T) {
	loadExtraDays()
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
		{mustDate(2023, 9, 21), ""},
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

func TestIsTurn(t *testing.T) {
	testcases := []struct {
		date godate.Date
		name string
		want bool
	}{
		{mustDate(2023, 9, 11), "陈煜珉", true},
		{mustDate(2023, 9, 12), "王书颖", true},
		{mustDate(2023, 9, 12), "杭票", true},
		{mustDate(2023, 9, 13), "王重言", true},
		{mustDate(2023, 9, 14), "史昌浩", true},
		{mustDate(2023, 9, 15), "史玮宸", true},
		{mustDate(2023, 9, 15), "杨欣媛", true},
		{mustDate(2023, 9, 18), "孙谦", true},
		{mustDate(2023, 9, 18), "赵韵瑾", true},
		{mustDate(2023, 9, 18), "王子歇", false},
		{mustDate(2023, 9, 21), "赵韵瑾", false},
		{mustDate(2023, 9, 21), "王子歇", false},
		{mustDate(2023, 10, 7), "王子歇", true},
		{mustDate(2023, 10, 7), "朱诗玥", true},
		{mustDate(2023, 10, 8), "朱诗玥", false},
	}

	for _, testcase := range testcases {
		t.Run("", func(t *testing.T) {
			got := isTurn(testcase.name, testcase.date)
			if got != testcase.want {
				t.Errorf("%s 值日的是[%s]：want:%v, got:%v",
					testcase.date, testcase.name, testcase.want, got)
			}
		})
	}
}

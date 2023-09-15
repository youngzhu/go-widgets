package main

import (
	"bufio"
	"github.com/youngzhu/godate"
	"os"
	"strconv"
	"strings"
)

/*
实现：有一组名单，轮流陪餐，遇到节假日顺延
*/

const (
	subject = "陪餐 11:20-12:20" // 邮件标题
)

func main() {
	//name := ""
}

func turn() int {
	return 0
}

var startDate = mustDate(2023, 9, 11) // 陪餐首次开始的时间

// 统计至截止日期（cutoffDate）上学的总天数（加上补班，减去节假日）
func countAllDays(cutoffDate godate.Date) int {
	count := 0

	cutoffDate, _ = cutoffDate.AddDay(1) // 当天也要放到for循环里比较

	it := startDate
	var err error
	for it.Before(cutoffDate.Time) {

		if !isWeekend(it) { // 工作日
			// 如果不是额外的假日，则+1
			if !containsDate(extraHolidays, it) {
				count++
			}
		} else { // 周末
			// 如果补班，则+1
			if containsDate(extraWorkdays, it) {
				count++
			}
		}

		it, err = it.AddDay(1)
		if err != nil {
			panic(err)
		}
	}

	return count
}

// todo 替换掉
func isWeekend(d godate.Date) bool {
	weekday := d.Weekday()
	return godate.Saturday == weekday || godate.Sunday == weekday
}

var (
	babies = make([]string, 0)

	extraHolidays = make([]godate.Date, 0)
	extraWorkdays = make([]godate.Date, 0)
)

func readExtraDays() {
	f, err := os.Open("extra_days.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		prefix := line[0]

		switch prefix {
		case '-':
			extraHolidays = append(extraHolidays, parseDate(line))
		case '+':
			extraWorkdays = append(extraWorkdays, parseDate(line))
		}
	}
}

func parseDate(s string) godate.Date {
	date := s[1:]
	ymd := strings.Split(date, ".")
	year, _ := strconv.Atoi(ymd[0])
	mon, _ := strconv.Atoi(ymd[1])
	day, _ := strconv.Atoi(ymd[2])

	result, err := godate.NewDateYMD(year, mon, day)
	if err != nil {
		panic(err)
	}
	return result
}

func readBabies() {
	f, err := os.Open("babies.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		babies = append(babies, scanner.Text())
	}
}

func mustDate(year, month, day int) (date godate.Date) {
	date, _ = godate.NewDateYMD(year, month, day)
	return
}

func containsDate(slice []godate.Date, date godate.Date) bool {
	for _, d := range slice {
		if date.IsTheSameDay(d) {
			return true
		}
	}
	return false
}

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
	startFrom = "2023-09-11" // 陪餐首次开始的时间

	subject = "陪餐 11:20-12:20" // 邮件标题
)

func main() {

}

func turn() int {
	return 0
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
		case '+':
			extraHolidays = append(extraHolidays, parseDate(line))
		case '-':
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

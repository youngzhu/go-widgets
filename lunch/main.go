package main

import (
	"bufio"
	"flag"
	"fmt"
	smail "github.com/youngzhu/go-smail"
	"github.com/youngzhu/godate"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
实现：有一组名单，轮流陪餐，遇到节假日顺延

提前一天邮件通知指定姓名的同学
*/

func init() {
	loadBabies()
	loadExtraDays()
}

// go run main.go -name 张三
func main() {
	name := flag.String("name", "", "学生的姓名")
	flag.Parse()
	log.Printf("学生姓名：%q\n", *name)
	if *name == "" {
		panic("请输入正确的学生姓名")
	}

	// 提前一天通知
	// 今天检查明天是否是该同学值班
	// 如果是，则发送提醒邮件
	tomorrow, _ := godate.Today().AddDay(1)
	log.Println(tomorrow)
	if result := isTurn(*name, tomorrow); result {
		subject := fmt.Sprintf("明日（%s)陪餐 11:20-12:20", tomorrow)
		err := smail.SendMail(subject, "")
		if err != nil {
			log.Println("邮件发送失败：" + err.Error())
		} else {
			log.Println("邮件发送成功")
		}
	}
}

func isTurn(name string, date godate.Date) bool {
	names := whoIs(date)
	return strings.Contains(names, name)
}

// 指定日期（date）该谁值班
func whoIs(date godate.Date) string {
	// 如果 date 是休息日，直接返回 ""
	if isOffDay(date) {
		return ""
	}
	count := countAllDays(date)
	idx := count%len(babies) - 1
	return babies[idx]
}

// 是否休息日
func isOffDay(date godate.Date) bool {
	//off1 := containsDate(extraHolidays, date)                     // 额外的休息日
	//off2 := isWeekend(date) && !containsDate(extraWorkdays, date) // 周末 且 没有补班的
	//
	//return off1 || off2

	if containsDate(extraHolidays, date) {
		return true
	}

	return date.IsWeekend() && !containsDate(extraWorkdays, date)
}

var startDate = godate.MustDate(2023, 9, 11) // 陪餐首次开始的时间

// 统计至截止日期（cutoffDate）上学的总天数（加上补班，减去节假日）
func countAllDays(cutoffDate godate.Date) int {
	count := 0

	cutoffDate, _ = cutoffDate.AddDay(1) // 当天也要放到for循环里比较

	it := startDate
	var err error
	for it.Before(cutoffDate) {

		if !it.IsWeekend() { // 工作日
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

var (
	babies = make([]string, 0)

	extraHolidays = make([]godate.Date, 0)
	extraWorkdays = make([]godate.Date, 0)
)

func loadExtraDays() {
	f, err := os.Open("extra_days.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	var dates []godate.Date
	for scanner.Scan() {
		line := scanner.Text()
		prefix := line[0]

		switch prefix {
		case '+', '-':
			dates = parseDate(line[1:])
		}

		switch prefix {
		case '-':
			extraHolidays = append(extraHolidays, dates...)
		case '+':
			extraWorkdays = append(extraWorkdays, dates...)
		}
	}
}

func parseDate(s string) []godate.Date {
	dates := make([]godate.Date, 0)

	if strings.Contains(s, "-") {
		// 日期区间
		split := strings.Split(s, "-")
		begin := split[0]
		end := split[1]
		beginDate := toDate(begin)
		endDate := toDate(end)
		for beginDate.Before(endDate) {
			dates = append(dates, beginDate)
			beginDate, _ = beginDate.AddDay(1)
		}
		dates = append(dates, endDate)
	} else {
		// 单个日期
		dates = append(dates, toDate(s))
	}

	return dates
}

func toDate(date string) godate.Date {
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

func loadBabies() {
	f, err := os.Open("babies.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		babies = append(babies, scanner.Text())
	}
}

func containsDate(slice []godate.Date, date godate.Date) bool {
	for _, d := range slice {
		if date.IsTheSameDay(d) {
			return true
		}
	}
	return false
}

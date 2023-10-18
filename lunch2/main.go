package main

import (
	"bufio"
	"fmt"
	smail "github.com/youngzhu/go-smail"
	"github.com/youngzhu/godate"
	"log"
	"os"
	"strconv"
	"strings"
)

/**
实现：每天邮件提醒陪餐轮值名单

1. 判断今天是否为工作日
	1.1 是，进入第2步
	1.2 否，结束
2. 计算从起始日到今天，共有多少工作日
3. 总天数 % 一轮下来经历的天数，即今天值班的索引 idx
4. babies[idx] 即今天值班的家长
*/

var startDate = godate.MustDate(2023, 9, 11) // 陪餐首次开始的时间

var (
	babies = make([]string, 0)

	extraHolidays = make([]godate.Date, 0)
	extraWorkdays = make([]godate.Date, 0)
)

func init() {
	loadBabies()
	loadExtraDays()
}

func main() {
	today := godate.Today()
	if isOffDay(today) {
		log.Println("休息！")
	} else {
		count := countAllDays(today)
		//idx := count%len(babies) - 1
		idx := (count - 1) % len(babies)
		subject := fmt.Sprintf("今日（%s)陪餐 11:20-12:20", today.FullStringCN())
		err := smail.SendMail(subject, babies[idx]+"家长")
		if err != nil {
			log.Println("邮件发送失败：" + err.Error())
		} else {
			log.Println("邮件发送成功")
		}
	}
}

// 统计至截止日期（cutoffDate）上学的总天数（加上补班，减去节假日）
func countAllDays(cutoffDate godate.Date) int {
	count := 0

	cutoffDate, _ = cutoffDate.AddDay(1) // 当天也要放到for循环里比较

	it := startDate
	var err error
	for it.Before(cutoffDate) {
		if !isOffDay(it) {
			count++
		}

		it, err = it.AddDay(1)
		if err != nil {
			panic(err)
		}
	}

	return count
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

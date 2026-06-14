// Package bocom 交通银行（Bank of Communications，简称BOCOM）
package bocom

import (
	"log"
	"reminder"
)

type Reminder struct{}

func (r Reminder) Remind() {
	checkin()
	salaryAward()
}

var today = reminder.Today

// 每月10号领20-10券
// 工资卡奖励
//
// 开薪星期五
// 每月（周五？）一次，领贴金券
func salaryAward() {
	if today.Day() == 1 {
		log.Println("交行工资卡奖励")
		dueOn, _ := today.AddDay(15)
		reminder.CreateTodo("交行储蓄卡-贴金券", dueOn, today)
	}
}

// 每月签到，小额的满减券
func checkin() {
	if today.Day() == 5 {
		log.Println("交行App签到")
		dueOn, _ := today.AddDay(10)
		reminder.CreateTodo("交行App签到", dueOn, today)
	}
}

// Package bocom 交通银行（Bank of Communications，简称BOCOM）
package bocom

import (
	"log"
	"reminder"
)

type Reminder struct{}

func (r Reminder) Remind() {
	checkin()
	//salaryAward()
}

var today = reminder.Today

// 每月10号领20-10券
// 工资卡奖励
func salaryAward() {
	if today.Day() == 10 {
		log.Println("交行工资卡奖励")
		dueOn, _ := today.AddDay(30)
		reminder.CreateTodo("交行储蓄卡20-10", dueOn, today)
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

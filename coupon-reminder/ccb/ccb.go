// Package ccb 建设银行
package ccb

import (
	"reminder"
)

type Reminder struct{}

func (r Reminder) Remind() {
	huiSaveMoney()
}

var today = reminder.Today

// 惠省钱
func huiSaveMoney() {
	if today.Day() == 3 {
		dueOn, _ := today.AddDay(10)
		reminder.CreateTodo("建行➡惠省钱", dueOn, today)
	}
}

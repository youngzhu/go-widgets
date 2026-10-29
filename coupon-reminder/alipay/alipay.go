// Package alipay 支付宝
package alipay

import (
	"reminder"
)

type Reminder struct{}

func (r Reminder) Remind() {
	commuterCard()
}

var today = reminder.Today

// 支付宝出行神卡
func commuterCard() {
	for _, d := range []int{1, 15} {
		if today.Day() == d {
			dueOn, _ := today.AddDay(10)
			reminder.CreateTodo("支付宝➡出行神卡", dueOn, today)
		}
	}
}

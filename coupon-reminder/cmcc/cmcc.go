// Package cmcc 中国移动，China Mobile Communications Corporation
package cmcc

import (
	"fmt"
	"log"
	"reminder"
)

type Reminder struct{}

func (r Reminder) Remind() {
	checkin()
	rightsMonthly()
}

var today = reminder.Today

// 每月1次权益超市
func rightsMonthly() {
	if today.Day() == 16 {
		log.Println("移动权益领取提醒")
		dueOn, _ := today.AddDay(10)
		reminder.CreateTodo("移动权益领取", dueOn, today)
	}
}

// 每月7次签到，下旬做，也不必连续
// 实现：每月15日，新建7条签到的TODO
func checkin() {
	if today.Day() == 19 {
		log.Println("移动签到")
		var content string
		dueOn, _ := today.AddDay(10)
		for i := 0; i < 7; i++ {
			content = fmt.Sprintf("移动签到%d/7", i+1)
			reminder.CreateTodo(content, dueOn, today)
		}
	}
}

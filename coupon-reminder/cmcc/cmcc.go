// Package cmcc 中国移动，China Mobile Communications Corporation
package cmcc

import (
	"fmt"
	"github.com/youngzhu/godate"
	"log"
	"reminder"
)

// 中国移动：每月1次权益超市

type CMCCReminder struct{}

func (r CMCCReminder) Remind() {
	checkin()
}

// 每月7次签到，下旬做，也不必连续
// 实现：每月15日，新建7条签到的TODO
func checkin() {
	today := godate.Today()
	if today.Day() == 15 {
		log.Println("移动签到")
		var content string
		dueOn, _ := today.AddDay(10)
		for i := 0; i < 7; i++ {
			content = fmt.Sprintf("移动签到%d/7", i+1)
			reminder.CreateTodo(content, dueOn.String(), today.String())
		}
	}
}

package eleme

import (
	"log"
	"reminder"
)

// 饿了么-话费特惠充权益50-1

type Reminder struct{}

func (r Reminder) Remind() {
	content := "饿了么-话费特惠充权益50-1"

	if today.Day() == 20 {
		log.Println(content)
		dueOn, _ := today.AddDay(5)
		reminder.CreateTodo(content, dueOn, today)
	}
}

var today = reminder.Today

package pingan

import (
	"log"
	"reminder"
)

// 平安好车主送的一些券，主要是兑换一些小商品

type Reminder struct{}

func (r Reminder) Remind() {
	if today.Day() == 11 {
		log.Println("平安好车主-用卡券")
		dueOn, _ := today.AddDay(15)
		reminder.CreateTodo("平安好车主-用卡券", dueOn.String(), today.String())
	}
}

var today = reminder.Today

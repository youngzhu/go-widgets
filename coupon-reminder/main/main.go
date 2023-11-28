package main

import (
	"log"
	"reminder"
	"reminder/cmcc"
)

// 将优惠券信息加入到Basecamp的Todo中，防止过期忘记

func init() {
	register("cmcc", cmcc.CMCCReminder{})
}

func main() {
	for k, v := range reminders {
		log.Println(k, "start...")
		v.Remind()
		log.Println(k, "DONE!")
	}
}

var reminders = make(map[string]reminder.Reminder)

func register(s string, r reminder.Reminder) {
	if _, exists := reminders[s]; exists {
		log.Println("Register Reminder:", s)
		reminders[s] = r
	}
}

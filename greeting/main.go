package main

import (
	"fmt"
	smail "github.com/youngzhu/go-smail"
	"time"
)

/*
实现：每天早、中、晚发送三封问候邮件

目的：
1. 测试邮件发送功能
2. 验证一个repository里是否可以配置多个action执行多个main函数
  答：可以
*/

const (
	goodMorning   = "早上好！"
	goodAfternoon = "中午好！"
	goodEvening   = "晚上好！"
)

func main() {
	msg := greeting(time.Now())
	if msg != "" {
		//fmt.Println("问候: " + msg)
		smail.SendMail("Greeting", msg)
	} else {
		fmt.Println("休息！勿扰！")
		smail.SendMail("勿扰 testing", "")
	}
}

func greeting(t time.Time) string {
	hour := t.Hour()
	switch hour {
	case 8:
		return goodMorning
	case 12:
		return goodAfternoon
	case 18:
		return goodEvening
	default:
		return ""
	}
}

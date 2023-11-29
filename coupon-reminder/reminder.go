package reminder

import (
	"github.com/youngzhu/go-basecamp"
	"github.com/youngzhu/go-smail"
	"github.com/youngzhu/godate"
	"time"
)

type Reminder interface {
	Remind()
}

const (
	projectName   = "MeTime"
	todoSetTitle  = "To-dos"
	todoListTitle = "券"
)

func CreateTodo(content string, dueOn, startsOn godate.Date) {
	todo := basecamp.Todo{
		Content:  content,
		DueOn:    dueOn.String(),
		StartsOn: startsOn.String(),
	}

	err := basecamp.CreateTodo(projectName, todoSetTitle, todoListTitle, todo)
	if err != nil {
		// 一个错，就会个个错，没必要继续了
		// 发邮件，然后结束
		smail.SendMail("优惠券添加失败！", err.Error())
		panic(err)
	}

	//
	time.Sleep(time.Millisecond * 500)
}

var (
	Today = godate.Today()
)

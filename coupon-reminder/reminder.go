package reminder

import (
	"github.com/youngzhu/go-basecamp"
	"github.com/youngzhu/go-smail"
	"github.com/youngzhu/godate"
)

type Reminder interface {
	Remind()
}

const (
	projectName   = "MeTime"
	todoSetTitle  = "To-dos"
	todoListTitle = "券"
)

func CreateTodo(content, dueOn, startsOn string) {
	todo := basecamp.Todo{
		Content:  content,
		DueOn:    dueOn,
		StartsOn: startsOn,
	}

	err := basecamp.CreateTodo(projectName, todoSetTitle, todoListTitle, todo)
	if err != nil {
		// 一个错，就会个个错，没必要继续了
		// 发邮件，然后结束
		smail.SendMail("优惠券添加失败！", err.Error())
		panic(err)
	}
}

var (
	Today = godate.Today()
)
package main

import (
	"fmt"
	"github.com/youngzhu/go-basecamp"
	"github.com/youngzhu/godate"
	"sort"
)

const (
	projectName   = "MeTime"
	todoSetTitle  = "To-dos"
	todoListTitle = "Loans"
)

func main() {

	var err error
	for _, todo := range createTodos(loans) {
		err = basecamp.AddTodo(projectName, todoSetTitle, todoListTitle, todo)
		if err != nil {
			panic(err)
		}
	}

}

type loanItem struct {
	name   string
	amount float32
	due    int // monthday
}

var loans = []loanItem{
	{"按揭", 3649, 1},
	{"车贷", 3651, 10},
	{"招行闪电贷1", 1956, 2},
	{"招行闪电贷2", 1881, 2},
	{"兴业现金分期", 2218, 14},
	{"浙商消费e贷", 3700, 29},
}

func createTodos(loans []loanItem) []basecamp.Todo {
	todos := make([]basecamp.Todo, len(loans))

	var totalAmount float32
	for _, loan := range loans {
		totalAmount += loan.amount
	}

	sort.Slice(loans, func(i, j int) bool {
		return loans[i].due < loans[j].due
	})

	today := godate.Today()
	for i, loan := range loans {
		dueOn := godate.MustDate(today.Year(), today.Month().IntValue(), loan.due)
		todos[i] = basecamp.Todo{
			Content: fmt.Sprintf("%s(%.0f%%) %.2f/%.2f", loan.name, loan.amount/totalAmount*100, loan.amount, totalAmount),
			DueOn:   dueOn.String(),
		}
	}

	return todos
}
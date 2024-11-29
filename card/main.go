package main

import (
	"github.com/youngzhu/go-basecamp"
	"github.com/youngzhu/godate"
	"log"
	"os"
)

// basecamp card table

func main() {
	today := godate.Today()
	// 每周五添加下周的，方便计划下周的工作安排
	if today.Weekday() != godate.Friday {
		log.Println("今天不是星期五")
		os.Exit(0)
	}

	monday := today.NextWorkday()
	friday, _ := monday.AddDay(4)

	title := monday.String() + " ~ " + friday.String()
	err := basecamp.CreateCard("Profession", "Card Table", "In progress",
		basecamp.Card{
			Title: title,
			DueOn: friday.String(),
		})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("%s, SUCCESS!!!", title)
	}
}

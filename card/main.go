package main

import (
	"github.com/youngzhu/go-basecamp"
	"github.com/youngzhu/godate"
	"log"
)

// basecamp card table

func main() {
	today := godate.Today()
	workdays := today.Workdays()
	monday := workdays[0]
	friday := workdays[4]

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

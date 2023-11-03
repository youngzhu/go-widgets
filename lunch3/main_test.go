package main

import (
	"fmt"
	"github.com/youngzhu/godate"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	tomorrow, _ := godate.Today().AddDay(1)

	base := tomorrow.Time
	start := base.Add(time.Hour * 11).Add(time.Minute * 20)
	end := start.Add(time.Hour)
	fmt.Println(base)
	fmt.Println(start)
	fmt.Println(end)
}

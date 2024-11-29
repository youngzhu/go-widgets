package main

import (
	"github.com/magiconair/properties/assert"
	"github.com/youngzhu/go-basecamp"
	"sort"
	"testing"
)

func TestCreateTodos(t *testing.T) {
	input := []loanItem{
		{"按揭", 4000, 1},
		{"车贷", 3700, 10},
	}

	want := []basecamp.Todo{
		{Content: "按揭(52%) 4000.00/7700.00", DueOn: "2024-11-01"},
		{Content: "车贷(48%) 3700.00/7700.00", DueOn: "2024-11-10"},
	}

	got := createTodos(input)

	if len(want) != len(got) {
		t.Errorf("Expected length %d, but got %d", len(want), len(got))
	}

	for i, item := range input {
		t.Run(item.name, func(t *testing.T) {
			if !todoEqual(want[i], got[i]) {
				t.Errorf("Expected %#v, \nbut got %#v", want[i], got[i])
			}
		})
	}

}

func todoEqual(a, b basecamp.Todo) bool {
	return a.Content == b.Content && a.DueOn == b.DueOn
}

func TestSort(t *testing.T) {
	loans := []loanItem{
		{"按揭", 0, 1},
		{"车贷", 0, 10},
		{"信用卡", 0, 3},
	}

	sort.Slice(loans, func(i, j int) bool {
		return loans[i].due < loans[j].due
	})

	//println(loans)
	assert.Equal(t, loans[0].due, 1)
	assert.Equal(t, loans[1].due, 3)
	assert.Equal(t, loans[2].due, 10)
}

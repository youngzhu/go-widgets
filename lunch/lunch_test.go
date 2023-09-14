package main

import "testing"

func TestReadBabies(t *testing.T) {
	want := 11

	babies, _ := readBabies()
	got := len(babies)

	if got != want {
		t.Errorf("%d babies, but got %d", want, got)
	}

}

package main

import "testing"

func TestArea(t *testing.T) {
	got := Area(6.0,12.0)
	want := 72.0
	if got != want {
		t.Errorf("got %.2f want %.2f",got,want)
	}
}

func TestPerimiter(t *testing.T) {
	got := Perimiter(10.0, 10.0)
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got,want)
	}
}

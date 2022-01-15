package main

import "testing"

type Rectangle struct {
	Width float64
	Height float64
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{6.0,12.0}
	got := Area(rectangle)
	want := 72.0
	if got != want {
		t.Errorf("got %.2f want %.2f",got,want)
	}
}

func TestPerimiter(t *testing.T) {
	rectangle := Rectangle{10.0,10.0}
	got := Perimiter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got,want)
	}
}

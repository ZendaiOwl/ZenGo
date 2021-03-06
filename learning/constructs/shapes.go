package main

import "math"

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base float64
	Height float64
}

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func Area(rectangle Rectangle) float64 {
	return (rectangle.Width * rectangle.Height)
}

func Perimiter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

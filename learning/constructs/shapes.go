package main

func Area(rectangle Rectangle) float64 {
	return (rectangle.Width * rectangle.Height)
}

func Perimiter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

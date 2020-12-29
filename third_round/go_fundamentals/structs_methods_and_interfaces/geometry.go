package main

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

//Perimeter function to calculate the perimeter of provided values
func Perimeter(rectangle Rectangle) float64 {
	res := 2 * (rectangle.Width + rectangle.Height)
	return res
}

//Area of a rectangle calculation
func Area(rectangle Rectangle) float64 {
	res := rectangle.Width * rectangle.Height
	return res
}

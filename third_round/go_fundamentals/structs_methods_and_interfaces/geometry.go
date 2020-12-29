package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

//Perimeter function to calculate the perimeter of provided values
func Perimeter(rectangle Rectangle) float64 {
	res := 2 * (rectangle.Width + rectangle.Height)
	return res
}

//Area of a rectangle calculation
func (r Rectangle) Area() float64 {
	return r.Width * r.Height

}

//Area of a circle calculation
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

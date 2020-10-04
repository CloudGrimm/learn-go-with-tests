package main

import "math"

//Rectangle struct to define its data
type Rectangle struct {
	Width  float64
	Height float64
}

//Circle struct to define its data
type Circle struct {
	Radius float64
}

// Triangle struct to define its data
type Triangle struct {
	Base   float64
	Height float64
}

// Shape inteface
type Shape interface {
	Area() float64
}

//Perimeter calculates area of rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

//Area calculates area of rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area method for circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area method for Triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

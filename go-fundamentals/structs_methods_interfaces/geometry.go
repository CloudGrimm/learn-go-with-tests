package structs_methods_interfaces

import (
	"fmt"
	"math"
)

// Introducing structs
type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base float64
	Height float64
}

// introducing interfaces
type Shape interface {
	Area() float64
}
//Function to calculate the perimeter of the given shape
func Perimeter(rectangle Rectangle) float64{
	return 2 * (rectangle.Width + rectangle.Height)
}

//Method to calculate area of the shape
func (r Rectangle) Area() float64{
	return r.Height * r.Width
}

func (c Circle) Area() float64{
	return math.Pi * (c.Radius * c.Radius)
}

func (t Triangle) Area() float64{
	return 0.5 * t.Base * t.Height
}

func main(){
	//input := Rectangle{10.0, 10.0}
	fmt.Println("Perimeter of the shape:  ")
	//fmt.Println("Area of the shape is: ", Area(input))
}
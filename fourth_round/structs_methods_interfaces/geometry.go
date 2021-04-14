package main

import "math"

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct{
	Radius float64
}
func Perimeter(rec Rectangle) (per float64){
	return 2 * (rec.Width + rec.Height)
}

func (rec Rectangle) Area() (area float64){
	area = rec.Width * rec.Height
	return
}

func (cir Circle) Area() (area float64){
	area = math.Pi * cir.Radius * cir.Radius
	return
}

type Shape interface {
	Area() float64
}
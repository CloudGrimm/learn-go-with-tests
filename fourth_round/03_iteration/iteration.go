package main

func Repeat(input string, x int) (output string){
	for i := 0; i<x; i++ {
		output += input
	}
	return output
}
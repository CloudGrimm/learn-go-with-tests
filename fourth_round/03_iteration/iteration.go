package main

func Repeat(input string) (output string){
	for x := 0; x<5; x++ {
		output += input
	}
	return output
}
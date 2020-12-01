package main

import "fmt"

//Hello function that returns english greeting
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("Tinashe"))
}

package main

import "fmt"

const englishHelloPrefix = "Hello, "

//Hello function that returns english greeting
func Hello(name string) string {
	if name == "" {
		return englishHelloPrefix + "World"
	} else {
		return englishHelloPrefix + name
	}

}

func main() {
	fmt.Println(Hello("Tinashe"))
}

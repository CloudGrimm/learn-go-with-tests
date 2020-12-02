package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanish = "Spanish"
const french = "French"
const frenchHelloPrefix = "Bonjour, "

//Hello function that returns english greeting
func Hello(name, lang string) string {
	if lang == spanish {
		return spanishHelloPrefix + name
	}
	if lang == french {
		return frenchHelloPrefix + name
	}
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Tinashe", ""))
}

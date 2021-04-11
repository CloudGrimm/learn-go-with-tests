package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanish = "Spanish"
const shona = "Shona"
const shonaHelloPrefix = "Zvirisei, "

func Hello(name, lang string) string {
	if name == ""{
		name = "World"
	}
	prefix := englishHelloPrefix

	switch lang{
	case spanish:
		prefix = spanishHelloPrefix
	case shona:
		prefix = shonaHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
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
	if lang == spanish {
		return spanishHelloPrefix + name
	}
	if lang == shona {
		return shonaHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
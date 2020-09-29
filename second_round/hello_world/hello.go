package main

import "fmt"

//writing constants
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const shonaHelloPrefix = "Hesi, "

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name
}
func greetingPrefix(lang string) (prefix string) {

	switch lang {
	case "Shona":
		prefix = shonaHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	name := "Tinashe"
	lang := "English"
	fmt.Println(Hello(name, lang))
}

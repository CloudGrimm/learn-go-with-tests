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

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string){
	switch lang{
	case spanish:
		prefix = spanishHelloPrefix
	case shona:
		prefix = shonaHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
func main() {
	fmt.Println(Hello("Chris", ""))
}
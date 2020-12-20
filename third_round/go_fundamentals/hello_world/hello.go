package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanish = "Spanish"
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const shona = "Shona"
const shonaHelloPrefix = "Madini, "
// comments

//Hello function that returns english greeting
func Hello(name, lang string) string {
	// if lang == spanish {
	// 	return spanishHelloPrefix + name
	// } else if lang == french {
	// 	return frenchHelloPrefix + name
	// }
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name
}
//comment   

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHelloPrefix
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
	fmt.Println(Hello("Tinashe", ""))
}

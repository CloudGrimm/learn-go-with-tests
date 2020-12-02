package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanish = "Spanish"
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const shona = "Shona"
const shonaHelloPrefix = "Madini, "

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

	prefix := englishHelloPrefix

	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case shona:
		prefix = shonaHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("Tinashe", ""))
}

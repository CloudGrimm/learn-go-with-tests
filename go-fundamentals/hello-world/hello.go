package main // packages are a way of grouping up related Go code together

import "fmt"

const englishHelloPrefix  = "Hello, "
const spanishHelloPrefix  = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	if name == ""{
		name = "World"
	}
	return greetingPrefix(lang) + name
	/**using if statements*/
	//if lang == "Spanish"{
	//	return spanishHelloPrefix + name
	//}
	//if lang == "English"{
	//	return englishHelloPrefix + name
	//}
	//if lang == "French"{
	//	return frenchHelloPrefix + name
	//}
	//return englishHelloPrefix + name

	///**Using the switch**/
	//prefix := englishHelloPrefix
	//
	//switch lang {
	//case "French":
	//	prefix = frenchHelloPrefix
	//case "Spanish":
	//	prefix = spanishHelloPrefix
	//}
	//return prefix + name
}

func greetingPrefix(lang string)(prefix string){
	switch lang {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	case "English":
		prefix = englishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
func main(){
	fmt.Println(Hello("Wilbrod", "Spanish"))
}
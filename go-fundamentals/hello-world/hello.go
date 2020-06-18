package main // packages are a way of grouping up related Go code together

import "fmt"

const englishHelloPrefix  = "Hello, "

func Hello(name string) string{
	if name != ""{
		return englishHelloPrefix + name
	}else{
		return englishHelloPrefix + "World"
	}

}
func main(){
	fmt.Println(Hello("Wilbrod"))
}
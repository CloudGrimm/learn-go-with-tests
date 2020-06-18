package main // packages are a way of grouping up related Go code together

import "fmt"

func Hello(name string) string{
	return "Hello, " + name
}
func main(){
	fmt.Println(Hello("Wilbrod"))
}
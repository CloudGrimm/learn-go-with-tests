// This is program is going to be exploring the anatomy of functions
// In general a function is a small piece of code that is dedicated to perform a task based on some input values
package main

import "fmt"

func dosomething() {
	fmt.Println("Hello world!")
}

// Function name convetion

// Function parameters

func greet(user string) {
	fmt.Println("Hello " + user)
}

//Return Values
func add(a, b int) int64 {
	c := a + b
	return int64(c)
}

//Multiple return values
func addMulti(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	dosomething()
	greet("Tinashe")
	res := add(1, 5)
	fmt.Println(res)
	addRes, multiRes := addMulti(2, 5)
	fmt.Println(addRes, multiRes)
	//using the blank identifier in Go
	_, multiRes2 := addMulti(1, 5)
	fmt.Println(multiRes2)
}

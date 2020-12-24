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

//Multiple return values with named return values
func addMulti(a, b int) (add, multi int) {
	add = a + b
	multi = a * b
	return
}

// Recursive function is that function that calls itself from inside the body.
// n! = n*(n-1)! where n > 0

func getFactorial(num int) int {
	if num > 1 {
		return num * getFactorial(num-1)
	} else {
		return 1
	}
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
	fmt.Println(getFactorial(2))
}

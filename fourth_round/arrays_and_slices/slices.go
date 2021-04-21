package main

import "fmt"

func main() {
	var s1 []int
	s2 := []int{1, 2, 3}
	//s3 := []int{4, 5, 6, 7}
	//s4 := []int{1, 2, 3}

	n1 := copy(s2, s1)
	fmt.Printf("n1=%d, s1=%v, s2=%v\n", n1, s1, s2)
    fmt.Println("s1 == nil", s1 == nil)
}
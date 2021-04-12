package main

func Sum(x []int) (sum int){
	for _, number := range x{
		sum += number
	}
	return sum
}
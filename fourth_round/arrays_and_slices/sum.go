package main

func Sum(x [5]int) (sum int){
	for i:=0; i < len(x); i++{
		sum += x[i]
	}
	return sum
}
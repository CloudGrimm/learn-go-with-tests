package main

func Sum(x []int) (sum int){
	for _, number := range x{
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int){

	for _, numbers := range numbersToSum{
		sums = append(sums, Sum(numbers))
	}
	return sums
}
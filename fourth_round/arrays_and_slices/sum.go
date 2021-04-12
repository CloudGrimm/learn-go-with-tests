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
	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int){
	for _, numbers := range numbersToSum{
		if len(numbers) == 0{
			sums = append(sums, 0)
		}else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}
	return
}
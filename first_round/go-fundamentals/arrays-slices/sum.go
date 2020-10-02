package arrays_slices

import "fmt"

func Sum(x []int) int{
	var summed int
	//for i:=0; i<3; i++{
	//	summed += x[i]
	//}

	/**Using range*/
	for _, number := range x{
		summed += number
	}
	return summed
}

func SumAll(numbersToSum ...[]int) []int{
	//lengthOfNumbers := len(numbersToSum)
	//sums := make([]int, lengthOfNumbers)
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numToSum ...[]int) []int{
	var sums []int
	for _, numbers := range numToSum{
		if len(numbers) ==0 {
			sums = append(sums, 0)
		}else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

func main(){
	x := []int{1,2,3}
	fmt.Println(Sum(x))
}

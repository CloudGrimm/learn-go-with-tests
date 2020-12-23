package main

//Sum function to add array of numbers
func Sum(input []int) int {
	res := 0
	// for i := 0; i < len(input); i++ {
	// 	res += input[i]
	// }
	for _, value := range input {
		res += value
	}

	return res
}

//SumAll function adds individual slices and returns their sums as a slice
func SumAll(numbersToSum ...[]int) (sums []int) {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)

	for _, value := range numbersToSum {
		sums = append(sums, Sum(value))
	}
	return
}

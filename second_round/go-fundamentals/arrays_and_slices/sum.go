package main

//Sum adds array values to get a total
func Sum(arr []int) int {
	var output int

	// for i := 0; i < len(arr); i++ {
	// 	output += arr[i]
	// }
	//refactored
	for _, i := range arr {
		output += i
	}
	return output
}

//SumAll adds various slices together
func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)
	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)

	// }
	//refactor
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

//SumAllTails is to add all numbers in a slice besides the head
func SumAllTails(tailsToSum ...[]int) []int {
	var tailsSum []int
	for _, numbers := range tailsToSum {
		if len(numbers) == 0 {
			tailsSum = append(tailsSum, 0)
		} else {
			tailsSum = append(tailsSum, Sum(numbers[1:]))
		}
	}
	return tailsSum
}

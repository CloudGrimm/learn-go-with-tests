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

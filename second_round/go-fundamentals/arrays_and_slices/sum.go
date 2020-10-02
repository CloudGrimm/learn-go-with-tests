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

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

func main(){
	x := []int{1,2,3}
	fmt.Println(Sum(x))
}

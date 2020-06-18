package arrays_slices

import (
	"testing"
)

func TestSum(t *testing.T){

	assertCorrectMessage := func(t *testing.T, expected, addition int, numbers []int){
		t.Helper()

		if expected != addition{
			t.Errorf("expected %d got %d given, %v", expected, addition, numbers)
		}
	}

	//t.Run("Collection of 3 numbers", func(t *testing.T) {
	//	numbers := []int{1,2,3}
	//
	//	addition := Sum(numbers)
	//	expected := 6
	//
	//	assertCorrectMessage(t, expected, addition, numbers)
	//})

	t.Run("Collection of any numbers", func(t *testing.T) {
		numbers := []int{1,2,3,4,5}

		addition := Sum(numbers)
		expected := 15

		assertCorrectMessage(t, expected, addition, numbers)
	})
}

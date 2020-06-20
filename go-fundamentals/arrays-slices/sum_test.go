package arrays_slices

import (
	"reflect"
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

func TestSumAll(t *testing.T){

	got := SumAll([]int{1,2}, []int{0,9})
	want := []int{3,9}

	if !reflect.DeepEqual(got,want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T){

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T){
		got := SumAllTails([]int{1,2,2}, []int{0,9})
		want := []int{4,9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0,9})
		want := []int{0,9}
		checkSums(t, got, want)
	})
}

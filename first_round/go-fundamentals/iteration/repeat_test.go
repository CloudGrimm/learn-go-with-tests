package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T){
	output := Repeat("a", 5)
	expected := "aaaaa"

	if output != expected{
		t.Errorf("expected %q got %q", expected, output)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("e", 5))
	// Output: eeeee
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
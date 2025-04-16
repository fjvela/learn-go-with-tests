package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

// Example functions are compiled whenever tests are executed.
// Because such examples are validated by the Go compiler, you can be confident your documentation's
//
//	examples always reflect current code behavior.
//
// https://go.dev/blog/examples
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

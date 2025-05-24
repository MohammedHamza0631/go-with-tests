package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	right := 4

	if sum != right {
		t.Errorf("expected '%d' but got '%d'", right, sum)
	}
}

// ExampleAdd shows how to use the Add function.
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

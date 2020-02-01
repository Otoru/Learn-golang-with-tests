package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	assertResult := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' and want '%q'.", got, want)
		}
	}

	t.Run("Sum 2 simple numbers", func(t *testing.T) {
		result := Add(2, 2)
		const expected int = 4
		assertResult(t, result, expected)
	})
}

func ExampleAdd() {
	result := Add(1, 5)
	fmt.Println(result)
	// Output: 6
}

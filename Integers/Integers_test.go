package integers

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestAdd(t *testing.T) {

	assertResult := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("Got '%q' and want '%q'.", got, want)
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

func BenchmarkAdd(b *testing.B) {
	firstChoice := rand.Intn(9)
	secondChoice := rand.Intn(9)
	for i := 0; i < b.N; i++ {
		Add(firstChoice, secondChoice)
	}
}

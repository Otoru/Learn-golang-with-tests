package iteration

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRepeat(t *testing.T) {

	assert := func(t *testing.T, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected '%q' and got '%q'.", repeated, expected)
		}
	}

	t.Run("Verify if Repeat function work", func(t *testing.T) {
		repeated := Repeat("a", 5)
		const expected string = "aaaaa"
		assert(t, repeated, expected)
	})

	t.Run("Verify if Repeat function work with another character", func(t *testing.T) {
		repeated := Repeat("b", 3)
		const expected string = "bbb"
		assert(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", rand.Intn(9))
	}
}

func ExampleRepeat() {
	result := Repeat("x", 3)
	fmt.Println(result)
	// Output: xxx
}

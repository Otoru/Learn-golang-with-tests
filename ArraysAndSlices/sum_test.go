package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	checkSum := func(t *testing.T, result, want int) {
		t.Helper()
		if result != want {
			t.Errorf("got '%d' and want '%d'.", result, want)
		}
	}

	t.Run("Sum collections of first size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Sum(numbers)
		const want int = 6
		checkSum(t, result, want)
	})

	t.Run("Sum collections of second size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		result := Sum(numbers)
		const want int = 10
		checkSum(t, result, want)
	})

	t.Run("Sum collections of third size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 6}
		result := Sum(numbers)
		const want int = 12
		checkSum(t, result, want)
	})
}

func TestSumAll(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int, size string) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%d' and want '%d' given %s array.", got, want, size)
		}
	}

	t.Run("Sum one array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 6}
		result := SumAll(numbers)
		var want []int = []int{12}
		checkSums(t, result, want, "one")
	})

	t.Run("Sum two array", func(t *testing.T) {
		result := SumAll([]int{1, 2, 3, 6}, []int{1, 2, 3, 6, 2})
		var want []int = []int{12, 14}
		checkSums(t, result, want, "two")
	})

	t.Run("Sum three array", func(t *testing.T) {
		result := SumAll([]int{1, 2, 3, 6}, []int{1, 2, 3, 6, 2}, []int{1, 2, 1, 6, 2})
		var want []int = []int{12, 14, 12}
		checkSums(t, result, want, "three")
	})

}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("Safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}

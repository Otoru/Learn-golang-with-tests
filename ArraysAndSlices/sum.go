package main

// Sum calculates the total from a slice of integers
func Sum(numbers []int) (sum int) {
	if len(numbers) == 0 {
		sum = 0
	} else {
		for _, number := range numbers {
			sum += number
		}
	}

	return
}

// SumAll function receives an array of arrays and returns an array with the
// sum of the children of each sub array of the given array.
func SumAll(numbersToSum ...[]int) []int {
	var length int = len(numbersToSum)
	sums := make([]int, length)

	for index, numbers := range numbersToSum {
		sums[index] = Sum(numbers)
	}

	return sums
}

// SumAllTails receives a collection of arrays of integers and for each one,
// returns a sum of its components, except the first.
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) > 0 {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, 0)
		}
	}

	return sums
}

func main() {}

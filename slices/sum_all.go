package slices

import (
	"github.com/DrxwDev/GoWithTest/arrays"
)

func SumAll(numbersToSum ...[]int) []int {
	// Valid Approach
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)

	// Better Approach
	var sums []int

	for _, numbers := range numbersToSum {
		// Valid Approach
		// sums[i] = arrays.Sum(numbers)

		// Better Approach
		sums = append(sums, arrays.Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))

	for i, num := range numbersToSum {
		if len(num) <= 1 {
			sums[i] = 0
		} else {
			sums[i] = arrays.Sum(num[1:])
		}
	}

	return sums
}

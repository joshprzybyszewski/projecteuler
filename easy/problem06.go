package easy

import (
	"fmt"
)

func SolveProblem6() string {
	/*
		The sum of the squares of the first ten natural numbers is,
			1^2 + 2^2 + ... + 10^2 = 385

		The square of the sum of the first ten natural numbers is,
			(1 + 2 + ... + 10)^2 = 55^2 = 3025

		Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is .
			3025 - 385 = 3025

		Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
	*/
	n := 100
	s1 := squareOfSum(n)
	s2 := sumOfSquares(n)
	ans := s1 - s2
	return fmt.Sprintf("%d = %d - %d", ans, s1, s2)
}

func sumOfSquares(n int) int {
	total := 0
	for k := 1; k <= n; k++ {
		total += (k * k)
	}
	return total
}

func squareOfSum(n int) int {
	total := 0
	for k := 1; k <= n; k++ {
		total += k
	}
	return total * total
}

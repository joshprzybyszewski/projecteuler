package easy

import (
	"fmt"
)

func SolveProblem28() {
	/*
		Starting with the number 1 and moving to the right in
		a clockwise direction a 5 by 5 spiral is formed as follows:

		21 22 23 24 25
		20  7  8  9 10
		19  6  1  2 11
		18  5  4  3 12
		17 16 15 14 13

		It can be verified that the sum of the numbers on the
		diagonals is 101.
		= 1 + (9 + 7 + 5 + 3) + (25 + 21 + 17 + 13)

		What is the sum of the numbers on the diagonals in a
		1001 by 1001 spiral formed in the same way?
	*/
	ans := getSumOfProblem28Spiral(1001)
	fmt.Printf("Problem 28 Answer: %v\n", ans)
}

func getSumOfProblem28Spiral(n int) int {
	if n%2 == 0 {
		// invalid input!
		return 0
	}
	sum := 0
	if n >= 1 {
		sum++
	}
	for i := 3; i <= n; i += 2 {
		square := i * i
		for c := 0; c < 4; c++ {
			sum += square
			square -= i - 1
		}
	}
	return sum
}

package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
)

func SolveProblem28() string {
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
	return fmt.Sprintf("%v", ans)
}

func getSumOfProblem28Spiral(n int) int {
	sum := 0
	for i := 1; i <= n; i += 2 {
		sum += mathUtils.Sum(getSpiralCorners(i))
	}
	return sum
}

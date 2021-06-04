package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
)

func SolveProblem30() {
	/*
		Surprisingly there are only three numbers that can
		be written as the sum of fourth powers of their digits:

		1634 = 1^4 + 6^4 + 3^4 + 4^4
		8208 = 8^4 + 2^4 + 0^4 + 8^4
		9474 = 9^4 + 4^4 + 7^4 + 4^4

		As 1 = 1^4 is not a sum it is not included.

		The sum of these numbers is 1634 + 8208 + 9474 = 19316.

		Find the sum of all the numbers that can be written
		as the sum of fifth powers of their digits.
	*/
	ans := solveProblem30()
	fmt.Printf("Problem 30 Answer: %v\n", ans)
}

func solveProblem30() int {
	nineRaisedFive := mathUtils.Raised(9, 5)
	max := nineRaisedFive * uint64(len(fmt.Sprintf("%d", nineRaisedFive)))
	for len(fmt.Sprintf("%d", max))*int(nineRaisedFive) < int(max) {
		max += nineRaisedFive
	}

	sum := 0
	for i := 2; i <= int(max); i++ {
		if isSumOfFifthPowersOfDigits(i) {
			sum += i
		}
	}

	return sum
}

func isSumOfFifthPowersOfDigits(
	i int,
) bool {
	if i <= 1 {
		return false
	}

	digits := mathUtils.StringToDigits(fmt.Sprintf("%d", i))
	for i, d := range digits {
		digits[i] = int(mathUtils.Raised(uint64(d), 5))
	}
	return mathUtils.Sum(digits) == i
}

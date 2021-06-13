package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
)

func SolveProblem52() string {
	/*
		It can be seen that the number, 125874, and its
		double, 251748, contain exactly the same digits,
		but in a different order.

		Find the smallest positive integer, x, such that
		2x, 3x, 4x, 5x, and 6x, contain the same digits.
	*/
	ans := getSmallestPositiveIntegerWithPermutationsUpTo(6)
	return fmt.Sprintf("%v", ans)
}

func getSmallestPositiveIntegerWithPermutationsUpTo(maxMult int) int {
	for i := 1; ; i++ {
		iDigits := mathUtils.ToDigits(i)
		all := true
		for m := 2; m <= maxMult; m++ {
			mxi := mathUtils.ToDigits(m * i)
			if !isPermutation(iDigits, mxi) {
				all = false
				break
			}
		}
		if all {
			return i
		}
	}
}

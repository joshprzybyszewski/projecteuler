package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
	"github.com/joshprzybyszewski/projecteuler/sequence"
)

func SolveProblem63() string {
	/*
		The 5-digit number, 16807=7^5, is also a fifth power. Similarly,
		the 9-digit number, 134217728=8^9, is a ninth power.

		How many n-digit positive integers exist which are also an nth power?
	*/
	ans := getProblem63Answer()
	return fmt.Sprintf("%v", ans)
}

func getProblem63Answer() int {
	maxDigitsToCheck := 18
	// 9^19, 9^20, and 9^21 are all too big for the default int space.
	// I found these with trial and error on google.
	// therefore, we start our total at 3.
	total := 3

	for digits := 1; digits <= maxDigitsToCheck; digits++ {
		seq := sequence.GetPower(digits)

		min := int(mathUtils.Raised(10, uint64(digits-1)))
		max := int(mathUtils.Raised(10, uint64(digits)))
		prev := total

		for n := uint(1); ; n++ {
			p := seq.GetNth(n)
			if p >= max {
				break
			}
			if p >= min {
				// log.Printf("%d digits: %d^%d = %d\n", digits, n, digits, p)
				total++
			}
		}
		if total == prev {
			// we would use this break with an unbounded for loop if our ints
			// could handle the whole space.
			// break
		}
	}

	return total
}

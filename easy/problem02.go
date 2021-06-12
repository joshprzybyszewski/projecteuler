package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/utils"
)

func SolveProblem2() string {
	// mathUtils.Sum(sequence.FibonaccisBelow(4000000))
	sum := sumOfEvenFibonaccisBelow(4000000)
	return fmt.Sprintf("%d", sum)
}

func sumOfEvenFibonaccisBelow(max int) int {
	a, b := 0, 1
	total := 0
	for {
		a, b = b, a+b
		if b >= max {
			return total
		}
		if utils.IsEven(b) {
			total += b
		}
	}
}

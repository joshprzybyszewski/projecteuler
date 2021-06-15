package easy

import (
	"fmt"
	"math/big"
)

func SolveProblem57() string {
	/*
		It is possible to show that the square root of two can be
		expressed as an infinite continued fraction.

		In the first one-thousand expansions, how many
		fractions contain a numerator with more digits than the denominator?
	*/
	ans := getProblem57Answer()
	return fmt.Sprintf("%v", ans)
}

func getProblem57Answer() int {
	total := 0

	num := big.NewInt(3)
	den := big.NewInt(2)

	for expansion := 1; expansion <= 1000; expansion++ {
		if len(num.String()) > len(den.String()) {
			total++
		}

		nextDen := big.NewInt(0).Add(num, den)
		nextNum := big.NewInt(0).Add(nextDen, den)

		den = nextDen
		num = nextNum
	}
	return total
}

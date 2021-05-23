package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/utils"
)

func SolveProblem14() {
	/*
		The following iterative sequence is defined for the set of positive integers:

		n → n/2 (n is even)
		n → 3n + 1 (n is odd)

		Using the rule above and starting with 13, we generate the following sequence:

		13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

		It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
		Although it has not been proved yet (Collatz Problem), it is thought that all starting
		numbers finish at 1.

		Which starting number, under one million, produces the longest chain?
	*/
	ans, ansLen := longestCollatzChain(1000000)
	fmt.Printf("Problem 14 Answer: %v (len of %d)\n", ans, ansLen)
}

func longestCollatzChain(maxStart int) (int, int) {
	best, bestLen := 0, 0
	for i := 1; i <= maxStart; i++ {
		myLen := collatzChainLen(i, 1)
		if myLen > bestLen {
			best = i
			bestLen = myLen
		}
	}
	return best, bestLen
}

func collatzChainLen(n int, steps int) int {
	if n <= 1 {
		return steps
	}
	if utils.IsEven(n) {
		return collatzChainLen(n/2, steps+1)
	}
	return collatzChainLen((3*n)+1, steps+1)
}

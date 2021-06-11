package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
	"github.com/joshprzybyszewski/projecteuler/primes"
)

func SolveProblem41() {
	/*
		What is the largest n-digit pandigital prime that exists?
	*/
	ans := getLargestPandigitalPrime()
	fmt.Printf("Problem 41 Answer: %v\n", ans)
}

func getLargestPandigitalPrime() int {
	// TODO this would solve it, but very slowly. Instead, I should
	// build all possible n-pandigital numbers, and check if they're prime.
	ps := primes.Below(8000000)

	for i := len(ps) - 1; i > 0; i-- {
		p := ps[i]
		digits := mathUtils.ToDigits(p)
		if mathUtils.IsNPandigital(digits, len(digits)) {
			return p
		}
	}
	return 0
}

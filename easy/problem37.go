package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
	"github.com/joshprzybyszewski/projecteuler/primes"
)

func SolveProblem37() {
	/*
		Find the sum of the only eleven primes that are both
		truncatable from left to right and right to left.
	*/
	all := getNTruncatablePrimes(11, 5000, nil)
	ans := mathUtils.Sum(all)
	fmt.Printf("Problem 37 Answer: %v = sum(%v)\n", ans, all)
}

func getNTruncatablePrimes(
	n, max int,
	found []int,
) []int {
	maxFound := 10
	if len(found) > 0 {
		maxFound = found[len(found)-1] + 1
	}
	ps := primes.Below(max)
	for _, p := range ps {
		if p > maxFound && isTruncatablePrime(p) {
			found = append(found, p)
		}
	}

	if len(found) >= n {
		return found[:n]
	}

	return getNTruncatablePrimes(n, max*10, found)
}

// assumes input is prime
func isTruncatablePrime(n int) bool {
	if n < 10 {
		return false
	}

	digits := mathUtils.ToDigits(n)

	for i := 0; i < len(digits); i++ {
		if digits[i]%2 == 0 {
			return false
		}
	}

	for i := 1; i < len(digits); i++ {
		n := spliceDigits(digits[:i], nil)
		if !primes.Is(n) {
			return false
		}

		n = spliceDigits(digits[len(digits)-i:], nil)
		if !primes.Is(n) {
			return false
		}
	}

	return true
}

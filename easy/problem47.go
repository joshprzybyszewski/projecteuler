package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/primes"
)

func SolveProblem47() string {
	/*
		The first two consecutive numbers to have two distinct prime factors are:

		14 = 2 × 7
		15 = 3 × 5

		The first three consecutive numbers to have three distinct prime factors are:

		644 = 2² × 7 × 23
		645 = 3 × 5 × 43
		646 = 2 × 17 × 19.

		Find the first four consecutive integers to have four distinct prime factors
		each. What is the first of these numbers?
	*/

	all := getConsecutiveNumbersWithNDistinctPrimeFactors(4)
	ans := all[0]
	return fmt.Sprintf("%v", ans)
}

func getConsecutiveNumbersWithNDistinctPrimeFactors(n int) []int {
	res := make([]int, 0, n)
	for i := 2; len(res) < n; i++ {
		myDistinct := getNumDistinctPrimeFactors(i)
		if myDistinct < n {
			res = res[:0]
			continue
		}
		res = append(res, i)
	}
	return res
}

func getNumDistinctPrimeFactors(n int) int {
	factors := primes.Factors(n)
	if len(factors) == 0 {
		return 0
	}

	nd := 1
	for i := 1; i < len(factors); i++ {
		if factors[i-1] != factors[i] {
			nd++
		}
	}
	return nd
}

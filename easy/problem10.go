package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/primes"
)

func SolveProblem10() {
	/*
		The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

		Find the sum of all the primes below two million.
	*/
	ans := sumOfPrimesBelow(2000000)
	fmt.Printf("Problem 10 Answer: %d\n", ans)
}

func sumOfPrimesBelow(max int) int {
	total := 0
	for _, p := range primes.Below(max) {
		total += p
	}
	return total
}

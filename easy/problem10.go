package easy

import (
	"fmt"
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
	total := 2
	for i := 3; i < max; i += 2 {
		if IsPrime(i) {
			total += i
		}
	}
	return total
}

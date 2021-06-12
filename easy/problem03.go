package easy

import (
	"fmt"
	"math"

	"github.com/joshprzybyszewski/projecteuler/primes"
)

func SolveProblem3() string {
	/*
		The prime factors of 13195 are 5, 7, 13 and 29.

		What is the largest prime factor of the number 600851475143 ?
	*/
	sum := largestPrimeFactor(600851475143)
	return fmt.Sprintf("%d", sum)
}

func largestPrimeFactor(n int) int {
	max := int(math.Floor(math.Sqrt(float64(n))))
	primesBelow := primes.Below(max)

	for i := len(primesBelow) - 1; i >= 0; i-- {
		p := primesBelow[i]
		if n%p == 0 {
			// it's a factor!
			return p
		}
	}
	return 1
}

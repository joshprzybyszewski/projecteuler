package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
	"github.com/joshprzybyszewski/projecteuler/primes"
)

func SolveProblem50() string {
	/*
		The prime 41, can be written as the sum of six consecutive primes:

		41 = 2 + 3 + 5 + 7 + 11 + 13

		This is the longest sum of consecutive primes that adds to a prime below one-hundred.

		The longest sum of consecutive primes below one-thousand that adds
		to a prime, contains 21 terms, and is equal to 953.

		Which prime, below one-million, can be written as the sum of the
		most consecutive primes?
	*/
	ans := getPrimeWithMostConsecutiveAdditivesBelow(1000000)
	return fmt.Sprintf("%v", ans)
}

func getPrimeWithMostConsecutiveAdditivesBelow(max int) int {
	ps := primes.Below(max)

	bestPrime := 0

	for n := 1; ; n++ {
		sum := mathUtils.Sum(ps[0:n])
		if sum >= max {
			break
		}

		if primes.Is(sum) {
			bestPrime = sum
			continue
		}

		for i := 1; ; i++ {
			sum = mathUtils.Sum(ps[i : n+i])
			if sum >= max {
				break
			}
			if primes.Is(sum) {
				bestPrime = sum
				break
			}
		}
	}

	return bestPrime
}

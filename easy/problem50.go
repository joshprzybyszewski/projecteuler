package easy

import (
	"fmt"

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

func getPrimeWithMostConsecutiveAdditivesBelowTrial2(max int) int {
	sum := 0
	lastSeen := 0
	for _, p := range primes.Below(max) {
		sum += p
		if sum >= max {
			break
		}
		if primes.Is(sum) {
			lastSeen = sum
		}
	}
	return lastSeen
}

func getPrimeWithMostConsecutiveAdditivesBelow(max int) int {
	bestP := 0
	var bestAdditives []int
	for _, p := range primes.Below(max) {
		additives := getConsecutivePrimesWithSum(p)
		if len(additives) > len(bestAdditives) {
			bestAdditives = additives
			bestP = p
			fmt.Printf("%d has %d additives\n", bestP, len(bestAdditives))
		}
	}
	fmt.Printf("%d has %d additives: %v\n", bestP, len(bestAdditives), bestAdditives)

	return bestP
}

func getConsecutivePrimesWithSum(p int) []int {
	if p <= 2 {
		// input must be greater than the first prime, 2
		return nil
	}

	ps := primes.Below(p)
	minI, maxI := 0, 1
	sum := ps[minI]

	for minI < maxI && maxI < len(ps) {
		next := ps[maxI]
		nextSum := sum + next
		switch {
		case nextSum == p:
			return ps[minI : maxI+1]
		case nextSum < p:
			sum = nextSum
			maxI++
		default:
			sum -= ps[minI]
			minI++
		}
	}

	return nil
}

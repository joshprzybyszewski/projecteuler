package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/primes"
)

func SolveProblem46() string {
	/*
		It was proposed by Christian Goldbach that every odd composite number can be written as the sum of a prime and twice a square.

		9 = 7 + 2×1^2
		15 = 7 + 2×2^2
		21 = 3 + 2×3^2
		25 = 7 + 2×3^2
		27 = 19 + 2×2^2
		33 = 31 + 2×1^2

		It turns out that the conjecture was false.

		What is the smallest odd composite that cannot be written as
		the sum of a prime and twice a square?
	*/
	ans := getGoldbachCounterExample()
	return fmt.Sprintf("%v", ans)
}

func getGoldbachCounterExample() int {
	for i := 3; ; i += 2 {
		if isGoldbachOtherConjectureCounterExample(i) {
			return i
		}
	}
}

func isGoldbachOtherConjectureCounterExample(n int) bool {
	if primes.Is(n) {
		return false
	}

	sq := 1 // squares
	s := 1  // odds
	for sq < n {
		diff := n - sq - sq
		if diff <= 0 {
			break
		}
		if primes.Is(diff) {
			// I can write n = diff + 2*sq
			return false
		}

		s += 2
		sq += s
	}

	return true
}

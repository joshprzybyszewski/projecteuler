package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/primes"
)

func SolveProblem7() string {
	/*
		By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
		we can see that the 6th prime is 13.

		What is the 10 001st prime number?
	*/
	ans := getNthPrime(10001)
	return fmt.Sprintf("%d", ans)
}

func getNthPrime(n int) int {
	below := primes.Below(n)
	if n < len(below) {
		return below[n]
	}

	nSeen := 0
	starting := 2
	if len(below) > 0 {
		nSeen = len(below) - 1
		starting = below[nSeen]
	}

	for i := starting; ; i++ {
		if primes.Is(i) {
			nSeen++
			if nSeen >= n {
				return i
			}
		}
	}
}

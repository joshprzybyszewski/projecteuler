package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
	"github.com/joshprzybyszewski/projecteuler/primes"
)

func SolveProblem35() {
	/*
		How many circular primes are there below one million?
	*/
	ans := getNumberOfCircularPrimesBelow(1000000)
	fmt.Printf("Problem 35 Answer: %v\n", ans)
}

func getNumberOfCircularPrimesBelow(max int) int {
	n := 0
	ps := primes.Below(max)
	for _, p := range ps {
		if isCircularPrime(p) {
			n++
		}
	}
	return n
}

// assumes the input is prime
func isCircularPrime(n int) bool {
	digits := mathUtils.ToDigits(n)

	for i := 1; i < len(digits); i++ {
		if digits[i] == 0 {
			return false
		}

		n := spliceDigits(digits[i:], digits[:i])
		if !primes.Is(n) {
			return false
		}
	}

	return true
}

func spliceDigits(a, b []int) int {
	ret := 0
	for _, d := range a {
		ret *= 10
		ret += d
	}
	for _, d := range b {
		ret *= 10
		ret += d
	}
	return ret
}

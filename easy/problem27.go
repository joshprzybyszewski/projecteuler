package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/primes"
)

func SolveProblem27() {
	/*
		Euler discovered the remarkable quadratic formula:
			n^2 + n + 41

		It turns out that the formula will produce 40 primes
		for the consecutive integer values 0 <= n <= 39. However,
		when n=40 is divisible by 41, and certainly when n=41 is
		clearly divisible by 41.

		The incredible formula n^2 - 79n + 1601 was discovered,
		which produces 80 primes for the consecutive values 0 <= n <= 79.
		The product of the coefficients, −79 and 1601, is −126479.

		Considering quadratics of the form:

		n^2 + an + b, where |a| < 1000 and |b| < 1000

		where |n| is the modulus/absolute value of n
		e.g. |11| = 11 and |-4|=4

		Find the product of the coefficients, a and b,
		for the quadratic expression that produces the
		maximum number of primes for consecutive values
		of n, starting with n=0.
	*/
	ans := getProblem27Answer()
	fmt.Printf("Problem 27 Answer: %v\n", ans)
}

func getProblem27Answer() int {

	max := 0
	prod := 0
	absMax := 999

	// we could just iterate through the prime numbers
	// (at least for b), but this works quick enough
	for a := absMax; a >= -absMax; a-- {
		for b := absMax; b >= -absMax; b-- {
			numC := getNumConsecutivePrimes(a, b)
			if numC > max {
				max = numC
				prod = a * b
			}
		}
	}

	return prod
}

func getNumConsecutivePrimes(
	a, b int,
) int {

	n := 0
	for {
		q := getQuadraticResult(n, 1, a, b)
		if q < 0 {
			q *= -1
		}
		if !primes.Is(q) {
			break
		}
		n++
	}

	return n
}

func getQuadraticResult(
	n int,
	a, b, c int,
) int {
	return a*(n*n) + b*n + c
}

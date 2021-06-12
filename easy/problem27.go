package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/primes"
)

func SolveProblem27() string {
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
	return fmt.Sprintf("%v", ans)
}

func getProblem27Answer() int {

	max := 0
	prod := 0

	check := func(a, b int) {
		numC := getNumConsecutivePrimes(a, b)
		if numC > max {
			max = numC
			prod = a * b
		}
	}

	// check through all pairs of primes below 1000
	p := primes.Below(1000)
	for _, a := range p {
		for _, b := range p {
			check(a, b)
			check(a, -b)
			check(-a, b)
			check(-a, -b)
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

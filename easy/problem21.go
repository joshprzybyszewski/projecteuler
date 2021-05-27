package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
)

func SolveProblem21() {
	/*
		Let d(n) be defined as the sum of proper
		divisors of n (numbers less than n which
		divide evenly into n).

		If d(a) = b and d(b) = a, where a ≠ b,
		then a and b are an amicable pair and
		each of a and b are called amicable numbers.

		For example, the proper divisors of 220 are
		1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110;
		therefore d(220) = 284. The proper divisors
		of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

		Evaluate the sum of all the amicable numbers under 10000.
	*/
	ans := sumOfAmicableNumbers(10000)
	fmt.Printf("Problem 21 Answer: %v\n", ans)
}

func sumOfAmicableNumbers(limit int) int {
	sum := 0
	for i := 1; i < limit; i++ {
		if isAmicableNumber(i) {
			sum += i
		}
	}
	return sum
}

func isAmicableNumber(n int) bool {
	s1 := sumOfDivisors(n)
	if n == s1 {
		return false
	}
	s2 := sumOfDivisors(s1)
	return n == s2
}

func sumOfDivisors(n int) int {
	return mathUtils.Sum(mathUtils.ProperDivisors(n))
}

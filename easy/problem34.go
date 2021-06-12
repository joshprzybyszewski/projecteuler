package easy

import (
	"fmt"
	"strconv"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
)

func SolveProblem34() string {
	/*
		145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

		Find the sum of all numbers which are equal to the sum
		of the factorial of their digits.
	*/
	ans := getProblem34Answer()
	return fmt.Sprintf("%v", ans)
}

func getProblem34Answer() int {
	sum := 0

	nineFactorial := mathUtils.Factorial(9)
	max := len(strconv.Itoa(nineFactorial)) * nineFactorial
	for i := 10; i < max; i++ {
		if isSumOfFactorialsOfDigits(i) {
			sum += i
		}
	}
	return sum
}

func isSumOfFactorialsOfDigits(n int) bool {
	if n < 10 {
		return false
	}
	digits := mathUtils.ToDigits(n)
	for i := range digits {
		digits[i] = mathUtils.Factorial(digits[i])
	}
	return mathUtils.Sum(digits) == n
}

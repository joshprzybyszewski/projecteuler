package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
	"github.com/joshprzybyszewski/projecteuler/primes"
)

func SolveProblem56() string {
	/*
		A googol (10^100) is a massive number: one followed by
		one-hundred zeros; 100^100 is almost unimaginably large:
		one followed by two-hundred zeros. Despite their size,
		the sum of the digits in each number is only 1.

		Considering natural numbers of the form, a^b,
		where a, b < 100, what is the maximum digital sum?
	*/
	ans := getProblem56Answer()
	return fmt.Sprintf("%v", ans)
}

func getProblem56Answer() int {
	max := 0
	for a := 1; a < 100; a++ {
		for b := 1; b < 100; b++ {
			sumOfDigits := getSumOfDigitsOfARaisedB(a, b)
			if sumOfDigits > max {
				max = sumOfDigits
			}
		}
	}
	return max
}

func getSumOfDigitsOfARaisedB(a, b int) int {
	factors := getAToTheBFactors(a, b)
	product := mathUtils.BigProduct(factors)
	digits := mathUtils.StringToDigits(product.String())
	return mathUtils.Sum(digits)
}

func getAToTheBFactors(a, b int) []int {
	base := primes.Factors(a)

	output := make([]int, 0, len(base)*b)
	for i := 0; i < b; i++ {
		output = append(output, base...)
	}
	return output
}

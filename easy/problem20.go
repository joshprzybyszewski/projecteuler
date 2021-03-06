package easy

import (
	"fmt"
	"sort"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
	"github.com/joshprzybyszewski/projecteuler/primes"
)

func SolveProblem20() string {
	/*
		Find the sum of the digits in the number 100!
	*/
	factors := getPrimeFactorialFactors(100)
	factors = remove10s(factors)
	product := mathUtils.BigProduct(factors)
	ans := mathUtils.Sum(mathUtils.StringToDigits(product.String()))
	return fmt.Sprintf("%v", ans)
}

func getPrimeFactorialFactors(bang int) []int {
	allFactors := make([]int, 0, bang)

	for d := 2; d <= bang; d++ {
		factors := primes.Factors(d)
		allFactors = append(allFactors, factors...)
	}

	return allFactors
}

func remove10s(factors []int) []int {
	sort.Ints(factors)
	for {
		i2 := indexOf(factors, 2)
		i5 := indexOf(factors, 5)
		if i2 == -1 || i5 == -1 {
			break
		}
		if i2 > i5 {
			i2, i5 = i5, i2
		}
		newFactors := make([]int, 0, len(factors)-2)
		newFactors = append(
			newFactors,
			factors[:i2]...,
		)
		newFactors = append(
			newFactors,
			factors[i2+1:i5]...,
		)
		newFactors = append(
			newFactors,
			factors[i5+1:]...,
		)
		factors = newFactors
	}
	return factors
}

func indexOf(s []int, target int) int {
	for i := range s {
		if s[i] == target {
			return i
		}
	}
	return -1
}

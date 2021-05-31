package easy

import (
	"fmt"
)

const (
	mantissaLimitMultiplier = 5
)

func SolveProblem26() {
	/*
		A unit fraction contains 1 in the numerator. The decimal
		representation of the unit fractions with denominators
		2 to 10 are given:

		1/2		= 	0.5
		1/3		= 	0.(3)
		1/4		= 	0.25
		1/5		= 	0.2
		1/6		= 	0.1(6)
		1/7		= 	0.(142857)
		1/8		= 	0.125
		1/9		= 	0.(1)
		1/10	= 	0.1

		Where 0.1(6) means 0.166666..., and has a 1-digit recurring
		cycle. It can be seen that 1/7 has a 6-digit recurring cycle.

		Find the value of d < 1000 for which 1/d contains the longest
		recurring cycle in its decimal fraction part.
	*/
	ans := getMaxCycleLengthOfDivisors(1000)
	fmt.Printf("Problem 26 Answer: %v\n", ans)
}

func getMaxCycleLengthOfDivisors(
	limit int,
) int {
	max := 0
	var bestCycle []int
	for d := 2; d <= limit; d++ {
		_, cycle := getDecimalMantissa(d)
		if len(cycle) > len(bestCycle) {
			bestCycle = cycle
			max = d
		}
	}
	return max
}

const (
	maxDecimalMantissaLenght = 10000
)

func getDecimalMantissa(
	divisor int,
) (prefix, cycle []int) {

	from := make([]int, 0, 8)
	res := make([]int, 0, 8)

	rem := 10
	for len(res) < maxDecimalMantissaLenght {
		for rem < divisor {
			from = append(from, rem)
			res = append(res, 0)
			rem *= 10
		}

		from = append(from, rem)
		res = append(res, rem/divisor)
		rem = rem % divisor
		if rem == 0 {
			return res, nil
		}
		rem *= 10

		lastFrom := lookForCycle(from, rem)
		if lastFrom >= 0 {
			return res[:lastFrom], res[lastFrom:]
		}
	}

	return res, nil
}

func lookForCycle(
	from []int,
	rem int,
) int {
	for i := len(from) - 1; i >= 0; i-- {
		if from[i] == rem {
			return i
		}
	}
	return -1
}

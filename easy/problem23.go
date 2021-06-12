package easy

import (
	"fmt"
	"sort"
)

type perfectNumberState int8

const (
	perfect   perfectNumberState = 1
	abundant  perfectNumberState = 2
	deficient perfectNumberState = 3
)

func SolveProblem23() string {
	/*
		A perfect number is a number for which the sum of its proper divisors
		is exactly equal to the number. For example, the sum of the proper
		divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that
		28 is a perfect number.

		A number n is called deficient if the sum of its proper divisors is
		less than n and it is called abundant if this sum exceeds n.

		As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the
		smallest number that can be written as the sum of two abundant
		numbers is 24. By mathematical analysis, it can be shown that all
		integers greater than 28123 can be written as the sum of two
		abundant numbers. However, this upper limit cannot be reduced any
		further by analysis even though it is known that the greatest number
		that cannot be expressed as the sum of two abundant numbers is less
		than this limit.

		Find the sum of all the positive integers which cannot be
		written as the sum of two abundant numbers.
	*/
	limit := 28124
	abundantNumbers := allAbundantNumbersBelow(limit)
	ans := getSumOfAllNonAbundantSumNumbersBelow(limit, abundantNumbers)
	return fmt.Sprintf("%v", ans)
}

func getSumOfAllNonAbundantSumNumbersBelow(
	limit int,
	abundantNumbers []int,
) int {
	sum := 0

	for i := 1; i < limit; i++ {
		if hasTwoElementsWithSum(i, abundantNumbers) {
			continue
		}
		sum += i
	}

	return sum
}

func hasTwoElementsWithSum(sum int, elems []int) bool {
	for _, a := range elems {
		if a > sum {
			break
		}
		target := sum - a

		i := sort.Search(len(elems), func(i int) bool {
			return target <= elems[i]
		})
		if i < len(elems) && elems[i] == target {
			return true
		}
	}
	return false
}

func allAbundantNumbersBelow(limit int) []int {
	res := make([]int, 0, 16)

	for i := 1; i < limit; i++ {
		if getNumberState(i) == abundant {
			res = append(res, i)
		}
	}

	return res
}

func getNumberState(n int) perfectNumberState {
	s := sumOfProperDivisors(n)
	switch {
	case s > n:
		return abundant
	case s == n:
		return perfect
	default:
		return deficient
	}
}

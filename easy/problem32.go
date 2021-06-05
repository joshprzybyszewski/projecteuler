package easy

import (
	"fmt"
	"sort"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
)

func SolveProblem32() {
	/*
		We shall say that an n-digit number is pandigital if it makes
		use of all the digits 1 to n exactly once; for example, the
		5-digit number, 15234, is 1 through 5 pandigital.

		The product 7254 is unusual, as the identity, 39 × 186 = 7254,
		containing multiplicand, multiplier, and product is 1 through 9 pandigital.

		Find the sum of all products whose multiplicand/multiplier/product
		identity can be written as a 1 through 9 pandigital.
	*/
	ans := getSumOfAllPandigitalNumbers()
	fmt.Printf("Problem 32 Answer: %v\n", ans)
}

func getSumOfAllPandigitalNumbers() int {
	products := make(map[int]struct{}, 32)

	tenDigits := 1000000000
	for d1 := 1; d1 < tenDigits; d1++ {
		for d2 := d1; d2 < tenDigits; d2++ {
			p := d1 * d2
			if p >= tenDigits {
				break
			}
			digits := allDigits(p, d1, d2)
			if isPandigital(digits) {
				products[p] = struct{}{}
			}
		}
	}

	sum := 0

	for p := range products {
		sum += p
	}

	return sum
}

func allDigits(s ...int) []int {
	all := make([]int, 0, 32)
	for _, e := range s {
		all = append(all, mathUtils.ToDigits(e)...)
	}
	return all
}

func isPandigital(all []int) bool {
	if len(all) != 9 {
		return false
	}
	sort.Ints(all)

	for i := range all {
		if all[i] != i+1 {
			return false
		}
	}

	return true
}

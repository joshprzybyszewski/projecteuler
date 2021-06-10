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

	// I'm not clever; I stole this from the internet.
	// my solution ran in about 5 minutes. By being
	// clever like this and filtering down, it's down
	// to 56ms
	runFrom := func(min1, max1, min2, max2 int) {
		for d1 := min1; d1 < max1; d1++ {
			for d2 := min2; d2 < max2; d2++ {
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
	}

	oneDigits := 1
	twoDigits := 10
	threeDigits := 100
	fourDigits := 1000
	fiveDigits := 10000

	runFrom(
		oneDigits, twoDigits,
		fourDigits, fiveDigits,
	)
	runFrom(
		twoDigits, threeDigits,
		threeDigits, fourDigits,
	)

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
	return isNPandigital(all, 9)
}

func isNPandigital(all []int, n int) bool {
	if len(all) != n {
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

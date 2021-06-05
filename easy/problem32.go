package easy

import (
	"fmt"
	"sort"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
	"github.com/joshprzybyszewski/projecteuler/primes"
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
	sum := 0

	for product := 1; product < 1000000000; product++ {
		if isProductPandigitalWithMultipliers(product) {
			sum += product
		}
	}

	return sum
}

func isProductPandigitalWithMultipliers(product int) bool {
	if hasDuplicateDigitsOrZeros(product) {
		return false
	}

	// we're counting on Divisors returning pairs of divisors
	// so that the entries at i and i+1 multiply to get product.
	// we're also betting on the first two entries being 1 and n
	divisors := mathUtils.GetDivisorsFromPrimeFactors(
		product,
		primes.Factors(product),
	)
	for i := 2; i+1 < len(divisors); i += 2 {
		if isPandigital(
			product,
			divisors[i],
			divisors[i+1],
		) {
			return true
		}
	}
	return false
}

func isPandigital(s ...int) bool {
	all := make([]int, 0, 9)
	for _, e := range s {
		all = append(all, mathUtils.ToDigits(e)...)
	}
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

func hasDuplicateDigitsOrZeros(e int) bool {
	digits := mathUtils.ToDigits(e)
	sort.Ints(digits)
	if digits[0] == 0 {
		return true
	}

	for i := 0; i < len(digits)-1; i++ {
		if digits[i] == digits[i+1] {
			return true
		}
	}

	return false
}

package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
	"github.com/joshprzybyszewski/projecteuler/primes"
)

func SolveProblem33() {
	/*
		The fraction 49/98 is a curious fraction, as an inexperienced
		mathematician in attempting to simplify it may incorrectly
		believe that 49/98 = 4/8, which is correct, is obtained by
		canceling the 9s.

		We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

		There are exactly four non-trivial examples of this type of
		fraction, less than one in value, and containing two digits
		in the numerator and denominator.

		If the product of these four fractions is given in its lowest
		common terms, find the value of the denominator.
	*/
	ans := getProblem33Answer()
	fmt.Printf("Problem 33 Answer: %v\n", ans)
}

func getProblem33Answer() int {
	ns, ds := getAllTrickyFractions()

	nPrimeFactors := make([]int, 0, len(ns))
	for _, n := range ns {
		nPrimeFactors = append(nPrimeFactors, primes.Factors(n)...)
	}

	dPrimeFactors := make([]int, 0, len(ds))
	for _, d := range ds {
		dPrimeFactors = append(dPrimeFactors, primes.Factors(d)...)
	}

	for di := 0; di < len(dPrimeFactors); di++ {
		shouldRemove := false
		for ni, n := range nPrimeFactors {
			if n == dPrimeFactors[di] {
				shouldRemove = true
				nPrimeFactors = append(nPrimeFactors[:ni], nPrimeFactors[ni+1:]...)
				break
			}
		}
		if shouldRemove {
			dPrimeFactors = append(dPrimeFactors[:di], dPrimeFactors[di+1:]...)
			di--
		}
	}

	return mathUtils.Product(dPrimeFactors)
}

func getAllTrickyFractions() (nums, dens []int) {
	for n := 11; n < 100; n++ {
		for d := n + 1; d < 100; d++ {
			if isTrickyFraction(n, d) {
				nums = append(nums, n)
				dens = append(dens, d)
			}
		}
	}
	return nums, dens
}

func isTrickyFraction(
	numerator,
	denominator int,
) bool {
	if numerator%10 == 0 || denominator%10 == 0 {
		// trivial
		return false
	}

	f := float64(numerator) / float64(denominator)

	nDigits := mathUtils.ToDigits(numerator)
	dDigits := mathUtils.ToDigits(denominator)

	var cmp float64
	if nDigits[0] == dDigits[0] {
		cmp = float64(nDigits[1]) / float64(dDigits[1])
	} else if nDigits[0] == dDigits[1] {
		cmp = float64(nDigits[1]) / float64(dDigits[0])
	} else if nDigits[1] == dDigits[0] {
		cmp = float64(nDigits[0]) / float64(dDigits[1])
	} else if nDigits[1] == dDigits[1] {
		cmp = float64(nDigits[0]) / float64(dDigits[0])
	}

	return cmp == f
}

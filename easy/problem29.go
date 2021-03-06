package easy

import (
	"fmt"
	"sort"

	"github.com/joshprzybyszewski/projecteuler/primes"
	"github.com/joshprzybyszewski/projecteuler/utils"
)

func SolveProblem29() string {
	/*
		How many distinct terms are in the sequence generated by a^b for
		2 ≤ a ≤ 100 and 2 ≤ b ≤ 100?
	*/
	ans := getNumDistinctTermsForARaisedB(2, 100, 2, 100)
	return fmt.Sprintf("%v", ans)
}

type problem29Params struct {
	minA, maxA int
	minB, maxB int
}

func getNumDistinctTermsForARaisedB(
	minA, maxA,
	minB, maxB int,
) int {

	params := problem29Params{
		minA: minA,
		maxA: maxA,
		minB: minB,
		maxB: maxB,
	}

	numDistinct := 0
	aNumbers := make([]numberWithFactors, 0, maxA-minA+1)

	numEntriesPerA := maxB - minB + 1

	for a := minA; a <= maxA; a++ {
		af := numberWithFactors{
			base:         a,
			primeFactors: primes.Factors(a),
		}
		aNumbers = append(aNumbers, af)

		subsets := getSubsets(af, aNumbers[:len(aNumbers)-1])

		numDistinct += numEntriesPerA

		if len(subsets) == 0 {
			continue
		}

		duplicates := 0

		for b := minB; b <= maxB; b++ {
			if isAlreadyIncluded(af, b, subsets, params) {
				duplicates++
			}
		}

		numDistinct -= duplicates
	}

	return numDistinct
}

func isAlreadyIncluded(
	a numberWithFactors,
	b int,
	subsets []numberWithFactors,
	params problem29Params,
) bool {

	aRaisedBFactors := multiplySlice(a.primeFactors, b)

	for _, subset := range subsets {

		if len(aRaisedBFactors)%len(subset.primeFactors) != 0 {
			continue
		}

		mult := len(aRaisedBFactors) / len(subset.primeFactors)
		if mult < params.minB || mult > params.maxB {
			continue
		}

		if equalIntSlices(
			multiplySlice(subset.primeFactors, mult),
			aRaisedBFactors,
		) {
			return true
		}
	}

	return false
}

func multiplySlice(
	s []int,
	m int,
) []int {
	res := make([]int, 0, len(s)*m)
	for i := 0; i < m; i++ {
		res = append(res, s...)
	}

	sort.Ints(res)
	return res
}

func equalIntSlices(
	a, b []int,
) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

type numberWithFactors struct {
	primeFactors []int
	base         int
}

func getSubsets(
	af numberWithFactors,
	others []numberWithFactors,
) []numberWithFactors {
	res := make([]numberWithFactors, 0, 2)

	for _, other := range others {
		if utils.IsSubset(
			af.primeFactors,
			other.primeFactors,
		) {
			res = append(res, other)
		}
	}

	return res
}

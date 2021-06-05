package mathUtils

import (
	"math"
	"sort"
)

// Divisors returns a slice of numbers that evenly divide into n
// Both 1 and n are included.
func Divisors(n int) []int {
	divisors := make([]int, 0, 4)
	max := int(math.Floor(math.Sqrt(float64(n))))
	for i := 1; i <= max; i++ {
		if n%i != 0 {
			continue
		}

		divisors = append(divisors, i)

		d := n / i
		if d != i {
			divisors = append(divisors, d)
		}
	}
	return divisors
}

// ProperDivisors are numbers less than n which divide evenly into n
func ProperDivisors(n int) []int {
	all := Divisors(n)
	proper := make([]int, 0, len(all))

	for _, d := range all {
		if d < n {
			proper = append(proper, d)
		}
	}

	return proper
}

// Divisors returns a slice of numbers that evenly divide into n
// Both 1 and n are included.
func GetDivisorsFromPrimeFactors(
	n int,
	primeFactors []int,
) []int {

	fc := newFactorCache(n)
	fc.getUniqueDivisorsFromPrimeFactors(1, primeFactors)
	return fc.divisors()
}

type factorCache struct {
	pairs    map[int]int
	smallers []int

	n int
}

func newFactorCache(n int) *factorCache {
	return &factorCache{
		n:        n,
		pairs:    make(map[int]int, 8),
		smallers: make([]int, 0, 8),
	}
}

func (fc *factorCache) addDivisor(d int) {
	if _, ok := fc.pairs[d]; ok {
		return
	}

	other := fc.n / d

	if other < d {
		if _, ok := fc.pairs[other]; ok {
			return
		}
		d, other = other, d
	}

	fc.pairs[d] = other
	fc.smallers = append(fc.smallers, d)
}

func (fc *factorCache) getUniqueDivisorsFromPrimeFactors(
	curDivisor int,
	otherFactors []int,
) {
	fc.addDivisor(curDivisor)

	for i, f := range otherFactors {
		newDivisor := curDivisor * f
		if fc.n%newDivisor == 0 {
			fc.getUniqueDivisorsFromPrimeFactors(
				curDivisor*f,
				otherFactors[i+1:],
			)
		}
	}
}

func (fc *factorCache) divisors() []int {
	sort.Ints(fc.smallers)
	res := make([]int, 0, 2*len(fc.smallers))
	for _, d := range fc.smallers {
		res = append(res, d)
		if other := fc.pairs[d]; other != d {
			res = append(res, other)
		}
	}
	return res
}

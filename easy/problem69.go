package easy

import (
	"fmt"
	"log"
	"sort"

	"github.com/joshprzybyszewski/projecteuler/primes"
	"github.com/joshprzybyszewski/projecteuler/utils"
)

func SolveProblem69() string {
	/*
		Euler's Totient function, φ(n) [sometimes called the phi function],
		is used to determine the number of numbers less than n which are
		relatively prime to n. For example, as 1, 2, 4, 5, 7, and 8, are
		all less than nine and relatively prime to nine, φ(9)=6.

		n	Relatively Prime	φ(n)	n/φ(n)
		2	1					1		2
		3	1,2					2		1.5
		4	1,3					2		2
		5	1,2,3,4				4		1.25
		6	1,5					2		3
		7	1,2,3,4,5,6			6		1.1666...
		8	1,3,5,7				4		2
		9	1,2,4,5,7,8			6		1.5
		10	1,3,7,9				4		2.5

		It can be seen that n=6 produces a maximum n/φ(n) for n ≤ 10.

		Find the value of n ≤ 1,000,000 for which n/φ(n) is a maximum.
	*/
	ans := getMaxNOverTotientWithMaxN(1000000)
	return fmt.Sprintf("%v", ans)
}

func getMaxNOverTotientWithMaxN(maxN int) int {
	// full disclosure: I came up with this solution after doing it the hard
	// way below.
	v := 1
	for _, p := range primes.Below(100) {
		nextV := v * p
		if nextV > maxN {
			return v
		}
		v = nextV
	}

	// the following takes about 2 minutes, and it reveals the sequence 1, 2, 6, 30, 210, 2310, 30030, 510510.
	var max float64
	var bestN int
	for n := 1; n <= maxN; n++ {
		v := getNOverTotient(n)
		if v > max {
			log.Printf("for n=%6d, n/totient(n) = %8v", n, v)
			max = v
			bestN = n
		}
	}
	return bestN
}

func getNOverTotient(n int) float64 {
	return float64(n) / float64(totient(n))
}

func totient(n int) int {
	rp := newRelativelyPrime(n)
	// return len(rp.all())
	return rp.num()
}

type relativelyPrime struct {
	n int

	// deduped and sorted
	factors []int
}

func newRelativelyPrime(n int) relativelyPrime {
	nPrimeFactors := primes.Factors(n)
	nPrimeFactors = utils.Dedupe(nPrimeFactors)
	sort.Ints(nPrimeFactors)

	return relativelyPrime{
		n:       n,
		factors: nPrimeFactors,
	}
}

func (rp relativelyPrime) num() int {
	if rp.n <= 0 {
		return 0
	} else if rp.n == 1 {
		return 1
	}

	output := rp.n
	for _, f := range rp.factors {
		output -= (output / f)
	}

	return output
}

func (rp relativelyPrime) all() []int {
	if rp.n <= 0 {
		return nil
	}

	output := make([]int, 0, rp.n/2)
	// always add one
	output = append(output, 1)
	// only check evens if n is not even
	if !utils.IsEven(rp.n) {
		for e := 2; e < rp.n; e += 2 {
			if rp.is(e) {
				output = append(output, e)
			}
		}
	}
	// always check odds
	for e := 3; e < rp.n; e += 2 {
		if rp.is(e) {
			output = append(output, e)
		}
	}

	return output
}

func (rp relativelyPrime) is(other int) bool {
	for _, f := range rp.factors {
		if other%f == 0 {
			return false
		}
	}

	return true
}

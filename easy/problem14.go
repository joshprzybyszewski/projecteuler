package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/utils"
)

func SolveProblem14() string {
	/*
		The following iterative sequence is defined for the set of positive integers:

		n → n/2 (n is even)
		n → 3n + 1 (n is odd)

		Using the rule above and starting with 13, we generate the following sequence:

		13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

		It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
		Although it has not been proved yet (Collatz Problem), it is thought that all starting
		numbers finish at 1.

		Which starting number, under one million, produces the longest chain?
	*/
	ans, ansLen := longestCollatzChain(1000000)
	return fmt.Sprintf("%v (len of %d)", ans, ansLen)
}

func longestCollatzChain(maxStart int) (int, uint) {
	ccc := newCollatzChainCache(maxStart)

	best, bestLen := 0, uint(0)
	for i := 1; i <= maxStart; i++ {
		myLen := ccc.collatzChainLen(i)
		if myLen > bestLen {
			best = i
			bestLen = myLen
		}
	}
	return best, bestLen
}

type collatzChainCache struct {
	cache []uint
}

func newCollatzChainCache(max int) *collatzChainCache {
	return &collatzChainCache{
		cache: make([]uint, max+1),
	}
}

func (ccc *collatzChainCache) collatzChainLen(
	n int,
) uint {

	if n < len(ccc.cache) && ccc.cache[n] > 0 {
		return ccc.cache[n]
	}

	var res uint

	switch {
	case n <= 1:
		res = 1
	case utils.IsEven(n):
		res = ccc.collatzChainLen(n/2) + 1
	default:
		res = ccc.collatzChainLen((3*n)+1) + 1
	}

	if n < len(ccc.cache) {
		ccc.cache[n] = res
	}

	return res
}

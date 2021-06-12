package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
	"github.com/joshprzybyszewski/projecteuler/primes"
)

func SolveProblem49() string {
	/*
		The arithmetic sequence, 1487, 4817, 8147, in which each of the
		terms increases by 3330, is unusual in two ways: (i) each of the
		three terms are prime, and, (ii) each of the 4-digit numbers are
		permutations of one another.

		There are no arithmetic sequences made up of three 1-, 2-, or 3-digit
		primes, exhibiting this property, but there is one other 4-digit
		increasing sequence.

		What 12-digit number do you form by concatenating the three terms
		in this sequence?
	*/
	ans := getProblem49Sequence()
	return fmt.Sprintf("%v", ans)
}

func getProblem49Sequence() []int {
	for _, start := range primes.Below(10000) {
		if start < 1000 {
			continue
		}
		if start == 1487 {
			// this is the one that we were given.
			continue
		}

		s0 := mathUtils.ToDigits(start)
		for add := 1; add < 10000; add++ {
			next := start + add
			if !primes.Is(next) {
				continue
			}

			s1 := mathUtils.ToDigits(next)
			if !isPermutation(s0, s1) {
				continue
			}

			next2 := next + add
			if !primes.Is(next2) {
				continue
			}

			s2 := mathUtils.ToDigits(next2)
			if !isPermutation(s0, s2) {
				continue
			}

			return []int{start, next, next2}
		}
	}
	return nil
}

func isPermutation(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	seen := make([]bool, len(a))
	for bi := 0; bi < len(b); bi++ {
		found := false
		for ai := 0; ai < len(a); ai++ {
			if seen[ai] {
				continue
			}
			if a[ai] == b[bi] {
				found = true
				seen[ai] = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

package easy

import (
	"fmt"
	"log"
	"sort"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
	"github.com/joshprzybyszewski/projecteuler/primes"
)

func SolveProblem37() string {
	/*
		Find the sum of the only eleven primes that are both
		truncatable from left to right and right to left.
	*/
	tpc := &truncatablePrimesCache{}
	tpc.populate()
	all := tpc.get()
	ans := mathUtils.Sum(all)
	log.Printf("Problem 37: %v = sum(%v)", ans, all)
	return fmt.Sprintf("%v", ans)
}

type truncatablePrimesCache struct {
	leftBySize  map[int][]truncPrime
	rightBySize map[int][]truncPrime

	both map[int]struct{}
}

func (tpc *truncatablePrimesCache) get() []int {
	res := make([]int, 0, 11)
	for n := range tpc.both {
		res = append(res, n)
	}
	sort.Ints(res)
	return res
}

func (tpc *truncatablePrimesCache) populate() {
	tpc.both = map[int]struct{}{}
	numExpected := 11

	singles := primes.Below(10)
	totalRange := primes.Below(1000000)
	splices := []func([]int, int) int{
		func(s []int, p int) int {
			return spliceDigits(s, []int{p})
		}, func(s []int, p int) int {
			return spliceDigits([]int{p}, s)
		},
	}

	for _, smaller := range totalRange {
		sDigits := mathUtils.ToDigits(smaller)
		for _, p := range singles {
			for _, splice := range splices {
				val := splice(sDigits, p)
				if isTruncatablePrime(val) {
					tpc.both[val] = struct{}{}
				}
				if len(tpc.both) >= numExpected {
					return
				}
			}
		}
	}
}

type truncPrime struct {
	digits []int
	n      int
}

func newTruncPrime(n int) truncPrime {
	return truncPrime{
		n:      n,
		digits: mathUtils.ToDigits(n),
	}
}

func (tp truncPrime) isBoth() bool {
	return !tp.isInvalid() && tp.isLeft() && tp.isRight()
}

func (tp truncPrime) isInvalid() bool {
	if tp.n < 10 || !primes.Is(tp.n) {
		return true
	}
	return false
}

func (tp truncPrime) isLeft() bool {
	// returns true if every left-side substring is prime
	for i := 1; i < len(tp.digits); i++ {
		n := spliceDigits(tp.digits[:i], nil)
		if !primes.Is(n) {
			return false
		}
	}

	return true
}

func (tp truncPrime) isRight() bool {
	// returns true if every right-side substring is prime
	for i := len(tp.digits) - 1; i > 0; i-- {
		n := spliceDigits(tp.digits[i:], nil)
		if !primes.Is(n) {
			return false
		}
	}

	return true
}

func isTruncatablePrime(n int) bool {
	return newTruncPrime(n).isBoth()
}

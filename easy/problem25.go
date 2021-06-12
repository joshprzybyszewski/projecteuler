package easy

import (
	"fmt"
	"math/big"

	"github.com/joshprzybyszewski/projecteuler/sequence"
)

func SolveProblem25() string {
	/*
		The Fibonacci sequence is defined by the recurrence relation:

		Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
		Hence the first 12 terms will be:

		F1 = 1
		F2 = 1
		F3 = 2
		F4 = 3
		F5 = 5
		F6 = 8
		F7 = 13
		F8 = 21
		F9 = 34
		F10 = 55
		F11 = 89
		F12 = 144
		The 12th term, F12, is the first term to contain three digits.


		What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
	*/
	ans := getBigFibonacciIndexWithNDigits(1000)
	return fmt.Sprintf("%v", ans)
}

func getBigFibonacciIndexWithNDigits(n int) uint {
	if n < 20 {
		return getFibonacciIndexWithNDigits(n)
	}

	minIndex := uint(4 * n)

	var ans uint
	cb := func(i uint, f *big.Int) bool {
		if i < minIndex {
			return true
		}
		if len(f.String()) >= n {
			ans = i + 1
			return false
		}
		return true
	}
	sequence.BigFibonaccisWhere(cb)

	return ans
}

func getFibonacciIndexWithNDigits(n int) uint {
	if n > 20 {
		// we need the big.Int to do this
		return 0
	}
	var ans uint
	cb := func(i uint, f uint64) bool {
		if len(fmt.Sprintf("%d", f)) >= n {
			ans = i + 1
			return false
		}
		return true
	}
	sequence.FibonaccisWhere(cb)

	return ans
}

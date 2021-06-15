package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/primes"
	"github.com/joshprzybyszewski/projecteuler/utils"
)

func SolveProblem58() string {
	/*
		If one complete new layer is wrapped around the
		spiral above, a square spiral with side length 9
		will be formed. If this process is continued, what
		is the side length of the square spiral for which
		the ratio of primes along both diagonals first falls
		below 10%?
	*/
	ans := getSideLengthOfSpiralWithLessThan10PercentPrimes()
	return fmt.Sprintf("%v", ans)
}

func getSideLengthOfSpiralWithLessThan10PercentPrimes() int {
	numPrimes := 0
	totalCorners := 0

	for sl := 1; ; sl += 2 {
		corners := getSpiralCorners(sl)
		numPrimes += getNumPrimes(corners)
		totalCorners += len(corners)
		isLess := numPrimes*10 < totalCorners
		if sl >= 9 && isLess {
			return sl
		}
		if sl > 30000 || corners[0] > 1000000000000 {
			return -1
		}
	}
}

func getNumPrimes(all []int) int {
	total := 0
	for _, n := range all {
		if primes.Is(n) {
			total++
		}
	}
	return total
}

func getSpiralCorners(n int) []int {
	if utils.IsEven(n) {
		// invalid input!
		return nil
	}
	if n == 1 {
		return []int{1}
	}
	res := make([]int, 0, 4)
	square := n * n
	for c := 0; c < 4; c++ {
		res = append(res, square)
		square -= n - 1
	}
	return res
}

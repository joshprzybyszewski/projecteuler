package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
	"github.com/joshprzybyszewski/projecteuler/primes"
)

func SolveProblem51() string {
	/*
		By replacing the 1st digit of the 2-digit number *3, it turns
		out that six of the nine possible values: 13, 23, 43, 53, 73,
		and 83, are all prime.

		By replacing the 3rd and 4th digits of 56**3 with the same digit,
		this 5-digit number is the first example having seven primes
		among the ten generated numbers, yielding the family: 56003,
		56113, 56333, 56443, 56663, 56773, and 56993. Consequently 56003,
		being the first member of this family, is the smallest prime with
		this property.

		Find the smallest prime which, by replacing part of the number
		(not necessarily adjacent digits) with the same digit, is part
		of an eight prime value family.
	*/
	ans := getSmallestPrimeWithNPrimeValueFamily(8)
	return fmt.Sprintf("%v", ans)
}

func getSmallestPrimeWithNPrimeValueFamily(n int) int {
	// chose 1,000,000 as a shot in the dark max value.
	ps := primes.Below(1000000)
	for _, p := range ps {
		digits, options := getIndexesForPrimeReplacement(p)
		for _, o := range options {
			family := getPrimeReplacements(digits, o)
			if len(family) >= n {
				return p
			}
		}
	}

	return -1
}

// assumes input is prime
func getIndexesForPrimeReplacement(
	p int,
) ([]int, [][]int) {
	res := make([][]int, 3)

	digits := mathUtils.ToDigits(p)
	for i, d := range digits {
		switch d {
		case 0:
			res[0] = append(res[0], i)
		case 1:
			res[1] = append(res[1], i)
		case 2:
			res[2] = append(res[2], i)
		}
	}

	return digits, res
}

func getPrimeReplacements(
	digits,
	indexs []int,
) []int {
	if len(indexs) == 0 {
		return nil
	}

	res := make([]int, 0, 2*len(indexs))
	for r := 0; r <= 9; r++ {
		if r < digits[indexs[0]] {
			continue
		}
		for _, i := range indexs {
			digits[i] = r
		}
		trial := mathUtils.DigitsToInt(digits)
		if primes.Is(trial) {
			res = append(res, trial)
		}
	}
	return res
}

package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/primes"
)

func SolveProblem5() string {
	/*
		2520 is the smallest number
		that can be divided by each of the numbers
		from 1 to 10 without any remainder.

		What is the smallest positive number that is
		evenly divisible by all of the numbers from 1 to 20?
	*/
	p := smallestDivisibleByRange(20)
	return fmt.Sprintf("%d", p)
}

func smallestDivisibleByRange(n int) int {
	allFactors := make([]int, 0, 4)
	countOfAllFactors := make(map[int]int, 4)

	for k := n; k > 1; k-- {
		kFactors := primes.Factors(k)
		kCount := getCountOfNums(kFactors)
		for factor, numNeeded := range kCount {
			for countOfAllFactors[factor] < numNeeded {
				allFactors = append(allFactors, factor)
				countOfAllFactors[factor]++
			}
		}
	}

	res := 1
	for _, f := range allFactors {
		res *= f
	}
	return res
}

func getCountOfNums(s []int) map[int]int {
	m := make(map[int]int, len(s))
	for _, e := range s {
		m[e]++
	}

	return m
}

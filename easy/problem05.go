package easy

import (
	"fmt"
)

func SolveProblem5() {
	/*
		2520 is the smallest number
		that can be divided by each of the numbers
		from 1 to 10 without any remainder.

		What is the smallest positive number that is
		evenly divisible by all of the numbers from 1 to 20?
	*/
	p := smallestDivisibleByRange(20)
	fmt.Printf("Problem 5 Answer: %d\n", p)
}

func smallestDivisibleByRange(n int) int {
	allFactors := make([]int, 0, 4)
	countOfAllFactors := make(map[int]int, 4)

	for k := n; k > 1; k-- {
		kFactors := GetPrimeFactorization(k)
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

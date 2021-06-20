package easy

import (
	"fmt"
	"sort"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
	"github.com/joshprzybyszewski/projecteuler/sequence"
)

func SolveProblem62() string {
	/*
		The cube, 41063625 (345^3), can be permuted to produce two
		other cubes: 56623104 (384^3) and 66430125 (405^3). In fact,
		41063625 is the smallest cube which has exactly three
		permutations of its digits which are also cube.

		Find the smallest cube for which exactly five permutations of its digits are cube.
	*/
	ans := getSmallestPermutationRepresentingNCubes(5)
	return fmt.Sprintf("%v (from %v)", ans[0], ans)
}

func getSmallestPermutationRepresentingNCubes(
	n int,
) []int {

	rangeMin := 1
	rangeMax := 9

	c := uint(1)

	for {
		allInRange := make([]int, 0, 4)
		for cube := sequence.Cubes.GetNth(c); cube < rangeMax; {
			allInRange = append(allInRange, cube)

			c++
			cube = sequence.Cubes.GetNth(c)
		}

		perms := getMaxPermutations(allInRange)
		if len(perms) >= n {
			sort.Ints(perms)
			return perms
		}

		rangeMin *= 10
		rangeMax *= 10
		rangeMax += 9
	}
}

func getMaxPermutations(
	options []int,
) []int {
	var best []int

	for i, val := range options {
		valDigits := mathUtils.ToDigits(val)
		valPerms := []int{val}

		for j := i + 1; j < len(options); j++ {
			jDigits := mathUtils.ToDigits(options[j])
			if isPermutation(valDigits, jDigits) {
				valPerms = append(valPerms, options[j])
			}
		}

		if len(valPerms) > len(best) {
			best = valPerms
		}
	}

	return best
}

package easy

import (
	"fmt"
	"strconv"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
)

func SolveProblem24() string {
	/*
		A permutation is an ordered arrangement of objects.
		For example, 3124 is one possible permutation of the
		digits 1, 2, 3 and 4. If all of the permutations are
		listed numerically or alphabetically, we call it
		lexicographic order. The lexicographic permutations
		of 0, 1 and 2 are:

		012   021   102   120   201   210

		What is the millionth lexicographic permutation
		of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
	*/
	ans := getLexicographicPermutationAtIndex(0, 9, 1000000)
	return fmt.Sprintf("%v", ans)
}

func getLexicographicPermutationAtIndex(
	minDigit, maxDigit int,
	oneIndexed int,
) string {
	if oneIndexed <= 0 {
		return `bad input`
	}

	toPlace := make([]int, 0, maxDigit-minDigit+1)
	for i := minDigit; i <= maxDigit; i++ {
		toPlace = append(toPlace, i)
	}

	str := ``

	d := oneIndexed - 1
	f := mathUtils.Factorial(len(toPlace) - 1)

	for len(toPlace) > 0 {
		i := d / f
		str += strconv.Itoa(toPlace[i])
		toPlace = append(toPlace[:i], toPlace[i+1:]...)
		if len(toPlace) == 0 {
			break
		}
		d = d % f
		f /= len(toPlace)
	}

	return str
}

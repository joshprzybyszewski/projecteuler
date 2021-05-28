package easy

import (
	"testing"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
	"github.com/stretchr/testify/assert"
)

func TestLargestPalindromeInRange(t *testing.T) {
	assert.Equal(t, 9009, largestPalindromeInRange(10, 100))
}

func TestSmallestDivisibleByRange(t *testing.T) {
	assert.Equal(t, 2520, smallestDivisibleByRange(10))
}

func TestSquareOfSum(t *testing.T) {
	assert.Equal(t, 3025, squareOfSum(10))
}

func TestSumOfSquares(t *testing.T) {
	assert.Equal(t, 385, sumOfSquares(10))
}

func TestGetNthPrime(t *testing.T) {
	assert.Equal(t, 13, getNthPrime(6))
}

func TestIsPythagoreanTriplet(t *testing.T) {
	assert.True(t, isPythagoreanTriplet(3, 4, 5))
	assert.True(t, isPythagoreanTriplet(3, 5, 4))
	assert.True(t, isPythagoreanTriplet(4, 3, 5))
	assert.True(t, isPythagoreanTriplet(4, 5, 3))
	assert.True(t, isPythagoreanTriplet(5, 3, 4))
	assert.True(t, isPythagoreanTriplet(5, 4, 3))

	assert.False(t, isPythagoreanTriplet(5, 5, 5))

	assert.True(t, isPythagoreanTriplet(5, 12, 13))
	assert.True(t, isPythagoreanTriplet(8, 15, 17))
}

func TestSumOfPrimesBelow(t *testing.T) {
	assert.Equal(t, 17, sumOfPrimesBelow(10))
}

func TestGetTriangleNumberWithMoreThanNDivisors(t *testing.T) {
	assert.Equal(t, 28, getTriangleNumberWithMoreThanNDivisors(5))
}

func TestLongestCollatzChain(t *testing.T) {
	//  1
	//  2 ->  1
	//  3 -> 10 ->  5 -> 16 ->  8 ->  4 ->  2 ->  1
	//  4 ->  2 ->  1
	//  5 -> 16 ->  8 ->  4 ->  2 ->  1
	//  6 ->  3 -> 10 ->  5 -> 16 ->  8 ->  4 ->  2 ->  1
	//  7 -> 22 -> 11 -> 34 -> 17 -> 52 -> 26 -> 13 -> 40 -> 20 -> 10 ->  5 -> 16 ->  8 ->  4 ->  2 ->  1
	//  8 ->  4 ->  2 ->  1
	//  9 -> 28 -> 14 ->  7 -> 22 -> 11 -> 34 -> 17 -> 52 -> 26 -> 13 -> 40 -> 20 -> 10 ->  5 -> 16 ->  8 ->  4 ->  2 ->  1
	// 10 ->  5 -> 16 ->  8 ->  4 ->  2 ->  1
	// 11 -> 34 -> 17 -> 52 -> 26 -> 13 -> 40 -> 20 -> 10 ->  5 -> 16 ->  8 ->  4 ->  2 ->  1
	// 12 ->  6 ->  3 -> 10 ->  5 -> 16 ->  8 ->  4 ->  2 ->  1
	// 13 -> 40 -> 20 -> 10 ->  5 -> 16 ->  8 ->  4 ->  2 ->  1
	// 14 ->  7 -> 22 -> 11 -> 34 -> 17 -> 52 -> 26 -> 13 -> 40 -> 20 -> 10 ->  5 -> 16 ->  8 ->  4 ->  2 ->  1
	// 15 -> 46 -> 23 -> 70 -> 35 -> 106-> 53 -> 160-> 80 -> 40 -> 20 -> 10 ->  5 -> 16 ->  8 ->  4 ->  2 ->  1
	// 16 ->  8 ->  4 ->  2 ->  1
	best, bestLen := longestCollatzChain(16)
	assert.Equal(t, 9, best)
	assert.Equal(t, 20, bestLen)
}

func TestGetSumOfTwoToThePower(t *testing.T) {
	ans := getSumOfTwoToThePower(15)
	assert.Equal(t, 26, ans)
}

func TestNumberToWord(t *testing.T) {
	assert.Equal(t, `fifteen`, numberToWord(15, true))
	assert.Equal(t, `three hundred and forty-two`, numberToWord(342, true))
	assert.Equal(t, `one hundred and fifteen`, numberToWord(115, true))
	assert.Equal(t, `nine hundred`, numberToWord(900, true))
	assert.Equal(t, `one thousand`, numberToWord(1000, true))
}

func TestNumLettersInWord(t *testing.T) {
	assert.Equal(t, 3, numLettersInWord(1))
	assert.Equal(t, 3, numLettersInWord(2))
	assert.Equal(t, 5, numLettersInWord(3))
	assert.Equal(t, 4, numLettersInWord(4))
	assert.Equal(t, 4, numLettersInWord(5))
}

func TestGetDownRightRoutesInGrid(t *testing.T) {
	assert.Equal(t, 6, getDownRightRoutesInGrid(2))
	assert.Equal(t, 20, getDownRightRoutesInGrid(3))
}

func TestCollapseMaxesDownTriangle(t *testing.T) {
	tri := [][]int{
		{3},
		{7, 4},
		{2, 4, 6},
		{8, 5, 9, 3},
	}
	maxes := collapseMaxesDownTriangle(tri)
	max := mathUtils.MaxInSlice(maxes)
	assert.Equal(t, 23, max)
}

func TestRemove10s(t *testing.T) {
	testCases := []struct {
		input []int
		exp   []int
	}{{
		input: []int{3, 5},
		exp:   []int{3, 5},
	}, {
		input: []int{3, 5, 2, 11},
		exp:   []int{3, 11},
	}, {
		input: []int{3, 5, 2, 11, 5},
		exp:   []int{3, 5, 11},
	}, {
		input: []int{3, 5, 2, 11, 5, 2, 2, 2},
		exp:   []int{2, 2, 3, 11},
	}, {
		input: []int{2, 5},
		exp:   []int{},
	}}

	for _, tc := range testCases {
		act := remove10s(tc.input)
		assert.Equal(t, tc.exp, act)
	}
}

func TestGetPrimeFactorialFactors(t *testing.T) {
	testCases := []struct {
		input int
		exp   []int
	}{{
		input: 2,
		exp:   []int{2},
	}, {
		input: 3,
		exp:   []int{2, 3},
	}, {
		input: 4,
		exp:   []int{2, 3, 2, 2},
	}, {
		input: 5,
		exp:   []int{2, 3, 2, 2, 5},
	}, {
		input: 6,
		exp:   []int{2, 3, 2, 2, 5, 2, 3},
	}, {
		input: 7,
		exp:   []int{2, 3, 2, 2, 5, 2, 3, 7},
	}, {
		input: 8,
		exp:   []int{2, 3, 2, 2, 5, 2, 3, 7, 2, 2, 2},
	}}

	for _, tc := range testCases {
		act := getPrimeFactorialFactors(tc.input)
		assert.Equal(t, tc.exp, act)
	}
}

func TestGetAlphabeticalValue(t *testing.T) {
	assert.Equal(t, 53, getAlphabeticalValue(`COLIN`))
	assert.Equal(t, 53, getAlphabeticalValue(`colin`))
}

func TestGetProblem22Answer(t *testing.T) {
	assert.Equal(t, 53, getProblem22Answer([]string{`COLIN`}))
	assert.Equal(t, 53+(2*53), getProblem22Answer([]string{`COLIN`, `COLIN`}))
}

func TestGetNumberState(t *testing.T) {
	assert.Equal(t, perfect, getNumberState(28))
	assert.Equal(t, abundant, getNumberState(12))
	assert.Equal(t, deficient, getNumberState(11))
}

func TestAllAbundantNumbersBelow(t *testing.T) {
	assert.Equal(t, []int{}, allAbundantNumbersBelow(11))
	assert.Equal(t, []int{}, allAbundantNumbersBelow(12))
	assert.Equal(t, []int{12}, allAbundantNumbersBelow(13))
	assert.Equal(t, []int{12, 18, 20, 24}, allAbundantNumbersBelow(28))
}

func TestHasTwoElementsWithSum(t *testing.T) {
	assert.True(t, hasTwoElementsWithSum(4, []int{2}))
	assert.False(t, hasTwoElementsWithSum(5, []int{2}))
	assert.True(t, hasTwoElementsWithSum(8, []int{2, 3, 4, 5}))
	assert.False(t, hasTwoElementsWithSum(11, []int{2, 3, 4, 5}))
}

func TestGetLexicographicPermutationAtIndex(t *testing.T) {
	assert.Equal(t, `012`, getLexicographicPermutationAtIndex(0, 2, 1))
	assert.Equal(t, `021`, getLexicographicPermutationAtIndex(0, 2, 2))
	assert.Equal(t, `102`, getLexicographicPermutationAtIndex(0, 2, 3))
	assert.Equal(t, `120`, getLexicographicPermutationAtIndex(0, 2, 4))
	assert.Equal(t, `201`, getLexicographicPermutationAtIndex(0, 2, 5))
	assert.Equal(t, `210`, getLexicographicPermutationAtIndex(0, 2, 6))

	assert.Equal(t, `0123`, getLexicographicPermutationAtIndex(0, 3, 1))
	assert.Equal(t, `0132`, getLexicographicPermutationAtIndex(0, 3, 2))
	assert.Equal(t, `0213`, getLexicographicPermutationAtIndex(0, 3, 3))
	assert.Equal(t, `0231`, getLexicographicPermutationAtIndex(0, 3, 4))
	assert.Equal(t, `0312`, getLexicographicPermutationAtIndex(0, 3, 5))
	assert.Equal(t, `0321`, getLexicographicPermutationAtIndex(0, 3, 6))
	assert.Equal(t, `1023`, getLexicographicPermutationAtIndex(0, 3, 7))
	assert.Equal(t, `1032`, getLexicographicPermutationAtIndex(0, 3, 8))
	assert.Equal(t, `1203`, getLexicographicPermutationAtIndex(0, 3, 9))
	assert.Equal(t, `1230`, getLexicographicPermutationAtIndex(0, 3, 10))
	assert.Equal(t, `1302`, getLexicographicPermutationAtIndex(0, 3, 11))
	assert.Equal(t, `1320`, getLexicographicPermutationAtIndex(0, 3, 12))
	assert.Equal(t, `2013`, getLexicographicPermutationAtIndex(0, 3, 13))
	assert.Equal(t, `2031`, getLexicographicPermutationAtIndex(0, 3, 14))
	assert.Equal(t, `2103`, getLexicographicPermutationAtIndex(0, 3, 15))
}

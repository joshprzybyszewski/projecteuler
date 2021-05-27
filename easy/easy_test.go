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

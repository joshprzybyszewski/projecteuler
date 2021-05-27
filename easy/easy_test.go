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

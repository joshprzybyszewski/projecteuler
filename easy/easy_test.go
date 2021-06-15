package easy

import (
	"log"
	"sort"
	"testing"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
	"github.com/joshprzybyszewski/projecteuler/primes"
	"github.com/joshprzybyszewski/projecteuler/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	assert.Equal(t, uint(20), bestLen)
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

func TestGetFib(t *testing.T) {

	assert.Equal(t, uint(1), getFibonacciIndexWithNDigits(1))   //  0 ->  1 = +1
	assert.Equal(t, uint(7), getFibonacciIndexWithNDigits(2))   //  1 ->  7 = +6
	assert.Equal(t, uint(12), getFibonacciIndexWithNDigits(3))  //  7 -> 12 = +5
	assert.Equal(t, uint(17), getFibonacciIndexWithNDigits(4))  // 12 -> 17 = +5
	assert.Equal(t, uint(21), getFibonacciIndexWithNDigits(5))  // 17 -> 21 = +4
	assert.Equal(t, uint(26), getFibonacciIndexWithNDigits(6))  // 21 -> 26 = +5
	assert.Equal(t, uint(31), getFibonacciIndexWithNDigits(7))  // 26 -> 31 = +5
	assert.Equal(t, uint(36), getFibonacciIndexWithNDigits(8))  // 31 -> 36 = +5
	assert.Equal(t, uint(40), getFibonacciIndexWithNDigits(9))  // 36 -> 40 = +4
	assert.Equal(t, uint(45), getFibonacciIndexWithNDigits(10)) // 40 -> 45 = +5
	assert.Equal(t, uint(50), getFibonacciIndexWithNDigits(11)) // 45 -> 50 = +5
	assert.Equal(t, uint(55), getFibonacciIndexWithNDigits(12)) // 50 -> 55 = +5
	assert.Equal(t, uint(60), getFibonacciIndexWithNDigits(13)) // ?? -> ?? = +5
	assert.Equal(t, uint(64), getFibonacciIndexWithNDigits(14)) // ?? -> ?? = +4

	testCases := []struct {
		digits int
		index  uint
	}{
		{digits: 1, index: 1},   // +1
		{digits: 2, index: 7},   // +6
		{digits: 3, index: 12},  //  +5
		{digits: 4, index: 17},  //  +5
		{digits: 5, index: 21},  // +4
		{digits: 6, index: 26},  //  +5
		{digits: 7, index: 31},  //  +5
		{digits: 8, index: 36},  //  +5
		{digits: 9, index: 40},  // +4
		{digits: 10, index: 45}, //  +5
		{digits: 11, index: 50}, //  +5
		{digits: 12, index: 55}, //  +5
		{digits: 13, index: 60}, //  +5
		{digits: 14, index: 64}, // +4
		{digits: 15, index: 69}, //  +5
		{digits: 16, index: 74}, //  +5
		{digits: 17, index: 79}, //  +5
		{digits: 18, index: 84}, //  +5
		{digits: 19, index: 88}, // +4
		{digits: 20, index: 93}, //  +5
		{digits: 21, index: 0},
	}

	for _, tc := range testCases {
		actIndex := getFibonacciIndexWithNDigits(tc.digits)
		assert.Equal(t, tc.index, actIndex, "desired %d digits with index %d, but got %d", tc.digits, tc.index, actIndex)
	}
}

func TestGetBigFibonacciIndexWithNDigits(t *testing.T) {

	testCases := []struct {
		digits int
		index  uint
	}{
		{digits: 1, index: 1},   // +1
		{digits: 2, index: 7},   // +6
		{digits: 3, index: 12},  //  +5
		{digits: 4, index: 17},  //  +5
		{digits: 5, index: 21},  // +4
		{digits: 6, index: 26},  //  +5
		{digits: 7, index: 31},  //  +5
		{digits: 8, index: 36},  //  +5
		{digits: 9, index: 40},  // +4
		{digits: 10, index: 45}, //  +5
		{digits: 11, index: 50}, //  +5
		{digits: 12, index: 55}, //  +5
		{digits: 13, index: 60}, //  +5
		{digits: 14, index: 64}, // +4
		{digits: 15, index: 69}, //  +5
		{digits: 16, index: 74}, //  +5
		{digits: 17, index: 79}, //  +5
		{digits: 18, index: 84}, //  +5
		{digits: 19, index: 88}, // +4
		{digits: 20, index: 93}, //  +5
		{digits: 21, index: 98},
		{digits: 22, index: 103},
		{digits: 23, index: 107},
	}

	for _, tc := range testCases {
		actIndex := getBigFibonacciIndexWithNDigits(tc.digits)
		assert.Equal(t, tc.index, actIndex, "desired %d digits with index %d, but got %d", tc.digits, tc.index, actIndex)
	}
}

func TestGetDecimalMantissa(t *testing.T) {
	testCases := []struct {
		divisor   int
		expPrefix []int
		expCycle  []int
	}{{
		divisor:   2,
		expPrefix: []int{5},
	}, {
		divisor:  3,
		expCycle: []int{3},
	}, {
		divisor:   4,
		expPrefix: []int{2, 5},
	}, {
		divisor:   5,
		expPrefix: []int{2},
	}, {
		divisor:   6,
		expPrefix: []int{1},
		expCycle:  []int{6},
	}, {
		divisor:  7,
		expCycle: []int{1, 4, 2, 8, 5, 7},
	}, {
		divisor:   8,
		expPrefix: []int{1, 2, 5},
	}, {
		divisor:  9,
		expCycle: []int{1},
	}, {
		divisor:   10,
		expPrefix: []int{1},
	}}

	for _, tc := range testCases {
		prefix, cycle := getDecimalMantissa(tc.divisor)
		if len(tc.expPrefix) == 0 {
			assert.Empty(t, prefix)
		} else {
			assert.Equal(t, tc.expPrefix, prefix)
		}
		if len(tc.expCycle) == 0 {
			assert.Empty(t, cycle)
		} else {
			assert.Equal(t, tc.expCycle, cycle)
		}
	}
}

func TestGetMaxCycleLengthOfDivisors(t *testing.T) {
	assert.Equal(t, 3, getMaxCycleLengthOfDivisors(5))
	assert.Equal(t, 7, getMaxCycleLengthOfDivisors(10))
}

func TestGetNumConsecutivePrimes(t *testing.T) {
	assert.Equal(t, 40, getNumConsecutivePrimes(1, 41))
	assert.Equal(t, 80, getNumConsecutivePrimes(-79, 1601))
}

func TestGetSumOfProblem28Spiral(t *testing.T) {
	assert.Equal(t, 1, getSumOfProblem28Spiral(1))
	assert.Equal(t, 25, getSumOfProblem28Spiral(3))
	assert.Equal(t, 101, getSumOfProblem28Spiral(5))
	assert.Zero(t, getSumOfProblem28Spiral(6))
}

func TestGetNumDistinctTermsForARaisedB(t *testing.T) {
	brute := bruteForceSequenceARaisedB(2, 5, 2, 5)
	// 4, 8, 9, 16, 25, 27, 32, 64, 81, 125, 243, 256, 625, 1024, 3125
	assert.Equal(t, []uint64{4, 8, 9, 16, 25, 27, 32, 64, 81, 125, 243, 256, 625, 1024, 3125}, brute)
	assert.Equal(t, 15, getNumDistinctTermsForARaisedB(2, 5, 2, 5))

	for max := uint64(6); max <= 16; max++ {
		brute = bruteForceSequenceARaisedB(2, max, 2, max)
		assert.Equal(t, len(brute), getNumDistinctTermsForARaisedB(2, int(max), 2, int(max)), `checked for %d`, int(max))
	}

	log.Printf("starting 2,8,2,8")
	brute = bruteForceSequenceARaisedB(2, 8, 2, 8)
	assert.Equal(t, len(brute), getNumDistinctTermsForARaisedB(2, 8, 2, 8))

}

func bruteForceSequenceARaisedB(
	minA, maxA,
	minB, maxB uint64,
) []uint64 {
	res := make([]uint64, 0, 8)
	for a := minA; a <= maxA; a++ {
		for b := minB; b <= maxB; b++ {
			aToTheB := mathUtils.Raised(a, b)
			if utils.Uint64Contains(res, aToTheB) {
				continue
			}
			res = append(res, aToTheB)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return res
}

func TestWaysToBuildSum(t *testing.T) {
	assert.Equal(t, 0, waysToBuildSum(0, []uint{1}))
	assert.Equal(t, 1, waysToBuildSum(1, []uint{1}))
	assert.Equal(t, 1, waysToBuildSum(2, []uint{1}))

	assert.Equal(t, 2, waysToBuildSum(2, []uint{2, 1})) // {2}, {1, 1}
	assert.Equal(t, 2, waysToBuildSum(3, []uint{2, 1})) // {2, 1}, {1, 1, 1}
	assert.Equal(t, 3, waysToBuildSum(4, []uint{2, 1})) // {2, 2}, {2, 1, 1}, {1, 1, 1, 1}
}

func TestIsTrickyFraction(t *testing.T) {
	assert.True(t, isTrickyFraction(49, 98))

	assert.False(t, isTrickyFraction(33, 66))
	assert.False(t, isTrickyFraction(30, 50))
}

func TestGetAllTrickyFractions(t *testing.T) {
	ns, ds := getAllTrickyFractions()
	assert.Len(t, ns, 4)
	assert.Len(t, ds, 4)
}

func TestIsSumOfFactorialsOfDigits(t *testing.T) {
	assert.True(t, isSumOfFactorialsOfDigits(145))
	assert.False(t, isSumOfFactorialsOfDigits(31))
}

func TestIsCircularPrime(t *testing.T) {
	assert.True(t, isCircularPrime(2))
	assert.True(t, isCircularPrime(13))
	assert.True(t, isCircularPrime(31))
	assert.True(t, isCircularPrime(197))
}

func TestSpliceDigits(t *testing.T) {
	assert.Zero(t, spliceDigits([]int{}, []int{}))

	assert.Equal(t, 13, spliceDigits([]int{1, 3}, []int{}))
	assert.Equal(t, 13, spliceDigits([]int{}, []int{1, 3}))

	assert.Equal(t, 12, spliceDigits([]int{1}, []int{2}))
	assert.Equal(t, 1225, spliceDigits([]int{1, 2}, []int{2, 5}))
}

func TestGetNumberOfCircularPrimesBelow(t *testing.T) {
	assert.Equal(t, 13, getNumberOfCircularPrimesBelow(100))
}

func TestIsPalindrome(t *testing.T) {
	assert.True(t, isPalindrome(585))
	assert.True(t, isBase2Palindrome(585))
}

func TestIsTruncatablePrime(t *testing.T) {
	assert.False(t, isTruncatablePrime(2))
	assert.False(t, isTruncatablePrime(3))
	assert.False(t, isTruncatablePrime(5))
	assert.False(t, isTruncatablePrime(7))

	assert.False(t, isTruncatablePrime(311))

	assert.True(t, primes.Is(7))
	assert.True(t, primes.Is(37))
	assert.True(t, primes.Is(379))
	assert.True(t, isTruncatablePrime(3797))
	assert.True(t, primes.Is(797))
	assert.True(t, primes.Is(97))
	assert.True(t, primes.Is(7))

	assert.False(t, isTruncatablePrime(3113))

	assert.True(t, isTruncatablePrime(23))
	assert.True(t, isTruncatablePrime(37))
	assert.True(t, isTruncatablePrime(53))
	assert.True(t, isTruncatablePrime(73))
	assert.True(t, isTruncatablePrime(313))
	assert.True(t, isTruncatablePrime(317))
	assert.True(t, isTruncatablePrime(373))
	assert.True(t, isTruncatablePrime(797))
	assert.True(t, isTruncatablePrime(3137))
	assert.True(t, isTruncatablePrime(3797))
	assert.True(t, isTruncatablePrime(739397))
}

func TestGetNTruncatablePrimes(t *testing.T) {
	truncatable := []int{
		23, 37, 53, 73,
		313, 317,
		373, 797,
		3137, 3797,
		739397,
	}
	tpc := &truncatablePrimesCache{}
	tpc.populate()
	assert.Equal(t, truncatable, tpc.get())
}

func TestIsRightAngleTriangle(t *testing.T) {
	assert.True(t, isRightAngleTriangle(20, 48, 52))
	assert.True(t, isRightAngleTriangle(24, 45, 51))
	assert.True(t, isRightAngleTriangle(30, 40, 50))
}

func TestGetNumberRightAngleTrianglesWithPerimeter(t *testing.T) {
	assert.Equal(t, 3, getNumberRightAngleTrianglesWithPerimeter(120))
}

func TestGetNumberRepresentedAtChampernowneDigit(t *testing.T) {
	rep := 1

	n := 1
	expIndex := 0
	expIndexMax := func(rep int) int {
		switch {
		case rep > 100000:
			// unsupported
			return -1
		case rep >= 10000:
			return 5
		case rep >= 1000:
			return 4
		case rep >= 100:
			return 3
		case rep >= 10:
			return 2
		case rep > 0:
			return 1
		default:
			return 0
		}
	}

	for rep < 100000 {
		actRep, index := getNumberRepresentedAtChampernowneDigit(n)
		require.Equal(t,
			rep,
			actRep,
			`unexpected result for getNumberRepresentedAtChampernowneDigit(%d): represents %d`,
			n, rep,
		)
		require.Equal(t,
			expIndex,
			index,
			`unexpected result for getNumberRepresentedAtChampernowneDigit(%d): represents %d`,
			n, rep,
		)

		require.Equal(t,
			mathUtils.ToDigits(rep)[expIndex],
			getChampernowneDigit(n),
			`unexpected result for getChampernowneDigit(%d): represents %d`,
			n, rep,
		)

		n++
		expIndex++
		if expIndex == expIndexMax(rep) {
			expIndex = 0
			rep++
		}
	}
}

func TestIsProblem43Property(t *testing.T) {
	assert.False(t, isProblem43Property(nil))

	assert.True(t, isProblem43Property([]int{1, 4, 0, 6, 3, 5, 7, 2, 8, 9}))
}

func TestIsGoldbachOtherConjectureCounterExample(t *testing.T) {
	assert.False(t, isGoldbachOtherConjectureCounterExample(9))
	assert.False(t, isGoldbachOtherConjectureCounterExample(15))
	assert.False(t, isGoldbachOtherConjectureCounterExample(21))
	assert.False(t, isGoldbachOtherConjectureCounterExample(25))
	assert.False(t, isGoldbachOtherConjectureCounterExample(27))
	assert.False(t, isGoldbachOtherConjectureCounterExample(33))

	// not a composite (it's a prime)
	assert.False(t, isGoldbachOtherConjectureCounterExample(11))

}

func TestGetConsecutiveNumbersWithNDistinctPrimeFactors(t *testing.T) {
	assert.Equal(t, []int{14, 15}, getConsecutiveNumbersWithNDistinctPrimeFactors(2))
	assert.Equal(t, []int{644, 645, 646}, getConsecutiveNumbersWithNDistinctPrimeFactors(3))
}

func TestGetNumDistinctPrimeFactors(t *testing.T) {
	assert.Equal(t, 2, getNumDistinctPrimeFactors(14))
	assert.Equal(t, 2, getNumDistinctPrimeFactors(15))
	assert.Equal(t, 3, getNumDistinctPrimeFactors(644))
	assert.Equal(t, 3, getNumDistinctPrimeFactors(645))
	assert.Equal(t, 3, getNumDistinctPrimeFactors(646))
}

func TestIsPermutation(t *testing.T) {
	assert.False(t, isPermutation([]int{1, 2, 3}, nil))
	assert.False(t, isPermutation([]int{1, 2, 3}, []int{2, 2, 1}))
	assert.False(t, isPermutation([]int{1, 2, 3}, []int{3, 0, 1}))

	assert.True(t, isPermutation([]int{1, 2, 3}, []int{2, 3, 1}))
	assert.True(t, isPermutation([]int{1, 2, 3}, []int{1, 2, 3}))
}

func TestGetPrimeWithMostConsecutiveAdditivesBelow(t *testing.T) {
	assert.Equal(t, 41, getPrimeWithMostConsecutiveAdditivesBelow(100))
	assert.Equal(t, 953, getPrimeWithMostConsecutiveAdditivesBelow(1000))
	assert.NotEqual(t, 978269, getPrimeWithMostConsecutiveAdditivesBelow(1000000))
}

func TestGetNToTheNSeriesAnswer(t *testing.T) {
	assert.Equal(t, 10405071317, getNToTheNSeriesAnswer(10))
}

func TestGetPrimeReplacements(t *testing.T) {
	act := getPrimeReplacements(mathUtils.ToDigits(13), []int{0})
	exp := []int{13, 23, 43, 53, 73, 83}
	assert.Equal(t, exp, act)

	act = getPrimeReplacements(mathUtils.ToDigits(56003), []int{2, 3})
	exp = []int{56003, 56113, 56333, 56443, 56663, 56773, 56993}
	assert.Equal(t, exp, act)

	act = getPrimeReplacements(mathUtils.ToDigits(17), nil)
	assert.Empty(t, act)
}

func TestGetIndexesForPrimeReplacement(t *testing.T) {
	digits, options := getIndexesForPrimeReplacement(13)
	expOptions := [][]int{
		nil,
		{0},
		nil,
	}
	assert.Equal(t, []int{1, 3}, digits)
	assert.Equal(t, expOptions, options)

	digits, options = getIndexesForPrimeReplacement(56003)
	expOptions = [][]int{
		{2, 3},
		nil,
		nil,
	}
	assert.Equal(t, []int{5, 6, 0, 0, 3}, digits)
	assert.Equal(t, expOptions, options)

	digits, options = getIndexesForPrimeReplacement(56113)
	expOptions = [][]int{
		nil,
		{2, 3},
		nil,
	}
	assert.Equal(t, []int{5, 6, 1, 1, 3}, digits)
	assert.Equal(t, expOptions, options)
}

func TestGetSmallestPrimeWithNPrimeValueFamily(t *testing.T) {
	act := getSmallestPrimeWithNPrimeValueFamily(6)
	assert.Equal(t, 13, act)

	act = getSmallestPrimeWithNPrimeValueFamily(7)
	assert.Equal(t, 56003, act)
}

func TestIsLychrelNumber(t *testing.T) {
	assert.False(t, isLychrelNumber(47))
	assert.False(t, isLychrelNumber(349))

	assert.True(t, isLychrelNumber(196))
}

func TestReverseInts(t *testing.T) {
	assert.Nil(t, reverseInts(nil))
	assert.Empty(t, reverseInts([]int{}))
	assert.Equal(t, []int{1}, reverseInts([]int{1}))
	assert.Equal(t, []int{2, 1}, reverseInts([]int{1, 2}))
	assert.Equal(t, []int{3, 2, 1}, reverseInts([]int{1, 2, 3}))
}

package primes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getAllCachers() []cacher {
	return []cacher{
		newSliceCache(),
	}
}

func TestCacherIs(t *testing.T) {
	cachers := getAllCachers()

	for _, c := range cachers {
		require.True(t, c.is(17))
	}
}

func TestCacherBelow(t *testing.T) {
	defer func(prevSlice int) {
		maxSliceCacheSize = prevSlice
	}(maxSliceCacheSize)
	maxSliceCacheSize = 3

	cachers := getAllCachers()

	for _, c := range cachers {
		act := c.below(4)
		assert.Equal(t, []int{2, 3}, act)

		act = c.below(5)
		assert.Equal(t, []int{2, 3}, act)

		act = c.below(6)
		assert.Equal(t, []int{2, 3, 5}, act)

		act = c.below(7)
		assert.Equal(t, []int{2, 3, 5}, act)

		act = c.below(8)
		assert.Equal(t, []int{2, 3, 5, 7}, act)
	}
}

func TestLoadingFromFile(t *testing.T) {
	sc := newSliceCacheFromFile([]string{`2`, `3`, `5`})
	assert.True(t, sc.is(7))
	act := sc.below(15)
	assert.Equal(t, []int{2, 3, 5, 7, 11, 13}, act)
	actOutput := sc.knownToString()
	assert.Equal(t, []string{`2`, `3`, `5`, `7`, `11`, `13`}, actOutput)

	sc2 := newSliceCacheFromFile(actOutput)
	assert.True(t, sc2.is(7))
	act = sc2.below(30)
	assert.Equal(t, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}, act)
	actOutput = sc2.knownToString()
	assert.Equal(t, []string{`2`, `3`, `5`, `7`, `11`, `13`, `17`, `19`, `23`, `29`}, actOutput)
}

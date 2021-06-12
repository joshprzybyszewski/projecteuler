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

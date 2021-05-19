package primes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactors(t *testing.T) {
	assert.Equal(t, []int{2, 5}, Factors(10))
	assert.Equal(t, []int{2, 3, 3}, Factors(18))
	assert.Equal(t, []int{3, 17}, Factors(51))
	assert.Equal(t, []int{2, 2, 2, 3, 5}, Factors(120))
}

func TestIs(t *testing.T) {
	assert.False(t, Is(-1))
	assert.False(t, Is(0))
	assert.False(t, Is(1))
	assert.True(t, Is(2))
	assert.True(t, Is(3))
	assert.False(t, Is(4))
	assert.True(t, Is(5))
	assert.False(t, Is(6))
	assert.True(t, Is(7))
	assert.False(t, Is(8))
	assert.False(t, Is(9))
	assert.False(t, Is(10))
	assert.True(t, Is(11))
	assert.False(t, Is(12))

	// just some large boys from wikipedia
	assert.True(t, Is(7919))
	assert.True(t, Is(999331))
}

func TestBelow(t *testing.T) {
	assert.Empty(t, Below(-1))
	assert.Empty(t, Below(0))
	assert.Empty(t, Below(1))
	assert.Empty(t, Below(2))
	assert.Equal(t, []int{2}, Below(3))
	assert.Equal(t, []int{2, 3}, Below(4))
	assert.Equal(t, []int{2, 3}, Below(5))
	assert.Equal(t, []int{2, 3, 5}, Below(6))

	assert.Equal(t,
		[]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29},
		Below(30),
	)
}

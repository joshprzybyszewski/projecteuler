package mathUtils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNPandigital(t *testing.T) {

	assert.True(t, IsNPandigital([]int{1}, 1))
	assert.True(t, IsNPandigital([]int{1, 2}, 2))
	assert.True(t, IsNPandigital([]int{1, 3, 2}, 3))
	assert.True(t, IsNPandigital([]int{1, 4, 3, 2}, 4))

	assert.False(t, IsNPandigital([]int{1, 4, 3, 2}, 5))
	assert.False(t, IsNPandigital([]int{1, 4, 3, 2}, 9))
	assert.False(t, IsNPandigital([]int{1, 4, 3, 2}, 10))
	assert.False(t, IsNPandigital([]int{1, 4, 3, 2}, 0))
	assert.False(t, IsNPandigital([]int{1, 4, 3, 2}, -1))
}

func TestIsPandigital(t *testing.T) {

	assert.False(t, IsPandigital([]int{1}))
	assert.False(t, IsPandigital([]int{1, 2}))
	assert.False(t, IsPandigital([]int{1, 3, 2}))
	assert.False(t, IsPandigital([]int{1, 4, 3, 2}))

	assert.True(t, IsPandigital([]int{1, 4, 3, 2, 9, 6, 8, 5, 7}))
}

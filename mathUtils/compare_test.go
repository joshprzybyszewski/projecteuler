package mathUtils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	assert.Equal(t, 0, Max(-1, 0))
	assert.Equal(t, 2, Max(1, 2))
	assert.Equal(t, 23, Max(23, 2))
}

func TestMaxInSlice(t *testing.T) {
	assert.Equal(t, 0, MaxInSlice(nil))
	assert.Equal(t, 0, MaxInSlice([]int{0, 0}))
	assert.Equal(t, 23, MaxInSlice([]int{1, 5, 23}))
}

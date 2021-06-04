package mathUtils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSqrt(t *testing.T) {
	s, ok := IntSqrt(16)
	assert.True(t, ok)
	assert.Equal(t, 4, s)

	s, ok = IntSqrt(15)
	assert.False(t, ok)
	assert.Equal(t, 0, s)
}

func TestRaised(t *testing.T) {
	assert.Equal(t, uint64(1), Raised(2, 0))
	assert.Equal(t, uint64(2), Raised(2, 1))
	assert.Equal(t, uint64(4), Raised(2, 2))
	assert.Equal(t, uint64(8), Raised(2, 3))
	assert.Equal(t, uint64(16), Raised(2, 4))

	assert.Equal(t, uint64(1), Raised(3, 0))
	assert.Equal(t, uint64(3), Raised(3, 1))
	assert.Equal(t, uint64(9), Raised(3, 2))
	assert.Equal(t, uint64(27), Raised(3, 3))
	assert.Equal(t, uint64(81), Raised(3, 4))
}

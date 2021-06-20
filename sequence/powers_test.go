package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCubes(t *testing.T) {
	assert.True(t, Cubes.Is(1))
	assert.True(t, Cubes.Is(8))
	assert.True(t, Cubes.Is(27))

	assert.False(t, Cubes.Is(24))
	assert.False(t, Cubes.Is(34))
	assert.False(t, Cubes.Is(44))
	assert.False(t, Cubes.Is(54))
}

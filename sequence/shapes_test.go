package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriangular(t *testing.T) {
	assert.True(t, Triangular.Is(1))
	assert.True(t, Triangular.Is(3))
	assert.True(t, Triangular.Is(6))
	assert.True(t, Triangular.Is(10))
	assert.True(t, Triangular.Is(15))
	assert.True(t, Triangular.Is(21))
	assert.True(t, Triangular.Is(28))
	assert.True(t, Triangular.Is(36))
	assert.True(t, Triangular.Is(45))
	assert.True(t, Triangular.Is(55))

	assert.False(t, Triangular.Is(24))
	assert.False(t, Triangular.Is(34))
	assert.False(t, Triangular.Is(44))
	assert.False(t, Triangular.Is(54))

	assert.True(t, Triangular.Is(40755))
}

func TestPentagonalIs(t *testing.T) {
	assert.True(t, Pentagonal.Is(1))
	assert.True(t, Pentagonal.Is(5))
	assert.True(t, Pentagonal.Is(12))
	assert.True(t, Pentagonal.Is(22))
	assert.True(t, Pentagonal.Is(35))
	assert.True(t, Pentagonal.Is(51))
	assert.True(t, Pentagonal.Is(70))
	assert.True(t, Pentagonal.Is(92))
	assert.True(t, Pentagonal.Is(117))
	assert.True(t, Pentagonal.Is(145))

	assert.False(t, Pentagonal.Is(24))
	assert.False(t, Pentagonal.Is(34))
	assert.False(t, Pentagonal.Is(44))
	assert.False(t, Pentagonal.Is(54))

	assert.True(t, Pentagonal.Is(40755))
}

func TestPentagonalGetNth(t *testing.T) {
	assert.Equal(t, 1, Pentagonal.GetNth(1))
	assert.Equal(t, 5, Pentagonal.GetNth(2))
	assert.Equal(t, 12, Pentagonal.GetNth(3))
	assert.Equal(t, 22, Pentagonal.GetNth(4))
	assert.Equal(t, 35, Pentagonal.GetNth(5))
	assert.Equal(t, 51, Pentagonal.GetNth(6))
	assert.Equal(t, 70, Pentagonal.GetNth(7))
	assert.Equal(t, 92, Pentagonal.GetNth(8))
	assert.Equal(t, 117, Pentagonal.GetNth(9))
	assert.Equal(t, 145, Pentagonal.GetNth(10))
}

func TestHexagonalGetNth(t *testing.T) {
	assert.Equal(t, 1, Hexagonal.GetNth(1))
	assert.Equal(t, 6, Hexagonal.GetNth(2))
	assert.Equal(t, 15, Hexagonal.GetNth(3))
	assert.Equal(t, 28, Hexagonal.GetNth(4))
	assert.Equal(t, 45, Hexagonal.GetNth(5))

	assert.True(t, Hexagonal.Is(40755))
}

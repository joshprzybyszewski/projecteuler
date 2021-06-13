package mathUtils

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChoose(t *testing.T) {
	assert.Equal(t, 1, Choose(1, 1))

	assert.Equal(t, 1, Choose(2, 0))
	assert.Equal(t, 2, Choose(2, 1))
	assert.Equal(t, 1, Choose(2, 2))

	assert.Equal(t, 1, Choose(3, 0))
	assert.Equal(t, 3, Choose(3, 1))
	assert.Equal(t, 3, Choose(3, 2))
	assert.Equal(t, 1, Choose(3, 3))

	assert.Equal(t, 1, Choose(4, 0))
	assert.Equal(t, 4, Choose(4, 1))
	assert.Equal(t, 6, Choose(4, 2))
	assert.Equal(t, 4, Choose(4, 3))
	assert.Equal(t, 1, Choose(4, 4))
}

func TestBigChoose(t *testing.T) {
	assert.Equal(t, big.NewInt(1), BigChoose(big.NewInt(1), big.NewInt(1)))

	assert.Equal(t, big.NewInt(1), BigChoose(big.NewInt(2), big.NewInt(0)))
	assert.Equal(t, big.NewInt(2), BigChoose(big.NewInt(2), big.NewInt(1)))
	assert.Equal(t, big.NewInt(1), BigChoose(big.NewInt(2), big.NewInt(2)))

	assert.Equal(t, big.NewInt(1), BigChoose(big.NewInt(3), big.NewInt(0)))
	assert.Equal(t, big.NewInt(3), BigChoose(big.NewInt(3), big.NewInt(1)))
	assert.Equal(t, big.NewInt(3), BigChoose(big.NewInt(3), big.NewInt(2)))
	assert.Equal(t, big.NewInt(1), BigChoose(big.NewInt(3), big.NewInt(3)))

	assert.Equal(t, big.NewInt(1), BigChoose(big.NewInt(4), big.NewInt(0)))
	assert.Equal(t, big.NewInt(4), BigChoose(big.NewInt(4), big.NewInt(1)))
	assert.Equal(t, big.NewInt(6), BigChoose(big.NewInt(4), big.NewInt(2)))
	assert.Equal(t, big.NewInt(4), BigChoose(big.NewInt(4), big.NewInt(3)))
	assert.Equal(t, big.NewInt(1), BigChoose(big.NewInt(4), big.NewInt(4)))

	assert.Equal(t, `242519269720337121015504`, BigChoose(big.NewInt(100), big.NewInt(25)).String())
}

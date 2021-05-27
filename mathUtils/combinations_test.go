package mathUtils

import (
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

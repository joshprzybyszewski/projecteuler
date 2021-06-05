package mathUtils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToDigits(t *testing.T) {
	assert.Empty(t, ToDigits(-123))

	assert.Equal(t, []int{0}, ToDigits(0))
	assert.Equal(t, []int{1}, ToDigits(1))
	assert.Equal(t, []int{1, 2}, ToDigits(12))
	assert.Equal(t, []int{1, 2, 3}, ToDigits(123))
}

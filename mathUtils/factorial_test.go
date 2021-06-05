package mathUtils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	assert.Equal(t, 0, Factorial(-1))
	assert.Equal(t, 1, Factorial(0))
	assert.Equal(t, 1, Factorial(1))
	assert.Equal(t, 2, Factorial(2))
	assert.Equal(t, 6, Factorial(3))
	assert.Equal(t, 24, Factorial(4))
	assert.Equal(t, 3628800, Factorial(10))
}

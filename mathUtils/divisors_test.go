package mathUtils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivisors(t *testing.T) {
	assert.ElementsMatch(t, []int{1, 2, 3, 6}, Divisors(6))
	assert.ElementsMatch(t, []int{1, 2, 3, 4, 6, 12}, Divisors(12))
	assert.ElementsMatch(t, []int{1, 7}, Divisors(7))
}

func TestProperDivisors(t *testing.T) {
	assert.ElementsMatch(t, []int{1, 2, 3}, ProperDivisors(6))
	assert.ElementsMatch(t, []int{1, 2, 3, 4, 6}, ProperDivisors(12))
	assert.ElementsMatch(t, []int{1}, ProperDivisors(7))
}

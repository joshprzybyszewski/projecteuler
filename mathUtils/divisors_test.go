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

func TestGetDivisorsFromPrimeFactors(t *testing.T) {
	assert.Equal(t, Divisors(6), GetDivisorsFromPrimeFactors(6, []int{2, 3}))
	assert.Equal(t, Divisors(12), GetDivisorsFromPrimeFactors(12, []int{2, 2, 3}))
	assert.Equal(t, Divisors(7), GetDivisorsFromPrimeFactors(7, []int{7}))

	assert.Equal(t, Divisors(24), GetDivisorsFromPrimeFactors(24, []int{2, 2, 2, 3}))
}

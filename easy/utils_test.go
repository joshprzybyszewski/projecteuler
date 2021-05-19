package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPrimeFactorization(t *testing.T) {
	assert.Equal(t, []int{2, 5}, GetPrimeFactorization(10))
	assert.Equal(t, []int{2, 3, 3}, GetPrimeFactorization(18))
	assert.Equal(t, []int{3, 17}, GetPrimeFactorization(51))
	assert.Equal(t, []int{2, 2, 2, 3, 5}, GetPrimeFactorization(120))
}

func TestIsPrime(t *testing.T) {
	assert.False(t, IsPrime(-1))
	assert.False(t, IsPrime(0))
	assert.False(t, IsPrime(1))
	assert.True(t, IsPrime(2))
	assert.True(t, IsPrime(3))
	assert.False(t, IsPrime(4))
	assert.True(t, IsPrime(5))
	assert.False(t, IsPrime(6))
	assert.True(t, IsPrime(7))
	assert.True(t, IsPrime(7))
}

package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfEvenFibonaccisBelow(t *testing.T) {
	// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
	exp := 2 + 8 + 34
	assert.Equal(t, exp, sumOfEvenFibonaccisBelow(100))
}

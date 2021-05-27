package mathUtils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	testCases := []struct {
		input []int
		exp   int
	}{{
		input: []int{2, 2},
		exp:   4,
	}, {
		input: []int{1, 1, 1, 1},
		exp:   1,
	}, {
		input: []int{},
		exp:   0,
	}}

	for _, tc := range testCases {
		act := Product(tc.input)
		assert.Equal(t, tc.exp, act)
	}
}

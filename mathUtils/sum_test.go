package mathUtils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	testCases := []struct {
		input []int
		exp   int
	}{{
		input: []int{2, 2},
		exp:   4,
	}, {
		input: []int{1, 1, 1, 1},
		exp:   4,
	}, {
		input: []int{},
		exp:   0,
	}}

	for _, tc := range testCases {
		act := Sum(tc.input)
		assert.Equal(t, tc.exp, act)
	}
}

func TestSumOfDigits(t *testing.T) {
	testCases := []struct {
		input int
		exp   int
	}{{
		input: 4,
		exp:   4,
	}, {
		input: 1111,
		exp:   4,
	}, {
		input: 0,
		exp:   0,
	}, {
		input: 102,
		exp:   3,
	}}

	for _, tc := range testCases {
		act := SumOfDigits(tc.input)
		assert.Equal(t, tc.exp, act)
	}
}

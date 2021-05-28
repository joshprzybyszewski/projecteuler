package mathUtils

import (
	"math"
	"math/big"
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

func TestBigProduct(t *testing.T) {
	testCases := []struct {
		input []int
		exp   *big.Int
	}{{
		input: []int{2, 2},
		exp:   big.NewInt(4),
	}, {
		input: []int{1, 1, 1, 1},
		exp:   big.NewInt(1),
	}, {
		input: []int{},
		exp:   big.NewInt(0),
	}, {
		input: []int{math.MaxInt32, math.MaxInt32, math.MaxInt32},
		exp: big.NewInt(1).Mul(
			big.NewInt(math.MaxInt32),
			big.NewInt(1).Mul(
				big.NewInt(math.MaxInt32),
				big.NewInt(math.MaxInt32),
			),
		),
	}}

	for _, tc := range testCases {
		act := BigProduct(tc.input)
		assert.Equal(t, tc.exp, act)
	}
}

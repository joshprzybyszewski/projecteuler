package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelativelyPrimeIs(t *testing.T) {
	testCases := []struct {
		input int
		other int
		exp   bool
	}{{
		input: 2,
		other: 1,
		exp:   true,
	}, {
		input: 3,
		other: 2,
		exp:   true,
	}, {
		input: 4,
		other: 3,
		exp:   true,
	}, {
		input: 4,
		other: 2,
		exp:   false,
	}, {
		input: 9,
		other: 2,
		exp:   true,
	}, {
		input: 9,
		other: 6,
		exp:   false,
	}}

	for _, tc := range testCases {
		assert.Equal(t, tc.exp, newRelativelyPrime(tc.input).is(tc.other))
	}
}

func TestRelativelyPrimeAll(t *testing.T) {
	testCases := []struct {
		input int
		exp   []int
	}{{
		input: 2,
		exp:   []int{1},
	}, {
		input: 3,
		exp:   []int{1, 2},
	}, {
		input: 4,
		exp:   []int{1, 3},
	}}

	for _, tc := range testCases {
		assert.Equal(t, tc.exp, newRelativelyPrime(tc.input).all())
	}
}

func TestTotient(t *testing.T) {
	testCases := []struct {
		input int
		exp   int
	}{{
		input: 2,
		exp:   1,
	}, {
		input: 3,
		exp:   2,
	}, {
		input: 4,
		exp:   2,
	}, {
		input: 5,
		exp:   4,
	}, {
		input: 6,
		exp:   2,
	}, {
		input: 7,
		exp:   6,
	}, {
		input: 8,
		exp:   4,
	}, {
		input: 9,
		exp:   6,
	}, {
		input: 10,
		exp:   4,
	}}

	for _, tc := range testCases {
		assert.Equal(t, tc.exp, totient(tc.input))
	}
}

func TestGetMaxNOverTotientWithMaxN(t *testing.T) {
	assert.Equal(t, 6, getMaxNOverTotientWithMaxN(10))
}

func TestRelativelyPrimeNum(t *testing.T) {
	for n := 1; n <= 1000; n++ {
		rp := newRelativelyPrime(n)
		all := rp.all()
		assert.Equal(t, rp.num(), len(all), `num did not match for %v`, all)
	}
}

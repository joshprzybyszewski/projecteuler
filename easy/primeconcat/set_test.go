package primeconcat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetAdd(t *testing.T) {
	testCases := []struct {
		s   *set
		w   int
		exp *set
	}{{
		s:   newSet([]int{3, 5}),
		w:   7,
		exp: nil,
	}, {
		s:   newSet([]int{3, 7}),
		w:   109,
		exp: newSet([]int{3, 7, 109}),
	}}

	for _, tc := range testCases {
		act := tc.s.add(tc.w)
		assert.Equal(t, tc.exp, act)
	}
}

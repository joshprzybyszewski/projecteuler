package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSingleEntry(t *testing.T) {
	assert.False(t, IsSingleEntry(nil))
	assert.False(t, IsSingleEntry([]int{}))
	assert.False(t, IsSingleEntry([]int{1, 2}))

	assert.True(t, IsSingleEntry([]int{1}))
	assert.True(t, IsSingleEntry([]int{1, 1}))
}

func TestUint64Contains(t *testing.T) {
	assert.False(t, Uint64Contains(nil, 1))
	assert.False(t, Uint64Contains([]uint64{}, 1))
	assert.False(t, Uint64Contains([]uint64{1, 2}, 3))
	assert.True(t, Uint64Contains([]uint64{1, 2}, 1))
}

func TestNumSubEntries(t *testing.T) {
	assert.Zero(t, NumSubEntries(nil, nil))
	assert.Zero(t, NumSubEntries([]int{}, nil))
	assert.Zero(t, NumSubEntries(nil, []int{}))
	assert.Zero(t, NumSubEntries([]int{}, []int{}))
	assert.Zero(t, NumSubEntries([]int{1}, []int{}))
	assert.Zero(t, NumSubEntries([]int{1}, []int{2}))
	assert.Zero(t, NumSubEntries([]int{2, 2, 3, 2, 3, 3}, []int{2, 3}))
	assert.Zero(t, NumSubEntries([]int{2, 2, 3, 3, 3}, []int{2, 3}))
	assert.Zero(t, NumSubEntries([]int{2, 2, 3, 3}, []int{2, 2, 3}))

	assert.Equal(t, 1, NumSubEntries([]int{2}, []int{2}))
	assert.Equal(t, 2, NumSubEntries([]int{2, 2}, []int{2}))
	assert.Equal(t, 2, NumSubEntries([]int{2, 2, 3, 3}, []int{2, 3}))
	assert.Equal(t, 2, NumSubEntries([]int{2, 2, 2, 2}, []int{2, 2}))
	assert.Equal(t, 3, NumSubEntries([]int{2, 2, 2, 3, 3, 3}, []int{2, 3}))
	assert.Equal(t, 3, NumSubEntries([]int{2, 2, 2, 2, 2, 2}, []int{2, 2}))

	assert.Zero(t, NumSubEntries([]int{2, 2, 5, 5}, []int{2, 5, 5}))
}

func TestIsSubset(t *testing.T) {
	assert.False(t, IsSubset(nil, nil))
	assert.False(t, IsSubset([]int{}, nil))
	assert.False(t, IsSubset(nil, []int{}))
	assert.False(t, IsSubset([]int{}, []int{}))
	assert.False(t, IsSubset([]int{1}, []int{2}))
	assert.False(t, IsSubset([]int{2, 3, 3}, []int{3, 2}))
	assert.False(t, IsSubset([]int{2}, []int{2, 2}))

	assert.True(t, IsSubset([]int{1}, []int{}))
	assert.True(t, IsSubset([]int{2, 2, 3, 2, 3, 3}, []int{2, 3}))
	assert.True(t, IsSubset([]int{2, 2, 3, 3, 3}, []int{2, 3}))
	assert.True(t, IsSubset([]int{2, 2, 3, 3}, []int{2, 2, 3}))

	assert.True(t, IsSubset([]int{2}, []int{2}))
	assert.True(t, IsSubset([]int{2, 2}, []int{2}))
	assert.True(t, IsSubset([]int{2, 2, 3, 3}, []int{2, 3}))
	assert.True(t, IsSubset([]int{2, 2, 2, 2}, []int{2, 2}))
	assert.True(t, IsSubset([]int{2, 2, 2, 3, 3, 3}, []int{2, 3}))
	assert.True(t, IsSubset([]int{2, 2, 2, 2, 2, 2}, []int{2, 2}))
	assert.True(t, IsSubset([]int{2, 2, 5, 5}, []int{2, 5, 5}))
}

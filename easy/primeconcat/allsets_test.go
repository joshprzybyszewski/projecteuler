package primeconcat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllSetsAdd(t *testing.T) {
	as := NewAllSets()

	assert.Equal(t, []int{3, 7}, as.GetMinSet(2))
	assert.Equal(t, []int{3, 37, 67}, as.GetMinSet(3))
	assert.Equal(t, []int{3, 7, 109, 673}, as.GetMinSet(4))
	// assert.Equal(t, []int{13, 5197, 5701, 6733, 8389}, as.GetMinSet(5))
}

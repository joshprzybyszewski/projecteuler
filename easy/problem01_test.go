package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOf3sAnd5sBelow(t *testing.T) {
	assert.Equal(t, 23, sumOf3sAnd5sBelow(10))

	// 3 + 5 + 6 + 9 + 10 + 12 + 15 + 18 = 78
	assert.Equal(t, 78, sumOf3sAnd5sBelow(20))
}

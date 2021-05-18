package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOf3sAnd5sBelow(t *testing.T) {
	assert.Equal(t, 23, sumOf3sAnd5sBelow(10))
}

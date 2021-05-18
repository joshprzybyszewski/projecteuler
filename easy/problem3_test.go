package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestPrimeFactor(t *testing.T) {
	assert.Equal(t, 29, largestPrimeFactor(13195))
}

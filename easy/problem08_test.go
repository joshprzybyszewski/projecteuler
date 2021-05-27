package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindGreatestRollingProduct(t *testing.T) {
	assert.Equal(t, 20, findGreatestRollingProduct([]int{2, 3, 4, 5}, 2))
	assert.Equal(t, 80, findGreatestRollingProduct([]int{2, 3, 4, 5, 4, 3, 2}, 3))

	populateProblem8()
	assert.Equal(t, 5832, findGreatestRollingProduct(problem8Digits, 4))
}

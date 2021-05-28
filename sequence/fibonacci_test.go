package sequence

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacisBelow(t *testing.T) {
	act := FibonaccisBelow(10)
	exp := []int{
		1, 1,
		2, 3,
		5, 8,
	}
	assert.Equal(t, exp, act)

	act = FibonaccisBelow(5)
	exp = []int{
		1, 1,
		2, 3,
	}
	assert.Equal(t, exp, act)

	act = FibonaccisBelow(100)
	exp = []int{
		1, 1,
		2, 3,
		5, 8,
		13,
		21,
		34,
		55,
		89,
	}
	assert.Equal(t, exp, act)
}

func TestBigFibonacisWhere(t *testing.T) {
	exp := []string{
		`1`, `1`, `2`, `3`, `5`, `8`,
		`13`, `21`, `34`, `55`, `89`,
		`144`, `233`, `377`, `610`, `987`,
		`1597`, `2584`, `4181`, `6765`,
		`10946`, `17711`, `28657`, `46368`, `75025`,
		`121393`, `196418`, `317811`, `514229`, `832040`,
		`1346269`, `2178309`, `3524578`, `5702887`, `9227465`,
		`14930352`,
	}
	BigFibonaccisWhere(func(i uint, f *big.Int) bool {
		if int(i) >= len(exp) {
			return false
		}
		assert.Equal(t, exp[i], f.String())
		return true
	})
}

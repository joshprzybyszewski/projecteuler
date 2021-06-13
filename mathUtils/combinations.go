package mathUtils

import (
	"math/big"
)

// n choose k = ( n! ) / ( k! * (n-k)! )
func Choose(n, k int) int {
	if n < 1 || k < 0 || k > n {
		return 0
	}

	choose := 1
	for m := n; m > n-k; m-- {
		choose = choose * m
	}
	for d := k; d > 1; d-- {
		choose = choose / d
	}

	return choose
}

func BigChoose(n, k *big.Int) *big.Int {
	if n.Sign() < 1 || k.Sign() < 0 || k.Cmp(n) > 0 {
		return big.NewInt(0)
	}

	nLessK := big.NewInt(0)
	nLessK.Sub(n, k)

	choose := big.NewInt(1)
	for m := big.NewInt(n.Int64()); m.Cmp(nLessK) > 0; {
		choose.Mul(choose, m)
		m.Sub(m, big.NewInt(1))
	}
	for d := big.NewInt(k.Int64()); d.Sign() > 0; {
		choose.Div(choose, d)
		d.Sub(d, big.NewInt(1))
	}

	return choose
}

package mathUtils

import "math/big"

func Product(elems []int) int {
	if len(elems) == 0 {
		return 0
	}

	p := 1
	for _, e := range elems {
		if e == 0 {
			return 0
		}
		p *= e
	}
	return p
}

func BigProduct(elems []int) *big.Int {
	if len(elems) == 0 {
		return big.NewInt(0)
	}

	p := big.NewInt(1)
	for _, e := range elems {
		if e == 0 {
			return big.NewInt(0)
		}
		p = p.Mul(p, big.NewInt(int64(e)))
	}
	return p
}

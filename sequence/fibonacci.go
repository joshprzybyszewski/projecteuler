package sequence

import "math/big"

func FibonaccisBelow(
	max int,
) []int {
	a, b := 0, 1
	res := make([]int, 0, max/4)
	res = append(res, b)
	for {
		a, b = b, a+b
		if b >= max {
			break
		}
		res = append(res, b)
	}
	return res
}

func FibonaccisWhere(
	callback func(i uint, f uint64) bool,
) {
	if callback == nil {
		return
	}

	if !callback(0, 1) {
		return
	}

	i, a, b := uint(1), uint64(0), uint64(1)
	for {
		a, b = b, a+b
		if !callback(i, b) {
			break
		}
		i++
	}
}

func BigFibonaccisWhere(
	callback func(i uint, f *big.Int) bool,
) {
	if callback == nil {
		return
	}

	if !callback(0, big.NewInt(1)) {
		return
	}

	i := uint(1)
	a := big.NewInt(0)
	b := big.NewInt(1)
	for {

		sum := big.NewInt(0)
		newB := sum.Add(a, b)
		if !callback(i, newB) {
			break
		}
		a = b
		b = newB
		i++
	}
}

package mathUtils

import "math"

func IntSqrt(n int) (int, bool) {
	s := int(math.Floor(math.Sqrt(float64(n))))
	if s*s == n {
		return s, true
	}

	return 0, false
}

func Raised(a, b uint64) uint64 {
	res := uint64(1)

	for p := uint64(0); p < b; p++ {
		res *= a
	}

	return res
}

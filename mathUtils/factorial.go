package mathUtils

func Factorial(n int) int {
	if n < 0 {
		return 0
	}

	perm := 1
	for m := n; m > 0; m-- {
		perm = perm * m
	}

	return perm
}

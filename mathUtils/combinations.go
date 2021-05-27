package mathUtils

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

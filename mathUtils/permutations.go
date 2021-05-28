package mathUtils

// nPr = ( n! ) / ( (n-r)! )
func Permutations(n, r int) int {
	if n < 0 || r < 0 || r > n {
		return 0
	}

	perm := 1
	for m := n; m > n-r; m-- {
		perm = perm * m
	}

	return perm
}

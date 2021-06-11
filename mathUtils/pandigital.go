package mathUtils

import "sort"

func IsPandigital(all []int) bool {
	return IsNPandigital(all, 9)
}

func IsNPandigital(all []int, n int) bool {
	if n > 10 {
		// unsupported
		return false
	}
	if len(all) != n {
		return false
	}
	sort.Ints(all)

	diff := 1
	if n == 10 {
		diff = 0
	}

	for i := range all {
		if all[i] != i+diff {
			return false
		}
	}

	return true
}

package utils

func Uint64Contains(s []uint64, t uint64) bool {
	for _, e := range s {
		if t == e {
			return true
		}
	}
	return false
}

func IntContains(s []int, t int) bool {
	for _, e := range s {
		if t == e {
			return true
		}
	}
	return false
}

func IsSingleEntry(s []int) bool {
	if len(s) == 0 {
		return false
	}

	return len(s) == getNumSameElems(s)
}

// assumes the inputs are sorted
func IsSubset(
	superset, subset []int,
) bool {
	if len(superset) == 0 || len(subset) > len(superset) {
		return false
	}

	superIndex := -1
	for _, s := range subset {
		superIndex++
		for superIndex < len(superset) && superset[superIndex] != s {
			superIndex++
		}
		if superIndex >= len(superset) {
			return false
		}
	}

	return true
}

func NumSubEntries(
	superset, subset []int,
) int {
	if len(superset) == 0 || len(subset) == 0 {
		return 0
	}

	runSub := getNumSameElems(subset)
	runSuper := getNumSameElems(superset)
	if runSuper < runSub || runSuper%runSub != 0 {
		return 0
	}

	mult := runSuper / runSub

	if len(subset)*mult != len(superset) {
		return 0
	}

	for i, e := range superset {
		if e != subset[i/mult] {
			return 0
		}
	}

	return mult
}

func getNumSameElems(
	s []int,
) int {
	if len(s) == 0 {
		return 0
	}

	f := s[0]
	for i, e := range s {
		if e != f {
			return i
		}
	}
	return len(s)
}

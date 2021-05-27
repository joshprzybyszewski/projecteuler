package mathUtils

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxInSlice(s []int) int {
	if len(s) == 0 {
		return 0
	}
	max := s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}

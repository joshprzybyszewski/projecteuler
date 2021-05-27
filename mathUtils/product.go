package mathUtils

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

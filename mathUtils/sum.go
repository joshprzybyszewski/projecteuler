package mathUtils

func Sum(elems []int) int {
	total := 0
	for _, e := range elems {
		total += e
	}
	return total
}

func SumOfDigits(n int) int {
	return Sum(ToDigits(n))
}

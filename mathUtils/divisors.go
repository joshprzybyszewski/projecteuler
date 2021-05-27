package mathUtils

import "math"

// Divisors returns a slice of numbers that evenly divide into n
// Both 1 and n are included.
func Divisors(n int) []int {
	divisors := make([]int, 0, 4)
	max := int(math.Floor(math.Sqrt(float64(n))))
	for i := 1; i <= max; i++ {
		if n%i != 0 {
			continue
		}

		divisors = append(divisors, i)

		d := n / i
		if d != i {
			divisors = append(divisors, d)
		}
	}
	return divisors
}

// ProperDivisors are numbers less than n which divide evenly into n
func ProperDivisors(n int) []int {
	all := Divisors(n)
	proper := make([]int, 0, len(all))

	for _, d := range all {
		if d < n {
			proper = append(proper, d)
		}
	}

	return proper
}

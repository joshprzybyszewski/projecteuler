package easy

import (
	"fmt"
	"math"
)

func SolveProblem3() {
	/*
		The prime factors of 13195 are 5, 7, 13 and 29.

		What is the largest prime factor of the number 600851475143 ?
	*/
	sum := largestPrimeFactor(600851475143)
	fmt.Printf("Problem 3 Answer: %d\n", sum)
}

func largestPrimeFactor(n int) int {
	max := int(math.Floor(math.Sqrt(float64(n))))
	for f := max; f > 1; f-- {
		if n%f != 0 {
			// not a factor
			continue
		}
		isPrime := true
		maxD := f / 2
		for d := 2; d < maxD; d++ {
			if f%d == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			return f
		}
	}
	return 1
}

package easy

import "fmt"

func SolveProblem2() {
	sum := sumOfEvenFibonaccisBelow(4000000)
	fmt.Printf("Problem 2 Answer: %d\n", sum)
}

func sumOfEvenFibonaccisBelow(max int) int {
	a, b := 0, 1
	total := 0
	for {
		a, b = b, a+b
		if b >= max {
			return total
		}
		if b%2 == 0 {
			total += b
		}
	}
}

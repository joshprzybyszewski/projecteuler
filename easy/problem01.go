package easy

import "fmt"

func SolveProblem1() string {
	/*
		If we list all the natural numbers below 10 that are multiples
		of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

		Find the sum of all the multiples of 3 or 5 below 1000.
	*/
	sum := sumOf3sAnd5sBelow(1000)
	return fmt.Sprintf("%d", sum)
}

func sumOf3sAnd5sBelow(max int) int {
	total := 0
	for i := 1; i < max; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	return total
}

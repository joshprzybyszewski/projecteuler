package easy

import (
	"fmt"
)

func SolveProblem15() {
	/*
		Starting in the top left corner of a 2×2 grid, and only
		being able to move to the right and down, there are
		exactly 6 routes to the bottom right corner.

		How many such routes are there through a 20×20 grid?
	*/
	ans := getDownRightRoutesInGrid(20)
	fmt.Printf("Problem 15 Answer: %v\n", ans)
}

func getDownRightRoutesInGrid(n int) int {
	total := 0
	for k := 0; k <= n; k++ {
		choose := nChooseK(n, k)
		total += (choose * choose)
	}
	return total
}

func nChooseK(n, k int) int {
	if n < 1 || k < 0 || k > n {
		return 0
	}

	choose := 1
	for m := n; m > n-k; m-- {
		choose = choose * m
	}
	for d := k; d > 1; d-- {
		choose = choose / d
	}

	return choose
}

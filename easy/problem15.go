package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
)

func SolveProblem15() string {
	/*
		Starting in the top left corner of a 2×2 grid, and only
		being able to move to the right and down, there are
		exactly 6 routes to the bottom right corner.

		How many such routes are there through a 20×20 grid?
	*/
	ans := getDownRightRoutesInGrid(20)
	return fmt.Sprintf("%v", ans)
}

func getDownRightRoutesInGrid(n int) int {
	total := 0
	for k := 0; k <= n; k++ {
		choose := mathUtils.Choose(n, k)
		total += (choose * choose)
	}
	return total
}

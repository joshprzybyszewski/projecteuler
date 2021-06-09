package easy

import (
	"fmt"
)

func SolveProblem39() {
	/*
		If p is the perimeter of a right angle triangle with integral
		length sides, {a,b,c}, there are exactly three solutions for p = 120.

		{20,48,52}, {24,45,51}, {30,40,50}

		For which value of p ≤ 1000, is the number of solutions maximised?
	*/
	ans := getPerimeterWithMaxNumberTrianglesBelow(1001)
	fmt.Printf("Problem 39 Answer: %v\n", ans)
}

func getPerimeterWithMaxNumberTrianglesBelow(maxPerimeter int) int {
	bestPerim := -1
	bestPerimNum := -1

	for p := 1; p < maxPerimeter; p++ {
		n := getNumberRightAngleTrianglesWithPerimeter(p)
		if n > bestPerimNum {
			bestPerim = p
			bestPerimNum = n
		}
	}
	return bestPerim
}

func getNumberRightAngleTrianglesWithPerimeter(p int) int {
	numTriangles := 0
	for a := 1; a < p-2; a++ {
		pLessA := p - a
		for b := a; b < pLessA-1; b++ {
			c := pLessA - b
			if isRightAngleTriangle(a, b, c) {
				numTriangles++
			}
		}
	}
	return numTriangles
}

func isRightAngleTriangle(a, b, c int) bool {
	return c*c == (a*a)+(b*b)
}

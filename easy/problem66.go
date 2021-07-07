package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/sequence"
)

func SolveProblem66() string {
	/*
		Consider quadratic Diophantine equations of the form:

		x^2 – Dy^2 = 1

		For example, when D=13, the minimal solution in x is 649^2 – 13×180^2 = 1.

		It can be assumed that there are no solutions in positive integers when D is square.

		By finding minimal solutions in x for D = {2, 3, 5, 6, 7}, we obtain the following:

		3^2 – 2×2^2 = 1
		2^2 – 3×1^2 = 1
		9^2 – 5×4^2 = 1
		5^2 – 6×2^2 = 1
		8^2 – 7×3^2 = 1

		Hence, by considering minimal solutions in x for D ≤ 7, the largest x is obtained when D=5.
		Find the value of D ≤ 1000 in minimal solutions of x for which the largest value of x is obtained.
	*/
	ans := getMaxXForMinDiophantinesAtOrBelow(1000)
	return fmt.Sprintf("%v", ans)
}

func getMaxXForMinDiophantinesAtOrBelow(max int) int {
	var maxX uint
	bestD := -1
	for d := 2; d <= max; d++ {
		if sequence.Squares.Is(d) {
			continue
		}
		x, _ := getMinDiophantine(d)
		// fmt.Printf("getMinDiophantine(%d) = %d, %d\n", d, x, y)
		if x > maxX {
			maxX = x
			bestD = d
			fmt.Printf("\tupdating maxX, bestD = %d, %d\n", maxX, bestD)
		}
	}
	return bestD
}

func getMinDiophantine(d int) (x, y uint) {
	var ok bool
	for y = 1; ; y++ {
		if y > 10000000 {
			fmt.Printf("d is huge. %d\n", d)
			sqx := 1 + (d * sequence.Squares.GetNth(y))
			return uint(sqx), y
		}
		sqx := 1 + (d * sequence.Squares.GetNth(y))
		if x, ok = sequence.Squares.GetPosition(sqx); ok {
			return x, y
		}
	}
}

func getMinDiophantineOriginal(d int) (x, y uint) {
	for x = 2; ; x++ {
		sqx := sequence.Squares.GetNth(x)
		for y = 1; y < x; y++ {

			sqy := sequence.Squares.GetNth(y)
			v := sqx - (d * sqy)
			if v == 1 {
				return x, y
			} else if v < 1 {
				break
			}
		}
	}
}

package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/sequence"
)

func SolveProblem45() string {
	/*
		Triangle, pentagonal, and hexagonal numbers are generated by the following formulae:

		Triangle	 	Tn=n(n+1)/2	 	1, 3, 6, 10, 15, ...
		Pentagonal	 	Pn=n(3n−1)/2	 	1, 5, 12, 22, 35, ...
		Hexagonal	 	Hn=n(2n−1)	 	1, 6, 15, 28, 45, ...

		It can be verified that T285 = P165 = H143 = 40755.

		Find the next triangle number that is also pentagonal and hexagonal.
	*/
	ans := getProblem45Answer()
	return fmt.Sprintf("%v", ans)
}

func getProblem45Answer() int {
	for h := uint(144); ; h++ {
		n := sequence.Hexagonal.GetNth(h)
		if sequence.Pentagonal.Is(n) && sequence.Triangular.Is(n) {
			return n
		}
	}
}

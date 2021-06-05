package easy

import (
	"fmt"
)

func SolveProblem31() {
	/*
		In the United Kingdom the currency is made up
		of pound (£) and pence (p). There are eight
		coins in general circulation:

		1p, 2p, 5p, 10p, 20p, 50p, £1 (100p), and £2 (200p).
		It is possible to make £2 in the following way:

		1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p

		How many different ways can £2 be made using any number of coins?
	*/
	coins := []uint{
		1, 2, 5, 10, 20, 50, 100, 200,
	}
	ans := waysToBuildSum(200, coins)
	fmt.Printf("Problem 31 Answer: %v\n", ans)
}

func waysToBuildSum(
	target uint,
	options []uint,
) int {
	if target == 0 || len(options) == 0 {
		return 0
	}

	total := 0
	c := options[0]
	for numC := uint(0); numC*c <= target; numC++ {
		switch {
		case numC*c == target:
			total++
		default:
			total += waysToBuildSum(target-numC*c, options[1:])
		}
	}

	return total
}

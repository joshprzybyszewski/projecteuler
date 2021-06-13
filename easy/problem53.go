package easy

import (
	"fmt"
	"log"
	"math/big"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
)

func SolveProblem53() string {
	/*
		How many, not necessarily distinct, values of "n choose r" for 1 <= n <= 100 are greater than 1 million?
	*/
	ans := choosesMoreThan(1, 100, 1000000)
	return fmt.Sprintf("%v", ans)
}

func choosesMoreThan(
	minN, maxN int64,
	v int64,
) int {
	bigV := big.NewInt(v)
	sum := 0
	for n := minN; n <= maxN; n++ {
		for r := int64(0); r <= n; r++ {
			c := mathUtils.BigChoose(
				big.NewInt(n),
				big.NewInt(r),
			)
			if c.Cmp(bigV) > 0 {
				log.Printf("%d choose %d = %d\n", n, r, c)
				sum++
			}
		}
	}
	return sum
}

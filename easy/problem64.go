package easy

import (
	"fmt"
	"log"

	"github.com/joshprzybyszewski/projecteuler/sequence"
	"github.com/joshprzybyszewski/projecteuler/utils"
)

const (
	maxCycleDetectLength = 10000
)

func SolveProblem64() string {
	/*
		How many continued fractions for n <= 10,000 have an odd period?
	*/
	ans := getNumSqrtContinuedFractionsWithOddPeriodsAtOrBelow(10000)
	return fmt.Sprintf("%v", ans)
}

func getNumSqrtContinuedFractionsWithOddPeriodsAtOrBelow(n int) int {
	total := 0

	for i := 1; i <= n; i++ {
		prefix, cycle := sqrtAsContinuedFraction(i)
		log.Printf("sqrt(%d) = [%v;(%v)], period = %d", i, prefix, cycle, len(cycle))
		if !utils.IsEven(len(cycle)) {
			total++
		}
	}
	return total
}

func sqrtAsContinuedFraction(r int) (prefix, repeat []int) {
	if sequence.Squares.Is(r) {
		return nil, nil
	}

	a0 := 1
	for sequence.Squares.GetNth(uint(a0+1)) < r {
		a0++
	}

	/*
		       1		            1
		----------------   ----------------   ----------------
			   1         =    sqrt(r) + a0  =  int(a01/(r-a0*a0)) +
		  ------------        ------------
		  sqrt(r) - a0        r - a0*a0
	*/

	// return detectCycle(sqrtContinuedFractionCycleOutputer{
	// 	of: r,
	// })
	return nil, nil
}

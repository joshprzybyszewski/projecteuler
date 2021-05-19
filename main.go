package main

import (
	"fmt"
	"time"

	"github.com/joshprzybyszewski/projecteuler/easy"
)

var (
	highestSolved = 100
)

func main() {
	solveAll()
}

func solveAll() {
	for i := 1; i <= highestSolved; i++ {
		solve(i)
	}
}

func solve(puzzleNum int) {
	solved := true
	defer func(t0 time.Time) {
		if solved {
			fmt.Printf("\tTook: %s\n", time.Since(t0))
		}
	}(time.Now())
	switch puzzleNum {
	case 1:
		easy.SolveProblem1()
	case 2:
		easy.SolveProblem2()
	case 3:
		// takes a long time to generate the prime cache
		// easy.SolveProblem3()
		solved = false
	case 4:
		easy.SolveProblem4()
	case 5:
		easy.SolveProblem5()
	case 6:
		easy.SolveProblem6()
	case 7:
		// takes a long time
		// easy.SolveProblem7()
		solved = false
	case 8:
		easy.SolveProblem8()
	case 9:
		easy.SolveProblem9()
	case 10:
		// takes a long time to find the first 2 million primes
		// easy.SolveProblem10()
		solved = false
	case 11:
		easy.SolveProblem11()
	default:
		solved = false
		highestSolved = puzzleNum - 1
	}
}

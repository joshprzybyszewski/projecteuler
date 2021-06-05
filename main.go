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
		easy.SolveProblem7()
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
	case 12:
		easy.SolveProblem12()
	case 13:
		easy.SolveProblem13()
	case 14:
		easy.SolveProblem14()
	case 15:
		easy.SolveProblem15()
	case 16:
		easy.SolveProblem16()
	case 17:
		easy.SolveProblem17()
	case 18:
		easy.SolveProblem18()
	case 19:
		easy.SolveProblem19()
	case 20:
		easy.SolveProblem20()
	case 21:
		easy.SolveProblem21()
	case 22:
		easy.SolveProblem22()
	case 23:
		// takes a long time
		// easy.SolveProblem23()
		solved = false
	case 24:
		easy.SolveProblem24()
	case 25:
		easy.SolveProblem25()
	case 26:
		easy.SolveProblem26()
	case 27:
		easy.SolveProblem27()
	case 28:
		easy.SolveProblem28()
	case 29:
		easy.SolveProblem29()
	case 30:
		easy.SolveProblem30()
	case 31:
		easy.SolveProblem31()
	case 32:
		easy.SolveProblem32()
	// TEMPLATE_MARKER //
	case 67:
		easy.SolveProblem67()
	default:
		solved = false
		highestSolved = puzzleNum - 1
	}
}

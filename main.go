package main

import "github.com/joshprzybyszewski/projecteuler/easy"

var (
	highestSolved = 3
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
	switch puzzleNum {
	case 1:
		easy.SolveProblem1()
	case 2:
		easy.SolveProblem2()
	case 3:
		easy.SolveProblem3()
	}
}

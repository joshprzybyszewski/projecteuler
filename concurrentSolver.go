package main

import "runtime"

type concurrentSolver struct {
	pns     chan int
	results chan solution
}

func newConcurrentSolver() *concurrentSolver {
	return &concurrentSolver{
		pns:     make(chan int, highestSolved),
		results: make(chan solution, highestSolved),
	}
}

func (cs *concurrentSolver) sendWork(highestPuzzleNum int) {
	defer close(cs.pns)

	for i := 1; i <= highestPuzzleNum; i++ {
		cs.pns <- i
	}
}

func (cs *concurrentSolver) getResults() []solution {
	defer close(cs.results)

	solutions := make([]solution, 0, highestSolved)
	for s := range cs.results {
		solutions = append(solutions, s)
		if len(solutions) == highestSolved {
			break
		}
	}
	return solutions
}

func (cs *concurrentSolver) startWorkers() {
	for i := 0; i < runtime.NumCPU(); i++ {
		go cs.work()
	}
}

func (cs *concurrentSolver) work() {
	for pn := range cs.pns {
		cs.results <- solve(pn)
	}
}

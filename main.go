package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/joshprzybyszewski/projecteuler/easy"
	"github.com/joshprzybyszewski/projecteuler/primes"
)

var (
	highestSolved = 100
)

func main() {
	primes.InitCache()
	defer primes.SaveCache()

	solveAll()
}

func solveAll() {
	numSolved := 0
	defer func(t0 time.Time) {
		fmt.Printf("\nSolved %d puzzles in %s\n", numSolved, time.Since(t0))
	}(time.Now())

	pns := make(chan int, highestSolved)
	results := make(chan solution, highestSolved)

	for i := 0; i < runtime.NumCPU(); i++ {
		go work(pns, results)
	}

	for i := 1; i <= highestSolved; i++ {
		pns <- i
	}
	close(pns)

	solutions := make([]solution, 0, highestSolved)
	for s := range results {
		solutions = append(solutions, s)
		if len(solutions) == highestSolved {
			break
		}
	}
	sort.Slice(solutions, func(i, j int) bool {
		return solutions[i].num < solutions[j].num
	})

	_, output := buildOutput(solutions)
	writeOutputOfSolutions(output)

	sort.Slice(solutions, func(i, j int) bool {
		return solutions[i].duration < solutions[j].duration
	})
	numSolved, output = buildOutput(solutions)
	fmt.Print(output)
}

func buildOutput(
	solutions []solution,
) (int, string) {
	numSolved := 0
	var sb strings.Builder
	sb.WriteString("|Problem #|Answer|Duration|\n")
	sb.WriteString("|-:|-|-:|\n")
	for _, s := range solutions {
		if !s.solved {
			sb.WriteString(fmt.Sprintf("|%d|Skipped|N/A|\n", s.num))
			continue
		}
		numSolved++
		sb.WriteString(fmt.Sprintf("|%d|%s|%s|\n", s.num, s.answer, s.duration))
	}
	return numSolved, sb.String()
}

func writeOutputOfSolutions(fileString string) {
	answerFile, err := os.Create(`answers.md`)
	if err != nil {
		log.Fatal(err)
	}
	defer answerFile.Close()
	w := bufio.NewWriter(answerFile)
	_, err = w.WriteString(fileString)
	if err != nil {
		log.Fatal(err)
	}

	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

func work(
	c <-chan int,
	results chan solution,
) {
	for pn := range c {
		results <- solve(pn)
	}
}

type solution struct {
	answer   string
	solved   bool
	num      int
	duration time.Duration
}

func solve(puzzleNum int) solution {
	s := solution{
		num:    puzzleNum,
		solved: true,
	}
	go func() {
		<-time.After(5 * time.Second)
		if s.answer == `` && s.solved {
			log.Printf("Problem %d is taking a while", s.num)
		}
	}()

	t0 := time.Now()

	switch puzzleNum {
	case 1:
		s.answer = easy.SolveProblem1()
	case 2:
		s.answer = easy.SolveProblem2()
	case 3:
		s.answer = easy.SolveProblem3()
	case 4:
		s.answer = easy.SolveProblem4()
	case 5:
		s.answer = easy.SolveProblem5()
	case 6:
		s.answer = easy.SolveProblem6()
	case 7:
		s.answer = easy.SolveProblem7()
	case 8:
		s.answer = easy.SolveProblem8()
	case 9:
		s.answer = easy.SolveProblem9()
	case 10:
		s.answer = easy.SolveProblem10()
	case 11:
		s.answer = easy.SolveProblem11()
	case 12:
		s.answer = easy.SolveProblem12()
	case 13:
		s.answer = easy.SolveProblem13()
	case 14:
		s.answer = easy.SolveProblem14()
	case 15:
		s.answer = easy.SolveProblem15()
	case 16:
		s.answer = easy.SolveProblem16()
	case 17:
		s.answer = easy.SolveProblem17()
	case 18:
		s.answer = easy.SolveProblem18()
	case 19:
		s.answer = easy.SolveProblem19()
	case 20:
		s.answer = easy.SolveProblem20()
	case 21:
		s.answer = easy.SolveProblem21()
	case 22:
		s.answer = easy.SolveProblem22()
	case 23:
		s.answer = easy.SolveProblem23()
	case 24:
		s.answer = easy.SolveProblem24()
	case 25:
		s.answer = easy.SolveProblem25()
	case 26:
		s.answer = easy.SolveProblem26()
	case 27:
		s.answer = easy.SolveProblem27()
	case 28:
		s.answer = easy.SolveProblem28()
	case 29:
		s.answer = easy.SolveProblem29()
	case 30:
		s.answer = easy.SolveProblem30()
	case 31:
		s.answer = easy.SolveProblem31()
	case 32:
		s.answer = easy.SolveProblem32()
	case 33:
		s.answer = easy.SolveProblem33()
	case 34:
		s.answer = easy.SolveProblem34()
	case 35:
		s.answer = easy.SolveProblem35()
	case 36:
		s.answer = easy.SolveProblem36()
	case 37:
		s.answer = easy.SolveProblem37()
	case 38:
		s.answer = easy.SolveProblem38()
	case 39:
		s.answer = easy.SolveProblem39()
	case 40:
		s.answer = easy.SolveProblem40()
	case 41:
		s.answer = easy.SolveProblem41()
	case 42:
		s.answer = easy.SolveProblem42()
	case 43:
		s.answer = easy.SolveProblem43()
	case 44:
		s.answer = easy.SolveProblem44()
	case 45:
		s.answer = easy.SolveProblem45()
		// TEMPLATE_MARKER //
	case 67:
		s.answer = easy.SolveProblem67()
	default:
		s.solved = false
	}

	if s.solved {
		s.duration = time.Since(t0)
		if s.duration > time.Second {
			s.duration = s.duration.Truncate(time.Millisecond)
		} else if s.duration > time.Millisecond {
			s.duration = s.duration.Truncate(time.Microsecond)
		} else if s.duration > time.Microsecond {
			s.duration = s.duration.Truncate(time.Microsecond)
		}
	}

	return s
}

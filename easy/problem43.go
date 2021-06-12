package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
	"github.com/joshprzybyszewski/projecteuler/primes"
)

func SolveProblem43() string {
	/*
		The number, 1406357289, is a 0 to 9 pandigital number
		because it is made up of each of the digits 0 to 9 in
		some order, but it also has a rather interesting
		sub-string divisibility property.

		Let d1 be the 1st digit, d2 be the 2nd digit, and so on.
		In this way, we note the following:

		d2d3d4=406 is divisible by 2
		d3d4d5=063 is divisible by 3
		d4d5d6=635 is divisible by 5
		d5d6d7=357 is divisible by 7
		d6d7d8=572 is divisible by 11
		d7d8d9=728 is divisible by 13
		d8d9d10=289 is divisible by 17

		Find the sum of all 0 to 9 pandigital numbers with this property.
	*/
	ans := getSumOfAllProblem43PropertyNumbers()
	return fmt.Sprintf("%v", ans)
}

func getSumOfAllProblem43PropertyNumbers() int {
	c := &problem43AnswerCache{}
	c.populate()
	fmt.Printf("\tFound: %+v\n", c.answers)

	return mathUtils.Sum(c.answers)
}

func isProblem43Property(digits []int) bool {
	expNumDigits := 10
	if len(digits) != expNumDigits {
		return false
	}
	ps := primes.Below(20)

	numDigitsToSee := 3
	for i, p := range ps {
		if i >= expNumDigits-numDigitsToSee {
			break
		}
		sI := i + 1
		eI := sI + numDigitsToSee
		subNum := mathUtils.DigitsToInt(digits[sI:eI])
		if subNum%p != 0 {
			return false
		}
	}

	return true
}

type problem43AnswerCache struct {
	answers []int
}

func (c *problem43AnswerCache) populate() {
	starting := make([]int, 10)
	toPlace := []int{5, 2, 4, 6, 8, 1, 3, 7, 9}

	c.buildOption(starting, toPlace)
}

func (c *problem43AnswerCache) buildOption(cur, toPlace []int) {
	if len(toPlace) == 0 {
		if isProblem43Property(cur) {
			c.answers = append(c.answers, mathUtils.DigitsToInt(cur))
		}
		return
	}

	tp := toPlace[0]

	for i := 0; i < len(cur); i++ {
		if cur[i] != 0 {
			// this position has been placed
			continue
		}

		if i == 3 {
			// then the tp _must_ be divisible by 2
			if tp%2 != 0 {
				continue
			}
		} else if i == 5 {
			if tp != 0 && tp != 5 {
				continue
			}
		}

		cpy := make([]int, len(cur))
		copy(cpy, cur)
		cpy[i] = tp
		c.buildOption(cpy, toPlace[1:])
	}

}

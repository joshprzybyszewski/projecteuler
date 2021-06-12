package easy

import (
	"fmt"
)

func SolveProblem16() string {
	/*
		2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
		What is the sum of the digits of the number 2^1000?
	*/
	ans := getSumOfTwoToThePower(1000)
	return fmt.Sprintf("%v", ans)
}

func getSumOfTwoToThePower(power int) int {
	pos := make([]int8, power)
	pos[0] = 2
	for i := 1; i < power; i++ {
		timesTwo(pos)
	}
	return sum(pos)
}

func timesTwo(pos []int8) {
	for i := range pos {
		pos[i] *= 2
	}
	for i := range pos {
		if tens := pos[i] / 10; tens > 0 {
			pos[i+1] += tens
			pos[i] %= 10
		}
	}
}

func sum(pos []int8) int {
	total := 0
	for _, p := range pos {
		total += int(p)
	}
	return total
}

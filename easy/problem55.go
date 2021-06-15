package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
)

func SolveProblem55() string {
	/*
		How many Lychrel numbers are there below ten-thousand?
	*/
	ans := getNumLychrelNumbersBelow(10000)
	return fmt.Sprintf("%v", ans)
}

func getNumLychrelNumbersBelow(max int) int {
	total := 0
	for i := 1; i < max; i++ {
		if isLychrelNumber(i) {
			total++
		}
	}
	return total
}

// A number that never forms a palindrome through the reverse
// and add process is called a Lychrel number.
//
// For the purpose of problem 55, we shall assume that a number
// is Lychrel until proven otherwise. In addition you are given
// that for every number below ten-thousand, it will either
//   (i) become a palindrome in less than fifty iterations, or,
//  (ii) no one, with all the computing power that exists, has
// managed so far to map it to a palindrome.
func isLychrelNumber(n int) bool {
	running := n
	for i := 0; i < 50; i++ {
		running += mathUtils.DigitsToInt(
			reverseInts(
				mathUtils.ToDigits(running),
			),
		)
		if isPalindrome(running) {
			return false
		}
	}

	return true
}

func reverseInts(d []int) []int {
	for i := 0; i < len(d)/2; i++ {
		d[i], d[len(d)-1-i] = d[len(d)-1-i], d[i]
	}
	return d
}

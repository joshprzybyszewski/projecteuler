package easy

import (
	"fmt"
	"strconv"

	"github.com/joshprzybyszewski/projecteuler/utils"
)

func SolveProblem4() string {
	/*
		A palindromic number reads the same both ways.
		The largest palindrome made from the product of two
		2-digit numbers is 9009 = 91 × 99.

		Find the largest palindrome made from the product of two 3-digit numbers.
	*/
	p := largestPalindromeInRange(100, 1000)
	return fmt.Sprintf("%d", p)
}

func largestPalindromeInRange(min, max int) int {
	largest := -1
	for a := max - 1; a > min; a-- {
		for b := a; b > min; b-- {
			n := a * b
			if isPalindrome(n) && n > largest {
				largest = n
			}
		}

	}
	return largest
}

func isPalindrome(n int) bool {
	return utils.IsPalindrome(strconv.Itoa(n))
}

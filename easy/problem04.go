package easy

import (
	"fmt"
)

func SolveProblem4() {
	/*
		A palindromic number reads the same both ways.
		The largest palindrome made from the product of two
		2-digit numbers is 9009 = 91 × 99.

		Find the largest palindrome made from the product of two 3-digit numbers.
	*/
	p := largestPalindromeInRange(100, 1000)
	fmt.Printf("Problem 4 Answer: %d\n", p)
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
	s := fmt.Sprintf("%d", n)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

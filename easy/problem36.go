package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/utils"
)

func SolveProblem36() string {
	/*
		The decimal number, 585 = 0x1001001001 (binary),
		is palindromic in both bases.

		Find the sum of all numbers, less than one million,
		which are palindromic in base 10 and base 2.
	*/
	ans := sumOfAllDoublePalindromes(1000000)
	return fmt.Sprintf("%v", ans)
}

func sumOfAllDoublePalindromes(max int) int {
	sum := 0

	for i := 1; i < max; i++ {
		if isPalindrome(i) && isBase2Palindrome(i) {
			sum += i
		}
	}

	return sum
}

func isBase2Palindrome(n int) bool {
	return utils.IsPalindrome(fmt.Sprintf("%b", n))
}

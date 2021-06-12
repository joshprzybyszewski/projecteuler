package easy

import (
	"fmt"
)

func SolveProblem38() string {
	/*
		What is the largest 1 to 9 pandigital 9-digit number
		that can be formed as the concatenated product of an
		integer with (1,2, ... , n) where n > 1?
	*/
	// I found this by writing it out on a post it note
	// obviously `9` produces `918273645`, so we need something
	// which starts with a 9 and then has a number > 1 in the
	// second position.
	// but we also know we need n > 1, so it can't be `987654321`.
	// And if we have 9 in the first position, then we get 18 as
	// the leading values in n = 2, so we have 6 more digits to use.
	// This means we _must_ use n = 2. therefore, we need to allocate
	// 2, 3, 4, 5, 6, and 7 such that three go in n = 1 and three go
	// in n = 2. We know the second digit < 5, otherwise n = 2 will
	// bump the leading 18 to a 19 and we can't have that.
	// so the second digit is either 2, 3, or 4. It can't be four because
	// that would make n = 2 lead with 188 or 189, and that's no dice.
	// let's try 3: so 93AB && 18(6|7)CD where A=B=C=D = (2|4|5|6|7).
	// We know D = (2B)%10, so that means either D=4 && B=(2|7) OR
	// D=2 && B=6.
	ans := `932718654`
	return fmt.Sprintf("%v", ans)
}

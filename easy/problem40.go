package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
)

func SolveProblem40() string {
	/*
		An irrational decimal fraction is created by concatenating
		the positive integers:

		0.123456789101112131415161718192021...
		..1...5....0.^..5....0....5....0...

		It can be seen that the 12th digit of the fractional part is 1.

		If dn represents the nth digit of the fractional part,
		find the value of the following expression.

		d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000
	*/
	ans := getProblem40Answer()
	return fmt.Sprintf("%v", ans)
}

func getProblem40Answer() int {
	digits := []int{
		getChampernowneDigit(1),
		getChampernowneDigit(10),
		getChampernowneDigit(100),
		getChampernowneDigit(1000),
		getChampernowneDigit(10000),
		getChampernowneDigit(100000),
		getChampernowneDigit(1000000),
	}
	return mathUtils.Product(digits)
}

func getChampernowneDigit(i int) int {
	if i <= 0 {
		return 0
	}

	represents, index := getNumberRepresentedAtChampernowneDigit(i)

	return mathUtils.ToDigits(represents)[index]
}

func getNumberRepresentedAtChampernowneDigit(i int) (int, int) {
	if i <= 0 {
		return 0, 0
	}

	numberMin, numberMax := 1, 10
	indexMin, indexMax := 1, 10

	numDigits := len(mathUtils.ToDigits(numberMin))

	for indexMax <= i {
		numberMin = numberMax
		numberMax *= 10
		indexMin = indexMax

		numDigits = len(mathUtils.ToDigits(numberMin))
		numberDiff := numberMax - numberMin
		indexMax += numDigits * numberDiff
	}

	past := i - indexMin

	number := numberMin + (past / numDigits)
	numberDigits := mathUtils.ToDigits(number)

	return number, past % len(numberDigits)
}

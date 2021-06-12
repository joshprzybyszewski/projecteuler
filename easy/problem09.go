package easy

import (
	"fmt"
)

func SolveProblem9() string {
	/*
		A Pythagorean triplet is a set of three natural
		numbers, a < b < c, for which,

		a^2 + b^2 = c^2
		For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

		There exists exactly one Pythagorean triplet for which a + b + c = 1000.
		Find the product abc.
	*/
	a, b, c := getThousandPythagoreanTriplet()
	return fmt.Sprintf("%d = %d * %d * %d", a*b*c, a, b, c)
}

func getThousandPythagoreanTriplet() (int, int, int) {
	thousand := 1000
	for a := 1; a < thousand; a++ {
		for b := a + 1; a+b < thousand; b++ {
			c := thousand - b - a
			if isPythagoreanTriplet(a, b, c) {
				return a, b, c
			}
		}
	}
	return 0, 0, 0
}

func isPythagoreanTriplet(a, b, c int) bool {
	a2 := a * a
	b2 := b * b
	c2 := c * c
	switch {
	case c > a && c > b:
		return a2+b2 == c2
	case b > c && b > a:
		return a2+c2 == b2
	default:
		return c2+b2 == a2
	}
}

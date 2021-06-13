package easy

import (
	"fmt"
	"strconv"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
	"github.com/joshprzybyszewski/projecteuler/primes"
)

func SolveProblem48() string {
	/*
		The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.

		Find the last ten digits of the series,
		1^1 + 2^2 + 3^3 + ... + 1000^1000.
	*/
	ans := getNToTheNSeriesAnswer(1000)

	return fmt.Sprintf("%v", ans)
}

func getNToTheNSeriesAnswer(max int) int {
	sum := 0
	for n := 1; n <= max; n++ {
		factors := getNToTheNFactors(n)
		product := mathUtils.BigProduct(factors)
		str := product.String()
		if len(str) > 10 {
			str = str[len(str)-11:]
		}
		v, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		sum += v
		sum %= 100000000000
	}

	return sum
}

func getNToTheNFactors(n int) []int {
	base := primes.Factors(n)

	output := make([]int, 0, len(base)*n)
	for i := 0; i < n; i++ {
		output = append(output, base...)
	}
	return output
}

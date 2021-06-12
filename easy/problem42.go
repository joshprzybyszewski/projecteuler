package easy

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/joshprzybyszewski/projecteuler/sequence"
)

const (
	problem42TriangleFileName = `easy/p042_words.txt`
)

func SolveProblem42() string {
	/*
		The nth term of the sequence of triangle numbers is
		given by, tn = ½n(n+1); so the first ten triangle
		numbers are:

		1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

		Using words.txt, a 16K text file containing nearly two-thousand
		common English words, how many are triangle words?
	*/
	words := getProbelm42Words()
	ans := getCountOfTriangleWords(words)
	return fmt.Sprintf("%v", ans)
}

func getProbelm42Words() []string {
	dat, err := ioutil.ReadFile(problem42TriangleFileName)
	if err != nil {
		log.Fatalf("problem 42 cannot read file: %v", err)
	}
	names := strings.Split(string(dat), ",")

	for i, name := range names {
		names[i] = strings.Trim(name, `"`)
	}

	return names
}

func getCountOfTriangleWords(words []string) int {
	sum := 0
	for _, word := range words {
		if sequence.Triangular.Is(
			getAlphabeticalValue(word),
		) {
			sum++
		}
	}

	return sum
}

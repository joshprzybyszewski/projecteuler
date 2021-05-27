package easy

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

const (
	problem22NamesFileName = `easy/p022_names.txt`

	lowerAByte = byte('a')
	lowerZByte = byte('z')
	capAByte   = byte('A')
	capZByte   = byte('Z')
)

func SolveProblem22() {
	/*
		Using names.txt (right click and 'Save Link/Target As...'),
		a 46K text file containing over five-thousand first names,
		begin by sorting it into alphabetical order. Then working
		out the alphabetical value for each name, multiply this
		value by its alphabetical position in the list to obtain
		a name score.

		For example, when the list is sorted into alphabetical
		order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53,
		is the 938th name in the list. So, COLIN would obtain
		a score of 938 × 53 = 49714.

		What is the total of all the name scores in the file?
	*/

	names := getProblem22Names()

	ans := getProblem22Answer(names)
	fmt.Printf("Problem 22 Answer: %v\n", ans)
}

func getProblem22Names() []string {
	dat, err := ioutil.ReadFile(problem22NamesFileName)
	if err != nil {
		log.Fatalf("problem 22 cannot read file: %v", err)
	}
	names := strings.Split(string(dat), ",")

	for i, name := range names {
		names[i] = strings.Trim(name, `"`)
	}

	return names
}

func getProblem22Answer(names []string) int {
	sort.Strings(names)
	sum := 0
	for i, name := range names {
		sum += ((i + 1) * getAlphabeticalValue(name))
	}
	return sum
}

func getAlphabeticalValue(name string) int {
	sum := 0
	for i := 0; i < len(name); i++ {
		l := name[i]
		if capAByte <= l && l <= capZByte {
			sum += int(1 + l - capAByte)
		} else if lowerAByte <= l && l <= lowerZByte {
			sum += int(1 + l - lowerAByte)
		}
	}
	return sum
}

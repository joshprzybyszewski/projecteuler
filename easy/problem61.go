package easy

import (
	"fmt"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
	"github.com/joshprzybyszewski/projecteuler/sequence"
	"github.com/joshprzybyszewski/projecteuler/utils"
)

func SolveProblem61() string {
	/*

		Triangle, square, pentagonal, hexagonal, heptagonal, and octagonal
		numbers are all figurate (polygonal) numbers and are generated by
		the following formulae:

		Triangle	 	P3(n)=n(n+1)/2	 	1, 3, 6, 10, 15, ...
		Square	 		P4(n)=n^2	 		1, 4, 9, 16, 25, ...
		Pentagonal	 	P5(n)=n(3n−1)/2	 	1, 5, 12, 22, 35, ...
		Hexagonal	 	P6(n)=n(2n−1)	 	1, 6, 15, 28, 45, ...
		Heptagonal	 	P7(n)=n(5n−3)/2	 	1, 7, 18, 34, 55, ...
		Octagonal	 	P8(n)=n(3n−2)	 	1, 8, 21, 40, 65, ...

		The ordered set of three 4-digit numbers: 8128, 2882, 8281,
		has three interesting properties.
		1. The set is cyclic, in that the last two digits of each number
			is the first two digits of the next number (including the
			last number with the first).
		2. Each polygonal type: triangle (P3,127=8128), square (P4,91=8281),
			and pentagonal (P5,44=2882), is represented by a different
			number in the set.
		3. This is the only set of 4-digit numbers with this property.

		Find the sum of the only ordered set of six cyclic 4-digit
		numbers for which each polygonal type: triangle, square,
		pentagonal, hexagonal, heptagonal, and octagonal, is represented
		by a different number in the set.
	*/
	ans := getFourDigitPolygonalCycle(6)
	return fmt.Sprintf("%v = sum(%v)", mathUtils.Sum(ans), ans)
}

func getFourDigitPolygonalCycle(cycleLen int) []int {
	for fourDigits := 9999; fourDigits >= 1000; fourDigits-- {
		pc := newPolygonalCycle(fourDigits, cycleLen)
		res := pc.buildComplete()
		if res.isComplete() {
			return res.fourDigitNums
		}
	}
	return nil
}

var (
	allPolygonalSequences = []sequence.Sequence{
		sequence.Triangular,
		sequence.Square,
		sequence.Pentagonal,
		sequence.Hexagonal,
		sequence.Heptagonal,
		sequence.Octagonal,
	}
)

type polygonalCycle struct {
	fourDigitNums []int
	satisfies     []bool
}

func newPolygonalCycle(start, goal int) *polygonalCycle {
	pc := &polygonalCycle{
		fourDigitNums: []int{start},
		satisfies:     make([]bool, goal),
	}

	isAny := false
	for i, seq := range allPolygonalSequences[:len(pc.satisfies)] {
		if seq.Is(start) {
			isAny = true
			pc.satisfies[i] = true
		}
	}
	if !isAny {
		return nil
	}

	return pc
}

func (pc *polygonalCycle) buildComplete() *polygonalCycle {
	if pc == nil {
		return nil
	}
	if pc.isComplete() {
		return pc
	}

	for twoDigits := 10; twoDigits < 100; twoDigits++ {
		res := pc.buildNext(twoDigits).buildComplete()
		if res.isComplete() {
			return res
		}
	}

	return nil
}

func (pc *polygonalCycle) isComplete() bool {
	if pc == nil || len(pc.fourDigitNums) < len(pc.satisfies) {
		return false
	}

	for _, sat := range pc.satisfies {
		if !sat {
			return false
		}
	}

	return pc.fourDigitNums[len(pc.fourDigitNums)-1]%100 == pc.fourDigitNums[0]/100
}

func (pc *polygonalCycle) buildNext(twoDigit int) *polygonalCycle {
	nextFour := ((pc.fourDigitNums[len(pc.fourDigitNums)-1] % 100) * 100) + twoDigit

	if nextFour < 1000 || utils.IntContains(pc.fourDigitNums, nextFour) {
		return nil
	}

	satCpy := make([]bool, len(pc.satisfies))
	copy(satCpy, pc.satisfies)

	wasAny := false
	for i, seq := range allPolygonalSequences[:len(pc.satisfies)] {
		if seq.Is(nextFour) {
			wasAny = true
			satCpy[i] = true
		}
	}
	if !wasAny {
		// number does not fall in any of the sequences.
		return nil
	}

	numSeqs := 0
	for _, sat := range satCpy {
		if sat {
			numSeqs++
		}
	}
	if numSeqs != len(pc.fourDigitNums)+1 {
		return nil
	}

	sCpy := make([]int, len(pc.fourDigitNums), len(pc.fourDigitNums)+1)
	copy(sCpy, pc.fourDigitNums)
	sCpy = append(sCpy, nextFour)

	return &polygonalCycle{
		fourDigitNums: sCpy,
		satisfies:     satCpy,
	}
}

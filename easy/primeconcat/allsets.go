package primeconcat

import (
	"github.com/joshprzybyszewski/projecteuler/primes"
)

type AllSets struct {
	bySize [][]*set
}

func NewAllSets() *AllSets {
	return &AllSets{
		bySize: [][]*set{
			nil, // 0
			nil, // 1
		},
	}
}

func (as *AllSets) GetMinSet(n int) []int {
	as.buildUntilFirstWithSize(n)
	if len(as.bySize) <= n || len(as.bySize[n]) == 0 {
		return nil
	}
	return as.bySize[n][0].elems
}

func (as *AllSets) buildUntilFirstWithSize(
	size int,
) {
	for len(as.bySize) <= size {
		np := as.getNextPrime()

		as.bySize[1] = append(as.bySize[1], newSingleElemSet(np))

		as.addElemToAllExisting(np)
	}
}

func (as *AllSets) getNextPrime() int {
	return primes.GetNth(uint(len(as.bySize[1]) + 1))
}

func (as *AllSets) addElemToAllExisting(
	nextPrime int,
) {
	for size := len(as.bySize) - 1; size > 0; size-- {
		setsOfSize := as.bySize[size]
		for i := len(setsOfSize) - 1; i >= 0; i-- {
			as.add(
				setsOfSize[i].add(nextPrime),
			)
		}
	}
}

func (as *AllSets) add(s *set) {
	if s == nil {
		return
	}

	ss := s.size()

	for len(as.bySize) <= ss {
		as.bySize = append(as.bySize, make([]*set, 0, 16))
	}

	as.bySize[ss] = append(as.bySize[ss], s)
}

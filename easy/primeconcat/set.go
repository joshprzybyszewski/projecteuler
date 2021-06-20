package primeconcat

import (
	"log"
	"strconv"

	"github.com/joshprzybyszewski/projecteuler/primes"
)

type set struct {
	elems []int
}

func newSet(elems []int) *set {
	return &set{
		elems: elems,
	}
}

func newSingleElemSet(elem int) *set {
	return &set{
		elems: []int{elem},
	}
}

func (s *set) size() int {
	return len(s.elems)
}

func (s *set) add(
	with int,
) *set {
	if s == nil {
		return nil
	}

	wStr := strconv.Itoa(with)

	for i := len(s.elems) - 1; i >= 0; i-- {
		e := s.elems[i]
		if e == with {
			return nil
		}

		eStr := strconv.Itoa(e)

		we, err := strconv.Atoi(wStr + eStr)
		if err != nil {
			log.Fatalf("test prime failed: %v", err)
		}
		if !primes.Is(we) {
			return nil
		}

		ew, err := strconv.Atoi(eStr + wStr)
		if err != nil {
			log.Fatalf("test prime failed: %v", err)
		}
		if !primes.Is(ew) {
			return nil
		}
	}

	elemsCpy := make([]int, len(s.elems), len(s.elems)+1)
	copy(elemsCpy, s.elems)
	elemsCpy = append(elemsCpy, with)

	return newSet(elemsCpy)
}

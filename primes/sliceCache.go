package primes

import (
	"math"
	"sort"
	"strconv"
)

var _ cacher = (*sliceCache)(nil)

var (
	maxSliceCacheSize = 1 << 16
)

type sliceCache struct {
	primes []int
	known  int
}

func newSliceCache() *sliceCache {
	return &sliceCache{
		primes: []int{2, 3, 5},
		known:  5,
	}
}

func newSliceCacheFromFile(lines []string) *sliceCache {
	sc := &sliceCache{
		primes: make([]int, 0, len(lines)),
	}

	for _, line := range lines {
		p, err := strconv.Atoi(line)
		if err != nil {
			break
		}
		sc.primes = append(sc.primes, p)
	}
	if len(sc.primes) == 0 {
		return newSliceCache()
	}

	sc.known = sc.primes[len(sc.primes)-1]

	return sc
}

func (m *sliceCache) knownToString() []string {
	ret := make([]string, 0, len(m.primes))
	for _, p := range m.primes {
		ret = append(ret, strconv.Itoa(p))
	}
	return ret
}

func (m *sliceCache) buildTo(n int) {
	for m.known < n && len(m.primes) < maxSliceCacheSize {
		m.known++

		if m.knownIsPrime() {
			m.primes = append(m.primes, m.known)
		}
	}
}

func (m *sliceCache) knownIsPrime() bool {
	max := int(math.Floor(math.Sqrt(float64(m.known))))

	for _, p := range m.primes {
		if m.known%p == 0 {
			return false
		}
		if p > max {
			break
		}
	}

	return true
}

func (m *sliceCache) is(n int) bool {
	m.buildTo(n)

	if n > m.known {
		return m.isPrimeBeyondCache(n)
	}

	i := sort.Search(len(m.primes), func(i int) bool {
		return n <= m.primes[i]
	})
	return i < len(m.primes) && m.primes[i] == n
}

func (m *sliceCache) isPrimeBeyondCache(n int) bool {
	maxToCheck := int(math.Floor(math.Sqrt(float64(n))))

	for _, p := range m.primes {
		if n%p == 0 {
			return false
		}
	}

	for d := m.known + 1; d <= maxToCheck; d++ {
		if n%d == 0 {
			return false
		}
	}

	return true
}

func (m *sliceCache) below(max int) []int {
	m.buildTo(max)

	b := m.knownBelow(max)

	if max <= m.known {
		// the desired max is below what we know about
		// let's just return the slice of known numbers
		return b
	}

	// we need to build up the primes that are past our current
	// upper-bound on the cache.
	for n := b[len(b)-1]; n < max; n++ {
		if m.isPrimeBeyondCache(n) {
			b = append(b, n)
		}
	}

	return b
}

func (m *sliceCache) knownBelow(max int) []int {

	l := make([]int, 0, len(m.primes))

	for _, p := range m.primes {
		if p < max {
			l = append(l, p)
		}
	}

	return l
}

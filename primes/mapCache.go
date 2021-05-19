package primes

import (
	"math"
	"sort"
)

var _ cacher = (*mapCache)(nil)

var (
	maxMapCacheSize = 1 << 16
)

type mapCache struct {
	primeCache map[int]struct{}
	known      int
}

func newMapCache() *mapCache {
	return &mapCache{
		primeCache: map[int]struct{}{
			2: {},
			3: {},
			5: {},
		},
		known: 5,
	}
}

func (m *mapCache) buildTo(n int) {
	for m.known < n && len(m.primeCache) < maxMapCacheSize {
		m.known++

		if m.knownIsPrime() {
			m.primeCache[m.known] = struct{}{}
		}
	}
}

func (m *mapCache) knownIsPrime() bool {
	for p := range m.primeCache {
		if m.known%p == 0 {
			return false
		}
	}

	return true
}

func (m *mapCache) is(n int) bool {
	m.buildTo(n)

	if n > m.known {
		return m.isPrimeBeyondCache(n)
	}

	if _, ok := m.primeCache[n]; ok {
		return true
	}
	return false
}

func (m *mapCache) isPrimeBeyondCache(n int) bool {
	for p := range m.primeCache {
		if n%p == 0 {
			return false
		}
	}

	maxToCheck := int(math.Floor(math.Sqrt(float64(n))))

	for d := m.known + 1; d <= maxToCheck; d++ {
		if n%d == 0 {
			return false
		}
	}

	return true
}

func (m *mapCache) below(max int) []int {
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

func (m *mapCache) knownBelow(max int) []int {

	l := make([]int, 0, len(m.primeCache))

	for p := range m.primeCache {
		if p < max {
			l = append(l, p)
		}
	}

	sort.Ints(l)
	return l
}

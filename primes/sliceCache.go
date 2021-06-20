package primes

import (
	"log"
	"math"
	"sort"
	"strconv"
	"sync"
)

var _ cacher = (*sliceCache)(nil)

var (
	maxSliceCacheSize = 1 << 31
)

type sliceCache struct {
	primes []int
	known  int

	lock *sync.RWMutex
}

func newSliceCache() *sliceCache {
	return &sliceCache{
		primes: []int{2, 3, 5},
		known:  5,
		lock:   &sync.RWMutex{},
	}
}

func newSliceCacheFromFile(lines []string) *sliceCache {
	sc := &sliceCache{
		primes: make([]int, 0, len(lines)),
		lock:   &sync.RWMutex{},
	}

	for _, line := range lines {
		p, err := strconv.Atoi(line)
		if err != nil {
			break
		}
		sc.primes = append(sc.primes, p)
	}
	log.Printf("loaded %d primes from file\n", len(lines))
	if len(sc.primes) == 0 {
		return newSliceCache()
	}

	sc.known = sc.primes[len(sc.primes)-1]

	return sc
}

func (m *sliceCache) knownToString() []string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	ret := make([]string, 0, len(m.primes))
	for _, p := range m.primes {
		ret = append(ret, strconv.Itoa(p))
	}
	return ret
}

func (m *sliceCache) needsToBuild(n int) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.known < n || len(m.primes) >= maxSliceCacheSize
}

func (m *sliceCache) buildTo(n int) {
	if !m.needsToBuild(n) {
		return
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	for m.known < n && len(m.primes) < maxSliceCacheSize {
		m.known++

		if m.knownIsPrime() {
			m.primes = append(m.primes, m.known)
		}
	}
}

func (m *sliceCache) knownIsPrime() bool {
	// don't lock in here, because we have a lock already.
	// also don't use this method when you don't have a lock.
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

	m.lock.RLock()
	defer m.lock.RUnlock()

	i := sort.Search(len(m.primes), func(i int) bool {
		return n <= m.primes[i]
	})
	return i < len(m.primes) && m.primes[i] == n
}

func (m *sliceCache) isPrimeBeyondCache(n int) bool {
	maxToCheck := int(math.Floor(math.Sqrt(float64(n))))

	m.lock.RLock()
	defer m.lock.RUnlock()

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

func (m *sliceCache) nth(n uint) int {
	i := n - 1
	if int(i) >= maxSliceCacheSize {
		return -1
	}
	m.buildUntil(n)

	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.primes[i]
}

func (m *sliceCache) buildUntil(cacheSize uint) {
	if m.hasNElems(cacheSize) {
		return
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	for int(cacheSize) >= len(m.primes) && len(m.primes) < maxSliceCacheSize {
		m.known++

		if m.knownIsPrime() {
			m.primes = append(m.primes, m.known)
		}
	}
}

func (m *sliceCache) hasNElems(size uint) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return int(size) < len(m.primes)
}

func (m *sliceCache) below(max int) []int {
	m.buildTo(max)

	b, hadAll := m.knownBelow(max)

	if hadAll {
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

func (m *sliceCache) knownBelow(max int) ([]int, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	numToCopy := sort.Search(len(m.primes), func(i int) bool {
		return m.primes[i] >= max
	})

	cpy := make([]int, numToCopy)
	copy(cpy, m.primes)

	return cpy, max < m.known
}

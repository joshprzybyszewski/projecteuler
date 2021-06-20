package sequence

import (
	"sort"
	"sync"
)

type generator interface {
	generate(uint) int
}

type sequenceCache struct {
	lock *sync.RWMutex

	generator generator
	numbers   []int
}

func newSequenceCache(
	gen generator,
) *sequenceCache {
	return &sequenceCache{
		numbers:   make([]int, 0, 16),
		lock:      &sync.RWMutex{},
		generator: gen,
	}
}

func (c *sequenceCache) getNthFromCache(n uint) (int, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	i := int(n) - 1
	if i < len(c.numbers) {
		return c.numbers[i], true
	}
	return 0, false
}

func (c *sequenceCache) generateTo(
	n uint,
) {

	c.lock.Lock()
	defer c.lock.Unlock()

	for gn := uint(len(c.numbers) + 1); gn <= n; gn++ {
		next := c.generator.generate(gn)
		c.numbers = append(c.numbers, next)
	}
}

func (c *sequenceCache) generatePast(
	past int,
) {

	c.lock.Lock()
	defer c.lock.Unlock()

	var next int
	if len(c.numbers) > 0 {
		next = c.numbers[len(c.numbers)-1]
	}

	for gn := uint(len(c.numbers) + 1); next < past; gn++ {
		next = c.generator.generate(gn)
		c.numbers = append(c.numbers, next)
	}
}

func (c *sequenceCache) GetNth(n uint) int {
	v, ok := c.getNthFromCache(n)
	if ok {
		return v
	}
	c.generateTo(n)

	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.numbers[n-1]
}

func (c *sequenceCache) isInCache(d int) (bool, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if len(c.numbers) == 0 || c.numbers[len(c.numbers)-1] < d {
		return false, false
	}

	i := sort.Search(len(c.numbers), func(i int) bool {
		return d <= c.numbers[i]
	})
	exists := i < len(c.numbers) && c.numbers[i] == d

	return exists, true
}

func (c *sequenceCache) Is(d int) bool {

	is, ok := c.isInCache(d)
	if ok {
		return is
	}

	c.generatePast(d)
	is, ok = c.isInCache(d)
	if ok {
		return is
	}

	panic(`max slice not implemented`)
	c.lock.RLock()
	start := len(c.numbers)
	c.lock.RUnlock()

	for i := uint(start); ; i++ {
		pn := c.generator.generate(i)
		if pn == d {
			return true
		}
		if pn > d {
			return false
		}
	}
}

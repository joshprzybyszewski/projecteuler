package sequence

import (
	"sort"
	"sync"
)

var (
	Triangular ShapeSequence = newShapeSequenceCache(triangle)
	Pentagonal ShapeSequence = newShapeSequenceCache(pentagon)
	Hexagonal  ShapeSequence = newShapeSequenceCache(hexagon)
)

type ShapeSequence interface {
	GetNth(uint) int
	Is(int) bool
}

type shapeSequenceGenerator interface {
	generate(uint) int
}

type shapeSequenceCache struct {
	lock *sync.RWMutex

	generator shapeSequenceGenerator
	numbers   []int
}

func newShapeSequenceCache(
	generator shapeSequenceGenerator,
) *shapeSequenceCache {
	return &shapeSequenceCache{
		numbers:   make([]int, 0, 16),
		lock:      &sync.RWMutex{},
		generator: generator,
	}
}

func (c *shapeSequenceCache) getNthFromCache(n uint) (int, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	i := int(n) - 1
	if i < len(c.numbers) {
		return c.numbers[i], true
	}
	return 0, false
}

func (c *shapeSequenceCache) generateTo(
	n uint,
) {

	c.lock.Lock()
	defer c.lock.Unlock()

	for gn := uint(len(c.numbers) + 1); gn <= n; gn++ {
		next := c.generator.generate(gn)
		c.numbers = append(c.numbers, next)
	}
}

func (c *shapeSequenceCache) generatePast(
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

func (c *shapeSequenceCache) GetNth(n uint) int {
	v, ok := c.getNthFromCache(n)
	if ok {
		return v
	}
	c.generateTo(n)

	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.numbers[n-1]
}

func (c *shapeSequenceCache) isInCache(d int) (bool, bool) {
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

func (c *shapeSequenceCache) Is(d int) bool {

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

type shapeSides int

func (s shapeSides) generate(n uint) int {
	switch s {
	case triangle:
		return int(n * (n + 1) / 2)
	case pentagon:
		return int(n * (3*n - 1) / 2)
	case hexagon:
		return int(n * (2*n - 1))
	default:
		return 0
	}
}

const (
	triangle shapeSides = 3
	pentagon shapeSides = 5
	hexagon  shapeSides = 6
)

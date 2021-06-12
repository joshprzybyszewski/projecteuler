package primes

var (
	cache cacher = newSliceCache()
)

type cacher interface {
	is(int) bool
	below(int) []int

	knownToString() []string
}

func Is(n int) bool {
	if n <= 1 {
		return false
	}

	return cache.is(n)
}

func Below(max int) []int {
	return cache.below(max)
}

func Factors(k int) []int {
	if k < 1 {
		return nil
	} else if k == 1 {
		return []int{1}
	}
	primes := Below(k + 1)
	factors := make([]int, 0, 4)

	cur := k
	for _, p := range primes {
		for cur%p == 0 && cur > 1 {
			factors = append(factors, p)
			cur = cur / p
		}
		if cur == 1 {
			break
		}
	}

	return factors
}

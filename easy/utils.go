package easy

import "math"

const (
	maxCacheValue = 200000
)

var (
	primeCache   = make(map[int]struct{}, maxCacheValue/100)
	maxPrimeSeen int
)

func init() {
	primeCache[2] = struct{}{}
	primeCache[3] = struct{}{}
	primeCache[5] = struct{}{}
	maxPrimeSeen = 5
	buildCacheTo(10000)
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n >= maxCacheValue {
		buildCacheTo(maxCacheValue)
		return longerIsPrime(n)
	} else if _, ok := primeCache[n]; ok {
		return true
	} else if n < maxPrimeSeen {
		return false
	}

	buildCacheTo(n)

	_, ok := primeCache[n]
	return ok
}

func buildCacheTo(
	n int,
) {

	for ; maxPrimeSeen <= n && maxPrimeSeen <= maxCacheValue; maxPrimeSeen++ {
		if longIsPrime(maxPrimeSeen) {
			primeCache[maxPrimeSeen] = struct{}{}
		}
	}
}

func longIsPrime(n int) bool {
	for p := range primeCache {
		if n%p == 0 {
			return false
		}
	}

	return true
}

func longerIsPrime(n int) bool {
	for p := range primeCache {
		if n%p == 0 {
			return false
		}
	}

	maxToCheck := int(math.Floor(math.Sqrt(float64(n))))

	for d := maxPrimeSeen; d <= maxToCheck; d++ {
		if n%d == 0 {
			return false
		}
	}

	return true
}

func GetPrimes(max int) []int {
	p := make([]int, 0, 4)
	for i := 2; i <= max; i++ {
		if IsPrime(i) {
			p = append(p, i)
		}
	}
	return p
}

func GetPrimeFactorization(k int) []int {
	factors := make([]int, 0, 4)
	primes := GetPrimes(k)

	for cur := k; cur > 1; {
		for _, p := range primes {
			if cur%p == 0 {
				factors = append(factors, p)
				cur = cur / p
				break
			}
		}
	}

	return factors
}

func getCountOfNums(s []int) map[int]int {
	m := make(map[int]int, len(s))
	for _, e := range s {
		m[e]++
	}

	return m
}

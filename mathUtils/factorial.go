package mathUtils

var (
	cacheFactorials = &factorialCache{
		values: []int{1, 1, 2},
	}
)

type factorialCache struct {
	values []int
}

func (fc *factorialCache) get(n int) int {
	if n < 0 {
		return 0
	}
	if n < len(fc.values) {
		return fc.values[n]
	}

	newValues := make([]int, n+1)
	copy(newValues, fc.values)
	for i := len(fc.values); i < len(newValues); i++ {
		newValues[i] = newValues[i-1] * i
	}
	fc.values = newValues

	return fc.values[n]
}

func Factorial(n int) int {
	return cacheFactorials.get(n)
}

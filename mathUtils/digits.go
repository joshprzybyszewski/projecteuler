package mathUtils

const (
	zeroByte = byte('0')
	nineByte = byte('9')
)

func ToDigits(n int) []int {
	if n < 0 {
		return nil
	}
	return Uint64ToDigits(uint64(n))
}

func Uint64ToDigits(n uint64) []int {
	if n == 0 {
		return []int{0}
	}

	tmp := make([]int, 0, 4)
	for n > 0 {
		tmp = append(tmp, int(n%10))
		n /= 10
	}
	if len(tmp) <= 1 {
		return tmp
	}

	res := make([]int, len(tmp))
	for i := 0; i < len(tmp); i++ {
		res[len(res)-1-i] = tmp[i]
	}
	return res
}

func DigitsToInt(digits []int) int {
	if len(digits) == 0 {
		return 0
	}
	total := 0
	for _, d := range digits {
		total *= 10
		total += d
	}
	return total
}

func StringToDigits(str string) []int {
	res := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		if str[i] < zeroByte || str[i] > nineByte {
			return nil
		}
		res[i] = int(str[i] - zeroByte)
	}
	return res
}

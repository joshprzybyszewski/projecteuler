package mathUtils

import "strconv"

const (
	zeroByte = byte('0')
)

func ToDigits(n int) []int {
	str := strconv.Itoa(n)
	return StringToDigits(str)
}

func StringToDigits(str string) []int {
	res := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		res[i] = int(str[i] - zeroByte)
	}
	return res
}

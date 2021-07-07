package utils

func Dedupe(input []int) []int {
	out := make([]int, 0, len(input))
	outMap := make(map[int]struct{}, len(input))
	for _, e := range input {
		if _, ok := outMap[e]; !ok {
			out = append(out, e)
			outMap[e] = struct{}{}
		}
	}
	return out
}

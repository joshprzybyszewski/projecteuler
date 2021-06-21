package sequence

func init() {
	for p := firstPower; p < maxPower; p++ {
		allPowerSequences[p] = newSequenceCache(p)
	}

	Squares = allPowerSequences[2]
	Cubes = allPowerSequences[3]
}

var (
	allPowerSequences = make([]Sequence, int(maxPower)+1)

	Squares Sequence
	Cubes   Sequence
)

func GetPower(p int) Sequence {
	if p <= 0 || p >= len(allPowerSequences) {
		return nil
	}
	return allPowerSequences[p]
}

type powers int

func (s powers) generate(n uint) int {
	if s < firstPower || s > maxPower {
		panic(`dev error`)
	}
	output := n
	for i := firstPower; i < s; i++ {
		output *= n
	}
	return int(output)
}

const (
	firstPower powers = 1
	maxPower   powers = 20
)

package sequence

var (
	Squares Sequence = newSequenceCache(squares)
	Cubes   Sequence = newSequenceCache(cubes)
	Quads   Sequence = newSequenceCache(quads)
)

type powers int

func (s powers) generate(n uint) int {
	switch s {
	case squares:
		return int(n * n)
	case cubes:
		return int(n * n * n)
	case quads:
		return int(n * n * n * n)
	default:
		return 0
	}
}

const (
	squares powers = 2
	cubes   powers = 3
	quads   powers = 4
)

package sequence

var (
	Triangular Sequence = newSequenceCache(triangle)
	Square     Sequence = newSequenceCache(square)
	Pentagonal Sequence = newSequenceCache(pentagon)
	Hexagonal  Sequence = newSequenceCache(hexagon)
	Heptagonal Sequence = newSequenceCache(heptagon)
	Octagonal  Sequence = newSequenceCache(octagon)
)

type shapeSides int

func (s shapeSides) generate(n uint) int {
	switch s {
	case triangle:
		return int(n * (n + 1) / 2)
	case square:
		return int(n * n)
	case pentagon:
		return int(n * (3*n - 1) / 2)
	case hexagon:
		return int(n * (2*n - 1))
	case heptagon:
		return int(n * (5*n - 3) / 2)
	case octagon:
		return int(n * (3*n - 2))
	default:
		return 0
	}
}

const (
	triangle shapeSides = 3
	square   shapeSides = 4
	pentagon shapeSides = 5
	hexagon  shapeSides = 6
	heptagon shapeSides = 7
	octagon  shapeSides = 8
)

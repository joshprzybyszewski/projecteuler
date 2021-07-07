package sequence

type Sequence interface {
	GetNth(uint) int
	Is(int) bool

	// returns a 1-indexed position
	GetPosition(int) (uint, bool)
}

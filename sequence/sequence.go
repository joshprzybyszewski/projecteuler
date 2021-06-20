package sequence

type Sequence interface {
	GetNth(uint) int
	Is(int) bool
}

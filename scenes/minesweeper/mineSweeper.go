package minesweeper

const id string = "mineSweeper"

type MineSweeper struct {
	rows          int
	colomns       int
	bombsNumber   int
	bombsPosition []int
	field         [][]int
}

const (
	zero = iota
	one
	two
	three
	four
	five
	six
	seven
	eight
	close
	bomb
	flag
)

var (
	nums      [9]int = [9]int{zero, one, two, three, four, five, six, seven, eight}
	nextCheck []int
)

func (m *MineSweeper) GetId() string {
	return id
}

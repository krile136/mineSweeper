package minesweeper

import "github.com/krile136/mineSweeper/types/route"

const routeType route.RouteType = route.MineSweeper

type MineSweeper struct {
	rows          int
	columns       int
	bombsNumber   int
	bombsPosition []int
	field         [][]int
}

const (
	blockWidth = 32
	zero       = iota
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
	nums            [9]int = [9]int{zero, one, two, three, four, five, six, seven, eight}
	nextCheck       []int
	scrollX         float64
	scrollY         float64
	maxScrollX      float64
	maxScrollY      float64
	barLengthX      float64
	barLengthY      float64
	barSlideX       float64
	barSlideY       float64
	isBarDisplay    bool
	BarDisplayFrame int

	messages []message

	player status
	enemy  status

	GetExp int
)

func (m *MineSweeper) GetRouteType() route.RouteType {
	return routeType
}

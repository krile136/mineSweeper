package minesweeper

import (
	"github.com/krile136/mineSweeper/enum/route"
	"github.com/krile136/mineSweeper/scenes/minesweeper/character/characterDraw"
	"github.com/krile136/mineSweeper/scenes/minesweeper/character/characterStatus"
	"github.com/krile136/mineSweeper/scenes/minesweeper/explode"
	"github.com/krile136/mineSweeper/scenes/minesweeper/message/messages"
)

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
	nums             [9]int = [9]int{zero, one, two, three, four, five, six, seven, eight}
	nextCheck        []int
	scrollX          float64
	scrollY          float64
	maxScrollX       float64
	maxScrollY       float64
	barLengthX       float64
	barLengthY       float64
	barSlideX        float64
	barSlideY        float64
	isBarDisplay     bool
	BarDisplayFrame  int
	currentCombo     int
	currentComboTick int
	maxComboTick     int = 600 // 60FPSで10秒
	rainbowIndex     int = 0

	// messages []message
	displayMessages []messages.MessageInterface

	explodes explode.ExplodeCollection

	player characterStatus.CharacterStatusInterface
	enemy  characterStatus.CharacterStatusInterface

	playerDraw characterDraw.CharacterDrawInterface
	enemyDraw  characterDraw.CharacterDrawInterface

	GetExp int
	score int
)

func (m *MineSweeper) GetRouteType() route.RouteType {
	return routeType
}

package minesweeper

import (
	"math"

	"github.com/krile136/mineSweeper/enum/route"
	"github.com/krile136/mineSweeper/scenes/minesweeper/character/characterDraw"
	"github.com/krile136/mineSweeper/scenes/minesweeper/character/characterStatus"
	"github.com/krile136/mineSweeper/scenes/minesweeper/explode"
	"github.com/krile136/mineSweeper/scenes/minesweeper/message/messages"
	"github.com/krile136/mineSweeper/store"
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
	score  int

	isClear      bool = false
	clearTicks   int  = 0
	maxClearTick      = 180

	allOpenTick int = 0
	maxOpenTick int = 180
)

func (m *MineSweeper) GetRouteType() route.RouteType {
	return routeType
}

func (m *MineSweeper) init() {
	// シーン切替.Print時にstoreから行と列のデータを持ってくる
	m.rows = store.Data.MineSweeper.Rows
	m.columns = store.Data.MineSweeper.Columns
	m.bombsNumber = store.Data.MineSweeper.BombsNumber

	GetExp = 0
	score = 0
	isClear = false
	clearTicks = 0
	allOpenTick = 0

	// rowsとcolumnsからフィールドを作成
	m.field = make([][]int, m.rows)
	for i := 0; i < m.columns; i++ {
		m.field[i] = make([]int, m.columns)
	}

	explodes = explode.Create()

	// 爆弾を配置する
	m.placeBombs()

	isBarDisplay = false

	// スクロール可能値を計算する
	// ブロックの大きさは、setWIndowの幅 / Layoutの幅に拡大される
	maxScrollX = math.Max(0, (32*float64(store.Data.MineSweeper.Columns))-float64(store.Data.Layout.OutsideWidth))
	maxScrollY = math.Max(0, (32*float64(store.Data.MineSweeper.Rows)+float64(store.Data.Layout.BattleField))-float64(store.Data.Layout.OutsideHeight))

	// ゲームに関するデータを初期化する
	displayMessages = []messages.MessageInterface{}

	// 各キャラクターの初期ステータスなどが入った配列を初期化する
	initCharacterSlice()

	// playerとenemyに初期値をセットする
	setInitialCharacter()
}

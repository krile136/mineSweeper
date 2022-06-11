package minesweeper

const id string = "mineSweeper"

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

	GetExp          int
	PlayerLv        int
	PlayerHp        float64
	PlayerMaxHp     float64
	PlayerNextExp   int
	PlayerSpeed     int
	PlayerTick      int
	PlayerTurn      bool
	PlayerMove      int
	PlayerDiff      int
	PlayerAttack    int
	PlayerDefense   int
	PlayerActiveBar float64

	EnemyLv        int
	EnemyHp        float64
	EnemyMaxHp     float64
	EnemySpeed     int
	EnemyTick      int
	EnemyTurn      bool
	EnemyMove      int
	EnemyDiff      int
	EnemyAttack    int
	EnemyDefense   int
	EnemyActiveBar float64
)

func (m *MineSweeper) GetId() string {
	return id
}

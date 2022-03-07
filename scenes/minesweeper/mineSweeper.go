package minesweeper

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/krile136/mineSweeper/internal/draw"
	"github.com/krile136/mineSweeper/scenes/scene"
	"github.com/krile136/mineSweeper/store"
)

const id string = "mineSweeper"

type MineSweeper struct {
	rows        int
	colomns     int
	bombsNumber int
	field       [][]int
}

const (
	zero = iota
	one
	twe
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

func (m *MineSweeper) Update() error {
	// シーン切替時にstoreから行と列のデータを持ってくる
	if scene.Is_just_changed {
		m.rows = store.Data.MineSweeper.Rows
		m.colomns = store.Data.MineSweeper.Columns
		m.bombsNumber = store.Data.MineSweeper.BombsNumber

		// rowsとcolumnsからフィールドを作成
		m.field = make([][]int, m.rows)
		for i := 0; i < m.colomns; i++ {
			m.field[i] = make([]int, m.colomns)
		}

		// 爆弾を配置する
		m.placeBombs()

	} else {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			m.placeBombs()
		}
	}
	return nil
}

func (m *MineSweeper) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("random: %d", rand.Intn(100)))
	c := 0.5
	p := 32
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.colomns; j++ {
			x := (float64(j) + 0.5) * float64(p) * c
			y := (float64(i) + 0.5) * float64(p) * c
			switch m.field[i][j] {
			case bomb:
				draw.Draw(screen, "minesweeper", c, x, y, 0, p*3, 0, p, p)
			case close:
				draw.Draw(screen, "minesweeper", c, x, y, 0, 0, 0, p, p)

			}
		}
	}

}

func (m *MineSweeper) GetId() string {
	return id
}

func (m *MineSweeper) placeBombs() error {
	// フィールドを初期化
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.colomns; j++ {
			m.field[i][j] = close
		}
	}

	// 重複のないランダムな数字をrows * columns の中からBombsNumberだけ作り
	// それをfieldに入れて爆弾の位置の初期化を行う
	bombsPosition := make([]int, m.bombsNumber)
	count := 0
	for count < m.bombsNumber {
		position := rand.Intn(m.rows * m.colomns)
		if !inArray(bombsPosition, position) {
			bombsPosition[count] = position
			// fieldにbombを入れる
			// 行は、positionをrowで割った値（15 * 15 のとき、値が23なら 23 / 15 = 1....  なので1行目（0行目があるので実質2行目）
			// 列は、positionをrowsで割ったあまり (15 * 15 のとき、値が23なら 23 % 15 = 8　なので 8列目
			m.field[position/m.rows][position%m.rows] = bomb
			count++
		}
	}

	return nil
}

func inArray(array []int, needle int) bool {
	for _, val := range array {
		if needle == val {
			return true
		}
	}
	return false
}

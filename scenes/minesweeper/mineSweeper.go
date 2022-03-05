package minesweeper

import (
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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

		// フィールドを巡回して、bombsが入っていないところにcloseを入れておく
		for i := 0; i < m.rows; i++ {
			for j := 0; j < m.colomns; j++ {
				if m.field[i][j] != bomb {
					m.field[i][j] = close
				}
			}
		}
	}
	return nil
}

func (m *MineSweeper) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("random: %d", rand.Intn(100)))
}

func (m *MineSweeper) GetId() string {
	return id
}

func inArray(array []int, needle int) bool {
	for _, val := range array {
		if needle == val {
			return true
		}
	}
	return false
}

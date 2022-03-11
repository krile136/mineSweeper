package minesweeper

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/krile136/mineSweeper/scenes/scene"
	"github.com/krile136/mineSweeper/store"
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

	mouse_x, mouse_y := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		var blockWidth int = 32
		y := mouse_y / blockWidth
		x := mouse_x / blockWidth
		position := y*m.rows + x
		if inArray(m.bombsPosition, position) {
			m.field[y][x] = bomb
		} else {
			m.searchAround(x, y)
			for len(nextCheck) > 0 {
				log.Println(fmt.Sprintf("next position: %d", nextCheck[0]))
				search_y := nextCheck[0] / m.rows
				search_x := nextCheck[0] % m.rows

				m.searchAround(search_x, search_y)
			}
		}
	}
	return nil
}

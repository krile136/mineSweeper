package minesweeper

import (
	"fmt"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/krile136/mineSweeper/scenes/scene"
	"github.com/krile136/mineSweeper/store"
)

func (m *MineSweeper) Update() error {
	// シーン切替時にstoreから行と列のデータを持ってくる
	if scene.Is_just_changed {
		m.rows = store.Data.MineSweeper.Rows
		m.columns = store.Data.MineSweeper.Columns
		m.bombsNumber = store.Data.MineSweeper.BombsNumber

		// rowsとcolumnsからフィールドを作成
		m.field = make([][]int, m.rows)
		for i := 0; i < m.columns; i++ {
			m.field[i] = make([]int, m.columns)
		}

		// 爆弾を配置する
		m.placeBombs()

		isBarDisplay = false

		// スクロール可能値を計算する
		// ブロックの大きさは、setWIndowの幅 / Layoutの幅に拡大される
		maxScrollX = math.Max(0, (32*float64(store.Data.MineSweeper.Columns))-float64(store.Data.Layout.OutsideWidth))
		maxScrollY = math.Max(0, (32*float64(store.Data.MineSweeper.Rows))-float64(store.Data.Layout.OutsideHeight))
		log.Println(fmt.Sprintf("maxScrollX: %g", maxScrollX))

	} else {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			m.placeBombs()
		}
	}

	// スクロールしたときの処理
	wheelX, wheelY := ebiten.Wheel()
	scrollCorrectionValue := store.Data.Env.ScrollCorrectionValue
	scrollX = setBetween(-float64(maxScrollX), scrollX+wheelX*float64(scrollCorrectionValue), 0)
	scrollY = setBetween(-float64(maxScrollY), scrollY+wheelY*float64(scrollCorrectionValue), 0)

	// スクロールされている間だけスクロールバーのサイズと位置を計算する
	if wheelX != 0 || wheelY != 0 {
		isBarDisplay = true
		BarDisplayFrame = 30
		barLengthY = math.Max(0.5, float64(store.Data.Layout.OutsideHeight)/(float64(store.Data.Layout.OutsideHeight)+maxScrollY)) * float64(store.Data.Layout.OutsideHeight)
		barSlideY = ((float64(store.Data.Layout.OutsideHeight)-barLengthY)/maxScrollY)*math.Abs(scrollY) + barLengthY/2
		barLengthX = math.Max(0.5, float64(store.Data.Layout.OutsideWidth)/(float64(store.Data.Layout.OutsideWidth)+maxScrollX)) * float64(store.Data.Layout.OutsideWidth)
		barSlideX = ((float64(store.Data.Layout.OutsideWidth)-barLengthX)/maxScrollX)*math.Abs(scrollX) + barLengthX/2
	} else {
		if isBarDisplay {
			BarDisplayFrame -= 1
			if BarDisplayFrame <= 0 {
				isBarDisplay = false
			}
		}
	}

	// マウスの座標をスクロールの分だけ補正する
	mouse_x, mouse_y := ebiten.CursorPosition()
	y := (mouse_y - int(scrollY)) / blockWidth
	x := (mouse_x - int(scrollX)) / blockWidth

	// 左クリックしたときの処理
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		// クリックしたマスがフラグが立っていれば何もしない
		if m.field[y][x] != flag {
			position := y*m.rows + x
			if inArray(m.bombsPosition, position) {
				m.field[y][x] = bomb
			} else {
				m.searchAround(x, y)
				for len(nextCheck) > 0 {
					search_y := nextCheck[0] / m.rows
					search_x := nextCheck[0] % m.rows

					m.searchAround(search_x, search_y)
				}
			}
		}
	}

	// 右クリックしたときの処理
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonRight) {
		switch m.field[y][x] {
		case close:
			m.placeFlag(x, y)
		case flag:
			m.field[y][x] = close
		case one, two, three, four, five, six, seven, eight:
			m.searchAroundOnNumberField(x, y)
			for len(nextCheck) > 0 {
				search_y := nextCheck[0] / m.rows
				search_x := nextCheck[0] % m.rows

				m.searchAround(search_x, search_y)
			}
		default:
			// 何もしない
		}
	}
	return nil
}

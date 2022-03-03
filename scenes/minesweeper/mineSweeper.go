package minesweeper

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/krile136/mineSweeper/scenes/scene"
	"github.com/krile136/mineSweeper/store"
)

const id string = "mineSweeper"

type MineSweeper struct {
	rows    int
	colomns int
}

func (m *MineSweeper) Update() error {
	// シーン切替時にstoreから行と列のデータを持ってくる
	if scene.Is_just_changed {
		m.rows = store.Data.MineField.Rows
		m.colomns = store.Data.MineField.Columns
	}
	return nil
}

func (m *MineSweeper) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("rows: %d", m.rows))
}

func (m *MineSweeper) GetId() string {
	return id
}

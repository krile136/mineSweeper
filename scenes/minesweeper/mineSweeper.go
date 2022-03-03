package minesweeper

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const id string = "mineSweeper"

type MineSweeper struct {
}

func (m *MineSweeper) Update() error {
	return nil
}

func (m *MineSweeper) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "mine")
}

func (m *MineSweeper) GetId() string {
	return id
}

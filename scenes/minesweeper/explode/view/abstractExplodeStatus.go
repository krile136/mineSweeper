package view

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/krile136/mineSweeper/internal/draw"
)

type abstractExplodeView struct {
	x        float64
	y        float64
	tick     int
	sheets   int
	interval int
	filename string
	width    int
	pictureY int
}

func (o Orange) Draw(screen *ebiten.Image) {
	if o.tick > 0 {
		pictureX := o.CalcSheetsNumber()
		draw.Draw(screen, o.filename, 1, 1, o.x, o.y, 0, pictureX, o.pictureY, 32, 32)
	}
}

func (a *abstractExplodeView) CalcSheetsNumber() int {
	i := (a.tick / a.interval) % a.sheets
	return a.width * i
}

func (a *abstractExplodeView) IsFinish() bool {
	return a.tick > a.interval*a.sheets
}

func (a *abstractExplodeView) makeAbstractExplodeView(
	x float64,
	y float64,
	pictureY int,
) (ae *abstractExplodeView) {
	ae = &abstractExplodeView{
		x:        x,
		y:        y,
		tick:     0,
		sheets:   6,
		interval: 5,
		filename: "bomb",
		width:    32,
		pictureY: pictureY,
	}
	return
}

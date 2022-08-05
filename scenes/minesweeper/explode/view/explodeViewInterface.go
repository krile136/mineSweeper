package view

import "github.com/hajimehoshi/ebiten/v2"

type ExplodeViewInterface interface {
	New(x, y float64) (new ExplodeViewInterface)
	Update() ExplodeViewInterface
	Draw(screen *ebiten.Image)
	CalcSheetsNumber() int
	IsFinish() bool

	makeAbstractExplodeView(
		x float64,
		y float64,
		pictureY int,
	) (ae *abstractExplodeView)
}

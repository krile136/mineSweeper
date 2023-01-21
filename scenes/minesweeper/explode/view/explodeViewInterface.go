package view

import "github.com/hajimehoshi/ebiten/v2"

type ExplodeViewInterface interface {
	New(x, y float64, tick int, delay int) (new ExplodeViewInterface)
	Update() ExplodeViewInterface
	Draw(screen *ebiten.Image)
	CalcSheetsNumber() int
	IsNotFinish() bool

	makeAbstractExplodeView(
		x float64,
		y float64,
		pictureY int,
		tick int,
		delay int,
	) (ae *abstractExplodeView)
}

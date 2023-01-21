package view

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/krile136/mineSweeper/internal/draw"
)

type abstractExplodeView struct {
	x        float64
	y        float64
	tick     int
	delay    int
	sheets   int
	interval int
	filename string
	width    int
	pictureY int
}

// 画像を表示させる
func (o Orange) Draw(screen *ebiten.Image) {
	if o.tick > 0 {
		pictureX := o.CalcSheetsNumber()
		draw.Draw(screen, o.filename, 1, 1, o.x, o.y, 0, pictureX, o.pictureY, 32, 32)
	}
}

// 画像ファイルの何番目の画像かを計算する
func (a *abstractExplodeView) CalcSheetsNumber() int {
	i := (a.tick / a.interval) % a.sheets
	return a.width * i
}

// 一通りの描画が終わっていないかチェックする
func (a *abstractExplodeView) IsNotFinish() bool {
	return a.tick < a.interval*a.sheets
}

func (a *abstractExplodeView) makeAbstractExplodeView(
	x float64,
	y float64,
	pictureY int,
	tick int,
	delay int,
) (ae *abstractExplodeView) {
	ae = &abstractExplodeView{
		x:        x,
		y:        y,
		tick:     tick,
		delay:    delay,
		sheets:   6,
		interval: 5,
		filename: "bomb",
		width:    32,
		pictureY: pictureY,
	}
	return
}

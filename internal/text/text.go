package text

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/krile136/mineSweeper/store"
)

func DrawText(target *ebiten.Image, str string, x, y int, clr color.Color) {
	text.Draw(target, str, store.Data.Env.Font, x, y, clr)
}

package text

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/krile136/mineSweeper/store"
)

func DrawText(target *ebiten.Image, str string, x, y int, size string, clr color.Color) {
	switch size {
	case "L":
		text.Draw(target, str, store.Data.Font.Large, x, y, clr)
	case "M":
		text.Draw(target, str, store.Data.Font.Middle, x, y, clr)
	case "S":
		text.Draw(target, str, store.Data.Font.Small, x, y, clr)
	}
}

func Length(message string, size string) int {
	font := store.Data.Font.Middle
	switch size {
	case "L":
		font = store.Data.Font.Large
	case "S":
		font = store.Data.Font.Small
	}
	len := text.BoundString(font, message)
	return len.Max.X
}

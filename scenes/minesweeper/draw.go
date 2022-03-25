package minesweeper

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/krile136/mineSweeper/internal/draw"
)

func (m *MineSweeper) Draw(screen *ebiten.Image) {
	c := 1.0
	p := 32
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.colomns; j++ {
			x := (float64(j)+0.5)*float64(p)*c + scrollX
			y := (float64(i)+0.5)*float64(p)*c + scrollY
			switch m.field[i][j] {
			case close:
				draw.Draw(screen, "minesweeper", c, x, y, 0, 0, 0, p, p)
			case zero:
				draw.Draw(screen, "minesweeper", c, x, y, 0, p, 0, p, p)
			case flag:
				draw.Draw(screen, "minesweeper", c, x, y, 0, p*2, 0, p, p)
			case bomb:
				draw.Draw(screen, "minesweeper", c, x, y, 0, p*3, 0, p, p)
			case one:
				draw.Draw(screen, "minesweeper", c, x, y, 0, 0, p, p, p)
			case two:
				draw.Draw(screen, "minesweeper", c, x, y, 0, p, p, p, p)
			case three:
				draw.Draw(screen, "minesweeper", c, x, y, 0, p*2, p, p, p)
			case four:
				draw.Draw(screen, "minesweeper", c, x, y, 0, p*3, p, p, p)
			case five:
				draw.Draw(screen, "minesweeper", c, x, y, 0, p*4, p, p, p)
			case six:
				draw.Draw(screen, "minesweeper", c, x, y, 0, p*5, p, p, p)
			case seven:
				draw.Draw(screen, "minesweeper", c, x, y, 0, p*6, p, p, p)
			case eight:
				draw.Draw(screen, "minesweeper", c, x, y, 0, p*7, p, p, p)
			}
		}
	}

}

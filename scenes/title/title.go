package title

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/krile136/mineSweeper/scenes/scene"
)

const id string = "title"

type Title struct {
}

func (t *Title) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		scene.Id = "mineSweeper"
	}
	return nil
}

func (t *Title) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "title")
}

func (t *Title) GetId() string {
	return id
}

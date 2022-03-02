package title

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const id string = "title"

type Title struct {
}

func (t *Title) Update() error {
	// if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
	// 	gameInterface.Scene_id = "ending"
	// }
	return nil
}

func (t *Title) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "title")
}

func (t *Title) GetId() string {
	return id
}

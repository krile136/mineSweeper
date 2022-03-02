package game

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/krile136/mineSweeper/scenes/scene"
	"github.com/krile136/mineSweeper/scenes/title"
)

type Game struct {
}

func NewGame() (*Game, error) {
	scene.Display = &title.Title{}
	game := &Game{}

	return game, nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}

func (g *Game) Update() error {
	if scene.Display.GetId() != scene.Id {
		scene.Display = route[scene.Id]
		scene.Is_just_changed = true
	}
	scene.Display.Update()

	if scene.Is_just_changed {
		scene.Is_just_changed = false
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	scene.Display.Draw(screen)
}

package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/krile136/mineSweeper/game"
)

func main() {

	game, err := game.NewGame()
	if err != nil {
		panic(err)
	}

	ebiten.SetWindowSize(640, 640)

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joho/godotenv"

	"github.com/krile136/mineSweeper/game"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	game, err := game.NewGame()
	if err != nil {
		panic(err)
	}

	ebiten.SetWindowSize(320, 320)

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

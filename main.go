package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/krile136/mineSweeper/game"
)

func main() {
	game, err := game.NewGame()
	if err != nil {
		log.Printf("failed to create game:  %s", err.Error())
		panic(err)
	}

	// store.Data.Layoutの横幅と縦幅を一緒にしておくこと
	ebiten.SetWindowSize(640, 480) 

	if err := ebiten.RunGame(game); err != nil {
		log.Printf("failed to run game: %s", err.Error())
		panic(err)
	}
}

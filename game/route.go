package game

import (
	"github.com/krile136/mineSweeper/scenes/minesweeper"
	"github.com/krile136/mineSweeper/scenes/scene"
	"github.com/krile136/mineSweeper/scenes/title"
)

var route = map[string]scene.Scene{
	"title":       &title.Title{},
	"mineSweeper": &minesweeper.MineSweeper{},
}

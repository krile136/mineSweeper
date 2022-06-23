package game

import (
	"github.com/krile136/mineSweeper/scenes/minesweeper"
	"github.com/krile136/mineSweeper/scenes/scene"
	"github.com/krile136/mineSweeper/scenes/title"
	"github.com/krile136/mineSweeper/types/route"
)

var routeMap = map[route.RouteType]scene.Scene{
	route.Title:       &title.Title{},
	route.MineSweeper: &minesweeper.MineSweeper{},
}

package title

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/krile136/mineSweeper/enum/route"
	"github.com/krile136/mineSweeper/internal/text"
	"github.com/krile136/mineSweeper/scenes/scene"
	"github.com/krile136/mineSweeper/store"
)

const routeType route.RouteType = route.Title

type Title struct {
}


func (t *Title) Update() error {

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		var mx, my int = ebiten.CursorPosition()
		var centerX float64 = float64(store.Data.Layout.OutsideWidth) / 2
		var centerY float64 = float64(store.Data.Layout.OutsideHeight) / 2

		var start_length = text.Length("Start Game", "M")
		var margin int = 3
		if inBetween(int(centerX)-start_length/2-margin, mx, int(centerX)+start_length/2+margin) {
			if inBetween(int(centerY)+87, my, int(centerY)+100+margin) {
				scene.RouteType = route.MineSweeper
			}
		}
		var ranking_length = text.Length("Ranking", "M")
		if inBetween(int(centerX)-ranking_length/2-margin, mx, int(centerX)+ranking_length/2+margin) {
			if inBetween(int(centerY)+127, my, int(centerY)+140+margin) {
				scene.RouteType = route.Ranking
			}

		}

		// var gameover_length = text.Length("GameOver", "M")
		// if inBetween(int(centerX)-gameover_length/2-margin, mx, int(centerX)+gameover_length/2+margin) {
		// 	if inBetween(int(centerY)+167, my, int(centerY)+180+margin) {
		// 		scene.RouteType = route.GameOver
		// 	}
		// }

	}
	return nil
}

func (t *Title) Draw(screen *ebiten.Image) {

	var centerX float64 = float64(store.Data.Layout.OutsideWidth) / 2
	var centerY float64 = float64(store.Data.Layout.OutsideHeight) / 2
	text.DrawTextAtCenter(screen, "B A T T L E", int(centerX), int(centerY-15), "L", store.Data.Color.White)
	text.DrawTextAtCenter(screen, "MINE SWEEPER", int(centerX), int(centerY+15), "L", store.Data.Color.White)

	text.DrawTextAtCenter(screen, "Start Game", int(centerX), int(centerY+100), "M", store.Data.Color.White)

	text.DrawTextAtCenter(screen, "Ranking", int(centerX), int(centerY+140), "M", store.Data.Color.White)

	// text.DrawTextAtCenter(screen, "Gameover", int(centerX), int(centerY+180), "M", store.Data.Color.White)

}

func (t *Title) GetRouteType() route.RouteType {
	return routeType
}

// int型の値が最小と最大の間にあるかチェックする
func inBetween(min, val, max int) bool {
	if (val >= min) && (val <= max) {
		return true
	} else {
		return false
	}
}

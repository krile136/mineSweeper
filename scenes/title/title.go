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
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		scene.RouteType = route.MineSweeper
	}
	return nil
}

func (t *Title) Draw(screen *ebiten.Image) {

	var centerX float64 = float64(store.Data.Layout.OutsideWidth) / 2
	var centerY float64 = float64(store.Data.Layout.OutsideHeight) / 2
	text.DrawTextAtCenter(screen, "B A T T L E", int(centerX), int(centerY-15), "L", store.Data.Color.White)
	text.DrawTextAtCenter(screen, "MINE SWEEPER", int(centerX), int(centerY+15), "L", store.Data.Color.White)

	text.DrawTextAtCenter(screen, "Push Enter", int(centerX), int(centerY+100), "M", store.Data.Color.White)
}

func (t *Title) GetRouteType() route.RouteType {
	return routeType
}

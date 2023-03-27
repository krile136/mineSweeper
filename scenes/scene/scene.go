package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/krile136/mineSweeper/enum/route"
)

var (
	Is_just_changed bool = true 
	Display         Scene
	Next            Scene
	RouteType       route.RouteType = route.Login
)

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
	GetRouteType() route.RouteType
}

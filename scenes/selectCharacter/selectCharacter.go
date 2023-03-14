package selectcharacter

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/krile136/mineSweeper/enum/route"
)

const routeType route.RouteType = route.SelectCharacter

type SelectCharacter struct {
}

func (s *SelectCharacter) Update() error {

	return nil
}

func (s *SelectCharacter) Draw(screen *ebiten.Image) {

}

func (s *SelectCharacter) GetRouteType() route.RouteType {
	return routeType
}

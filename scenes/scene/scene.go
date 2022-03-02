package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Is_just_changed bool = false
	Display         Scene
	Id              string = "title"
)

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
	GetId() string
}

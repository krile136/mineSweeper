package explode

import (
	"github.com/krile136/mineSweeper/scenes/minesweeper/explode/view"
)

var ExplodeMap = map[ExplodeType]view.ExplodeViewInterface{
	Orange: view.Orange{},
}

func Create() (slice []view.ExplodeViewInterface) {
	explodeStruct := ExplodeMap[Orange]
	slice = append(slice, explodeStruct.New(100, 100), explodeStruct.New(150, 100))
	return
}

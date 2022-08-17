package explode

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/krile136/mineSweeper/scenes/minesweeper/explode/view"
)

var ExplodeMap = map[ExplodeType]view.ExplodeViewInterface{
	Orange: view.Orange{},
}

func Create() (collection ExplodeCollection) {
	slice := []view.ExplodeViewInterface{}
	explodeStruct := ExplodeCollection{slice}
	return explodeStruct
}

type ExplodeCollection struct {
	collection []view.ExplodeViewInterface
}

func (e ExplodeCollection) Add(eType ExplodeType, x, y float64, tick int) ExplodeCollection {
	newSlice := e.collection
	explode := ExplodeMap[eType].New(x, y, tick)
	newSlice = append(newSlice, explode)
	new := ExplodeCollection{newSlice}
	return new
}

func (e ExplodeCollection) Draw(screen *ebiten.Image) {
	for _, v := range e.collection {
		v.Draw(screen)
	}
}

func (e ExplodeCollection) Update() (ExplodeCollection, int) {
	new := e
	tmpExplodes := []view.ExplodeViewInterface{}
	finish := 0
	for _, explode := range e.collection {
		newExplode := explode.Update()
		if newExplode.IsNotFinish() {
			tmpExplodes = append(tmpExplodes, newExplode)
		} else {
			finish += 1
		}
	}
	new.collection = tmpExplodes
	return new, finish
}

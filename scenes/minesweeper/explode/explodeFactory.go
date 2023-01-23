package explode

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/krile136/mineSweeper/scenes/minesweeper/explode/view"
)

// 爆発の元を定義する
var ExplodeMap = map[ExplodeType]view.ExplodeViewInterface{
	Orange: view.Orange{},
}

// 爆発一覧を作成する
func Create() (collection ExplodeCollection) {
	slice := []view.ExplodeViewInterface{}
	explodeStruct := ExplodeCollection{slice}
	return explodeStruct
}

// 爆発のコレクション
type ExplodeCollection struct {
	collection []view.ExplodeViewInterface
}

// 爆発のコレクションに追加
func (e ExplodeCollection) Add(eType ExplodeType, x, y float64, tick int, delay int) ExplodeCollection {
	newSlice := e.collection
	explode := ExplodeMap[eType].New(x, y, tick, delay)
	newSlice = append(newSlice, explode)
	new := ExplodeCollection{newSlice}
	return new
}

// 爆発のコレクションを元に描画する
func (e ExplodeCollection) Draw(screen *ebiten.Image) {
	for _, v := range e.collection {
		v.Draw(screen)
	}
}

// 爆発のコレクションの各要素を更新する
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

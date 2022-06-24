package messages

import (
	"image/color"

	"github.com/krile136/mineSweeper/store"
)

// 敵が受けたダメージのメッセージ
type LevelUp struct {
	value string
	x     float64
	y     float64
	tick  int
	crl   color.Color
}

/*
-------------------- Public Method --------------------
*/

func (l LevelUp) New(value string) MessageInterface {
	new := LevelUp{}
	new.value = value
	new.x, new.y = l.defaultPosition()
	new.tick = 0
	new.crl = store.Data.Color.Orange
	return new

}
func (l LevelUp) String() string {
	return "Level UP !!"
}

func (l LevelUp) Update() MessageInterface {
	mx, my := l.moveValue(l.tick)

	new := l
	new.x = new.x + mx
	new.y = new.y + my
	new.tick = new.tick + 1
	return new
}

func (l LevelUp) IsExist() bool {
	return l.tick <= l.existTick()
}

func (l LevelUp) GetFieldForDraw() (value string, x, y int, crl color.Color) {
	value = l.value
	x = int(l.x)
	y = int(l.y)
	crl = l.crl
	return
}

/*
-------------------- Private Method --------------------
*/

func (l LevelUp) defaultPosition() (x, y float64) {
	x = 100
	y = 0
	return
}

func (l LevelUp) existTick() int {
	return 70
}

func (l LevelUp) moveValue(tick int) (x, y float64) {
	g := float64(-0.1)
	v := g * float64(l.existTick()) * -0.5
	diffY := v + g*float64(tick)

	x = 0
	y = diffY
	return
}

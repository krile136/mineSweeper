package messages

import (
	"image/color"

	"github.com/krile136/mineSweeper/store"
)

// 敵が受けたダメージのメッセージ
type LevelUp struct {
	*abstractMessage
}

/*
-------------------- Public Method --------------------
*/

func (l LevelUp) New(value string) MessageInterface {
	x, y, existTick, crl := l.defaultField()
	am := l.makeAbstractMessage(value, x, y, crl, existTick)
	new := LevelUp{am}
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

/*
-------------------- Private Method --------------------
*/

func (l LevelUp) defaultField() (x, y float64, existTick int, crl color.Color) {
	x = 100
	y = 0
	existTick = 70
	crl = store.Data.Color.Orange
	return
}

func (l LevelUp) moveValue(tick int) (x, y float64) {
	g := float64(-0.1)
	v := g * float64(l.existTick) * -0.5
	diffY := v + g*float64(tick)

	x = 0
	y = diffY
	return
}

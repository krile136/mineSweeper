package messages

import (
	"image/color"

	"github.com/krile136/mineSweeper/store"
)

// 敵が受けたダメージのメッセージ
type EnemyDamage struct {
	*abstractMessage
}

/*
-------------------- Public Method --------------------
*/

func (e EnemyDamage) New(value string) MessageInterface {
	x, y, existTick, crl := e.defaultField()
	am := e.makeAbstractMessage(value, x, y, crl, existTick)
	new := EnemyDamage{am}
	return new
}

func (e EnemyDamage) String() string {
	return "EnemyDamage"
}

func (e EnemyDamage) Update() MessageInterface {
	mx, my := e.moveValue(e.tick)

	new := e
	new.x = new.x + mx
	new.y = new.y + my
	new.tick = new.tick + 1
	return new
}

/*
-------------------- Private Method --------------------
*/

func (e EnemyDamage) defaultField() (x, y float64, existTick int, crl color.Color) {
	x = 210
	y = 55
	existTick = 40
	crl = store.Data.Color.Red
	return
}

func (e EnemyDamage) moveValue(tick int) (x, y float64) {
	g := float64(0.1)
	v := g * float64(e.existTick) * -0.5
	diffY := v + g*float64(tick)

	x = 1
	y = diffY
	return
}

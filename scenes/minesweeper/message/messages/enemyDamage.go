package messages

import (
	"image/color"

	"github.com/krile136/mineSweeper/store"
)

// 敵が受けたダメージのメッセージ
type EnemyDamage struct {
	value string
	x     float64
	y     float64
	tick  int
	crl   color.Color
}

/*
-------------------- Public Method --------------------
*/

func (e EnemyDamage) New(value string) MessageInterface {
	new := EnemyDamage{}
	new.value = value
	new.x, new.y = e.defaultPosition()
	new.tick = 0
	new.crl = store.Data.Color.Red
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

func (e EnemyDamage) IsExist() bool {
	return e.tick <= e.existTick()
}

func (e EnemyDamage) GetFieldForDraw() (value string, x, y int, crl color.Color) {
	value = e.value
	x = int(e.x)
	y = int(e.y)
	crl = e.crl
	return
}

/*
-------------------- Private Method --------------------
*/

func (e EnemyDamage) defaultPosition() (x, y float64) {
	x = 210
	y = 55
	return
}

func (e EnemyDamage) existTick() int {
	return 40
}

func (e EnemyDamage) moveValue(tick int) (x, y float64) {
	g := float64(0.1)
	v := g * float64(e.existTick()) * -0.5
	diffY := v + g*float64(tick)

	x = 1
	y = diffY
	return
}

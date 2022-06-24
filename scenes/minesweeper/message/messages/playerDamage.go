package messages

import (
	"image/color"

	"github.com/krile136/mineSweeper/store"
)

// プレイヤーが受けたダメージのメッセージ
type PlayerDamage struct {
	value string
	x     float64
	y     float64
	tick  int
	crl   color.Color
}

/*
-------------------- Public Method --------------------
*/

func (p PlayerDamage) New(value string) MessageInterface {
	new := PlayerDamage{}
	new.value = value
	new.x, new.y = p.defaultPosition()
	new.tick = 0
	new.crl = store.Data.Color.Red
	return new

}
func (p PlayerDamage) String() string {
	return "PlayerDamage"
}

func (p PlayerDamage) Update() MessageInterface {
	mx, my := p.moveValue(p.tick)

	new := p
	new.x = new.x + mx
	new.y = new.y + my
	new.tick = new.tick + 1
	return new
}

func (p PlayerDamage) IsExist() bool {
	return p.tick <= p.existTick()
}

func (p PlayerDamage) GetFieldForDraw() (value string, x, y int, crl color.Color) {
	value = p.value
	x = int(p.x)
	y = int(p.y)
	crl = p.crl
	return
}

/*
-------------------- Private Method --------------------
*/

func (p PlayerDamage) defaultPosition() (x, y float64) {
	x = 100
	y = 55
	return
}

func (p PlayerDamage) existTick() int {
	return 40
}

func (p PlayerDamage) moveValue(tick int) (x, y float64) {
	g := float64(0.1)
	v := g * float64(p.existTick()) * -0.5
	diffY := v + g*float64(tick)

	x = -1
	y = diffY
	return
}

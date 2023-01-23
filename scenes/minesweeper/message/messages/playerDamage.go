package messages

import (
	"image/color"

	"github.com/krile136/mineSweeper/store"
)

// プレイヤーが受けたダメージのメッセージ
type PlayerDamage struct {
	*abstractMessage
}

/*
-------------------- Public Method --------------------
*/

func (p PlayerDamage) New(value string) MessageInterface {
	x, y, existTick, crl := p.defaultField()
	am := p.makeAbstractMessage(value, x, y, crl, existTick)
	new := PlayerDamage{am}

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

/*
-------------------- Private Method --------------------
*/

func (p PlayerDamage) defaultField() (x, y float64, existTick int, crl color.Color) {
	x = float64(store.Data.Layout.OutsideWidth)/2 - 75
	y = 60
	existTick = 40
	crl = store.Data.Color.Red
	return
}

func (p PlayerDamage) moveValue(tick int) (x, y float64) {
	g := float64(0.1)
	v := g * float64(p.existTick) * -0.5
	diffY := v + g*float64(tick)

	x = -1
	y = diffY
	return
}

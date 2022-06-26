package messages

import (
	"image/color"
)

type abstractMessage struct {
	MessageInterface
	value string
	x     float64
	y     float64
	tick  int
	crl   color.Color

	existTick int
}

/*
-------------------- Public Method --------------------
*/

func (a *abstractMessage) GetFieldForDraw() (value string, x, y int, crl color.Color) {
	value = a.value
	x = int(a.x)
	y = int(a.y)
	crl = a.crl
	return
}

func (a *abstractMessage) IsExist() bool {
	return a.tick <= a.existTick
}

/*
-------------------- Private Method --------------------
*/

func (a *abstractMessage) makeAbstractMessage(value string, x, y float64, crl color.Color, existTick int) (am *abstractMessage) {
	am = &abstractMessage{
		value:     value,
		x:         x,
		y:         y,
		tick:      0,
		crl:       crl,
		existTick: existTick,
	}
	return

}

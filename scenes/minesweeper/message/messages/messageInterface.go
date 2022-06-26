package messages

import "image/color"

type MessageInterface interface {
	// public method
	New(string) MessageInterface
	String() string
	Update() MessageInterface
	IsExist() bool
	GetFieldForDraw() (value string, x, y int, crl color.Color)

	// private method
	makeAbstractMessage(value string, x, y float64, crl color.Color, existTick int) *abstractMessage
	defaultField() (x, y float64, existTick int, crl color.Color)
	moveValue(tick int) (x, y float64)
}

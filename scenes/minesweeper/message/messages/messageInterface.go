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
	defaultPosition() (x, y float64)
	existTick() int
	moveValue(tick int) (x, y float64)
}

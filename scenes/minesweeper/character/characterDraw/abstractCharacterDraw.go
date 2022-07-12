package characterDraw

import "math"

type abstractCharacterDraw struct {
	CharacterDrawInterface
	blinkingTick int
	positionX    float64
	positionY    float64
	direction    float64
	difference   float64
}

const topDifference float64 = 50
const DeadBlinkingMaxTick int = 100

func (a *abstractCharacterDraw) CanExecuteInvertAtTop() bool {
	return math.Abs(a.difference) >= topDifference
}

func (a *abstractCharacterDraw) IsFinishDeadBlinking() bool {
	return a.blinkingTick >= DeadBlinkingMaxTick
}

func (a *abstractCharacterDraw) addDirectionToDifference() {
	a.difference += a.direction * 5
}

func (a *abstractCharacterDraw) invertDirection() {
	a.direction = a.direction * -1
}

func (a *abstractCharacterDraw) PositionX() float64 {
	return a.positionX
}

func (a *abstractCharacterDraw) PositionY() float64 {
	return a.positionY
}

func (a *abstractCharacterDraw) Difference() float64 {
	return a.difference
}
func (a *abstractCharacterDraw) CurrentPosition() float64 {
	return a.positionX + a.difference
}

func (a *abstractCharacterDraw) makeAbstractCharacterDraw(positionX, positionY, direction, difference float64) (acd *abstractCharacterDraw) {
	acd = &abstractCharacterDraw{
		blinkingTick: 0,
		positionX:    positionX,
		positionY:    positionY,
		direction:    direction,
		difference:   difference,
	}
	return
}

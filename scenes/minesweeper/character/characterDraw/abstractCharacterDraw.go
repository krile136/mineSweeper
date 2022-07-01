package characterDraw

import "math"

type abstractCharacterDraw struct {
	CharacterDrawInterface
	blinkingTick int
	direction    float64
	difference   float64
}

const topDifference float64 = 50

func (a *abstractCharacterDraw) CanExecuteInvertAtTop() bool {
	return math.Abs(a.difference) >= topDifference
}

func (a *abstractCharacterDraw) CanExecuteInvertAtBase() bool {
	// enemyのときはマイナスがプラスになったことを判定
	// playerのときは directionが - 、differenceが +
	// enemyのときは directionが+ , differenceが-
	// かけると常に負になるので、これが 0より大きくなったときが反転タイミング
	return a.difference*a.direction >= 0
}

func (a *abstractCharacterDraw) addDirectionToDifference() {
	a.difference += a.direction * 5
}

func (a *abstractCharacterDraw) invertDirection() {
	a.direction = a.direction * -1
}

func (a *abstractCharacterDraw) makeAbstractCharacterDraw(direction, difference float64) (acd *abstractCharacterDraw) {
	acd = &abstractCharacterDraw{
		blinkingTick: 0,
		direction:    direction,
		difference:   difference,
	}
	return
}

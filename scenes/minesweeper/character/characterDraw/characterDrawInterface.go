package characterDraw

type CharacterDrawInterface interface {
	// public method
	New() CharacterDrawInterface
	ExecuteMoving() CharacterDrawInterface
	UpdateBlinking() CharacterDrawInterface
	InvertDirection() CharacterDrawInterface
	CanExecuteInvertAtTop() bool
	CanExecuteInvertAtBase() bool
	CanFinishAppearing() bool
	IsReturningToBase() bool
	IsShowBlinking() bool
	FinishTurn() CharacterDrawInterface
	IsFinishDeadBlinking() bool
	IsFinishAppearing() bool
	IsBlinking() bool
	SetInitialDraw() CharacterDrawInterface
	PositionX() float64
	PositionY() float64
	Difference() float64
	CurrentPosition() float64

	// private method
	makeAbstractCharacterDraw(positionX, positionY, direction, difference float64) (acd *abstractCharacterDraw)
}

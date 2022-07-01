package characterDraw

type CharacterDrawInterface interface {
	// public method
	New() CharacterDrawInterface
	ExecuteMoving() CharacterDrawInterface
	InvertDirection() CharacterDrawInterface
	CanExecuteInvertAtTop() bool
	CanExecuteInvertAtBase() bool
	FinishTurn() CharacterDrawInterface

	// private method
	makeAbstractCharacterDraw(direction, difference float64) (acd *abstractCharacterDraw)
}

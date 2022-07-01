package characterStatus

type CharacterStatusInterface interface {
	// public method
	New(lv int) (new CharacterStatusInterface)
	// String() string
	AttackTo(CharacterStatusInterface) CharacterStatusInterface
	LevelUp(exp int) (bool, CharacterStatusInterface)
	Update() CharacterStatusInterface
	InvertTurn() CharacterStatusInterface
	FinishTurn() CharacterStatusInterface
	Lv() int
	Hp() int
	MaxHp() int
	NextExp() int
	Turn() bool
	GetDamageAmount(CharacterStatusInterface) float64

	// private method
	// initialStatus()
	// growthRate()
	calcDamage(currentAttack, targetDefense float64) (damage float64)
	calcNextExp(lv int) (next int)
	getDefense() float64
	reduceHp(damage float64) CharacterStatusInterface
}

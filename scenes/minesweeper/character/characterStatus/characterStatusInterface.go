package characterStatus

type CharacterStatusInterface interface {
	// public method
	New(lv int) (new CharacterStatusInterface)
	// String() string
	AttackTo(CharacterStatusInterface) CharacterStatusInterface
	LevelUp(exp int) (bool, CharacterStatusInterface)
	Update() CharacterStatusInterface
	InvertTurn() CharacterStatusInterface
	SetTurn(boolean bool) CharacterStatusInterface
	FinishTurn() CharacterStatusInterface
	Lv() int
	Hp() int
	MaxHp() int
	NextExp() int
	Turn() bool
	Name() string
	ActiveBar() float64
	GetDamageAmount(CharacterStatusInterface) float64
	StopTimer() bool
	Dead() bool
	Appearing() bool
	AddCondition(cond condition) CharacterStatusInterface
	CanTurnOn() bool
	ResetCondition() CharacterStatusInterface

	SetInitialStatus() CharacterStatusInterface

	// private method
	// initialStatus()
	// growthRate()
	calcDamage(currentAttack, targetDefense float64) (damage float64)
	calcNextExp(lv int) (next int)
	getDefense() float64
	reduceHp(damage float64) CharacterStatusInterface
}

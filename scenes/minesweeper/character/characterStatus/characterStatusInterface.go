package characterStatus

type CharacterStatusInterface interface {
	// public method
	New(lv int) (new CharacterStatusInterface)
	// String() string
	AttackTo(CharacterStatusInterface) CharacterStatusInterface
	LevelUp(exp int) (bool, CharacterStatusInterface)
	Update() CharacterStatusInterface
	Lv() int
	Hp() int
	MaxHp() int
	NextExp() int
	Turn() bool

	// private method
	// initialStatus()
	// growthRate()
	calcNextExp(lv int) (next int)
	getDefense() float64
	reduceHp(damage float64) CharacterStatusInterface
}

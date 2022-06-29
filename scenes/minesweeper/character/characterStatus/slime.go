package characterStatus

import "math"

type Slime struct {
	*abstractCharacterStatus
}

// コンストラクタ
func (s Slime) New(lv int) (new CharacterStatusInterface) {
	ac := s.makeAbstractCharacterStatus(s.defaultField())
	new = Slime{ac}
	return
}

// 攻撃をする
func (s Slime) AttackTo(target CharacterStatusInterface) CharacterStatusInterface {
	damage := s.calcDamage(s.attack, target.getDefense())

	newTarget := target.reduceHp(damage)
	return newTarget
}

// レベルアップを実施する
func (s Slime) LevelUp(exp int) (bool, CharacterStatusInterface) {
	new := s
	isLevelUp, newLv, newHp, newMaxHp, newAttack, newDefense, newNextExp := new.addExp(exp)
	new.setBattleStatus(
		newLv,
		newHp,
		newMaxHp,
		newAttack,
		newDefense,
		newNextExp,
	)

	return isLevelUp, new
}

// 毎フレーム更新
func (s Slime) Update() CharacterStatusInterface {
	new := s
	new.tick += 1
	s.activeBar = float64(new.tick) / float64(new.speed)
	return new
}

// 守備力を取得する
func (s Slime) getDefense() float64 {
	return s.defense
}

// HPを減少させる
func (s Slime) reduceHp(damage float64) CharacterStatusInterface {
	new := s

	new.hp = math.Max(0, s.hp-damage)
	return new
}

// バトルに関するステータスをセットする
func (s Slime) setBattleStatus(
	lv int,
	hp, maxHp, attack, defense float64,
	nextExp int,
) CharacterStatusInterface {
	new := s
	new.lv = lv
	new.hp = hp
	new.maxHp = maxHp
	new.attack = attack
	new.defense = defense
	new.nextExp = nextExp

	return new
}

// デフォルトのフィールド値を取得する
func (s Slime) defaultField() (
	name string,
	lv int,
	hp, maxHp float64,
	nextExp, speed, tick int,
	attack, defense, activeBar float64,
	baseNextExp int,
	hpRate, attackRate, defenseRate float64,
) {
	name = "Player"
	lv = 1
	hp = 100
	maxHp = 100
	nextExp = 10
	speed = 100
	tick = 0
	attack = 10
	defense = 5
	activeBar = 0
	baseNextExp = 50
	hpRate = 30
	attackRate = 5
	defenseRate = 3

	return
}

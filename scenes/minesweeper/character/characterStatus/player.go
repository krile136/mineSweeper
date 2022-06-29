package characterStatus

import (
	"math"
)

type Player struct {
	*abstractCharacterStatus
}

// コンストラクタ
func (p Player) New(lv int) (new CharacterStatusInterface) {
	ac := p.makeAbstractCharacterStatus(p.defaultField())
	ac.nextExp = ac.calcNextExp(ac.lv)
	new = Player{ac}
	return
}

// 攻撃をする
func (p Player) AttackTo(target CharacterStatusInterface) CharacterStatusInterface {
	damage := p.calcDamage(p.attack, target.getDefense())

	newTarget := target.reduceHp(damage)
	return newTarget
}

// レベルアップを実施する
func (p Player) LevelUp(exp int) (bool, CharacterStatusInterface) {
	new := p
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
func (p Player) Update() CharacterStatusInterface {
	new := p
	new.tick += 1
	p.activeBar = float64(new.tick) / float64(new.speed)
	return new
}

// 守備力を取得する
func (p Player) getDefense() float64 {
	return p.defense
}

// HPを減少させる
func (p Player) reduceHp(damage float64) CharacterStatusInterface {
	new := p

	new.hp = math.Max(0, p.hp-damage)
	return new
}

// バトルに関するステータスをセットする
func (p Player) setBattleStatus(
	lv int,
	hp, maxHp, attack, defense float64,
	nextExp int,
) CharacterStatusInterface {
	new := p
	new.lv = lv
	new.hp = hp
	new.maxHp = maxHp
	new.attack = attack
	new.defense = defense
	new.nextExp = nextExp

	return new
}

// デフォルトのフィールド値を取得する
func (p Player) defaultField() (
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

package characterStatus

import (
	"fmt"
	"math"
)

type abstractCharacterStatus struct {
	CharacterStatusInterface
	name        string
	lv          int
	hp          float64
	maxHp       float64
	nextExp     int
	speed       int
	tick        int
	attack      float64
	defense     float64
	activeBar   float64
	baseNextExp int
	hpRate      float64
	attackRate  float64
	defenseRate float64
	turn        bool
	conditions  []condition
}

func (a *abstractCharacterStatus) Lv() int {
	return a.lv
}

func (a *abstractCharacterStatus) Hp() int {
	return int(a.hp)
}

func (a *abstractCharacterStatus) MaxHp() int {
	return int(a.maxHp)
}

func (a *abstractCharacterStatus) NextExp() int {
	return a.nextExp
}

func (a *abstractCharacterStatus) Turn() bool {
	return a.turn
}

func (a *abstractCharacterStatus) Name() string {
	return a.name
}

func (a *abstractCharacterStatus) ActiveBar() float64 {
	return a.activeBar
}

func (a *abstractCharacterStatus) StopTimer() bool {
	isDead := false
	isAppear := false
	for _, condition := range a.conditions {
		if condition == Dead {
			isDead = true
		}
		if condition == Appearing {
			isAppear = true
		}
	}
	return isDead || isAppear
}

func (a *abstractCharacterStatus) Dead() bool {
	for _, condition := range a.conditions {
		if condition == Dead {
			return true
		}
	}
	return false
}

func (a *abstractCharacterStatus) Appearing() bool {
	for _, condition := range a.conditions {
		if condition == Appearing {
			return true
		}
	}
	return false
}

func (a *abstractCharacterStatus) CanTurnOn() bool {
	return a.tick >= a.speed
}

func (a *abstractCharacterStatus) calcDamage(currentAttack, targetDefense float64) (damage float64) {
	damage = math.Max(1, math.Floor(currentAttack*0.5-targetDefense*0.25))
	return
}

// 経験値付与およびステータス変化
func (a *abstractCharacterStatus) addExp(exp int) (
	isLevelUp bool,
	lv int,
	hp, maxHp, attack, defense float64,
	nextExp int,
) {
	lv = a.lv
	hp = a.hp
	maxHp = a.maxHp
	attack = a.attack
	defense = a.defense
	isLevelUp = false
	nextExp = a.nextExp
	remainingExp := exp

	fmt.Printf("getExp: %d\n", exp)

	for exp > 0 {
		currentNextExp := nextExp
		remainingExp -= currentNextExp
		fmt.Printf("currentNextExp: %d       remainingExp: %d\n", currentNextExp, remainingExp)

		if remainingExp < 0 {
			nextExp = remainingExp * -1
			break
		}
		lv += 1
		maxHp += a.hpRate
		hp = maxHp
		attack += a.attackRate
		defense += a.defenseRate
		isLevelUp = true
		nextExp = a.calcNextExp(lv)
	}
	return
}

func (a *abstractCharacterStatus) makeAbstractCharacterStatus(
	name string,
	lv int,
	hp, maxHp float64,
	nextExp, speed, tick int,
	attack, defense, activeBar float64,
	baseNextExp int,
	hpRate, attackRate, defenseRate float64,
) (ac *abstractCharacterStatus) {
	ac = &abstractCharacterStatus{
		name:        name,
		lv:          lv,
		hp:          hp,
		maxHp:       maxHp,
		nextExp:     nextExp,
		speed:       speed,
		tick:        tick,
		attack:      attack,
		defense:     defense,
		activeBar:   activeBar,
		baseNextExp: baseNextExp,
		hpRate:      hpRate,
		attackRate:  attackRate,
		defenseRate: defenseRate,
	}
	return
}

func (a *abstractCharacterStatus) calcNextExp(lv int) (next int) {
	// 次の経験値の指数関数部分 base * (1.1 ^ (lv -1))
	nextExpExponential := int(math.Floor(float64(a.baseNextExp) * math.Pow(1.1, float64(lv-1))))

	// 次の経験値の比例関数部分 lv * 15
	nextExpProportional := lv * 15

	next = (nextExpExponential + nextExpProportional) / 2
	return
}

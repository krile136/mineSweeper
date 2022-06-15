package minesweeper

import (
	"fmt"
	"math"
)

// ステータス
type status struct {
	name         string
	lv           int
	hp           float64
	maxHp        float64
	nextExp      int
	speed        int
	tick         int
	blinkingTick int
	turn         bool
	move         float64
	diff         float64
	attack       float64
	defense      float64
	activeBar    float64
	destroyed    bool
	isAppearing  bool
}

// キャラクター区分
type Character int

const (
	PlayerBattlePositionX   = 110
	EnemyBattlePositionX    = 195
	BaseLevelUpExp          = 50
	DestroyBlinkingMaxTick  = 100
	DestroyBlinkingInterval = 3

	Player Character = iota
	Slime
)

func (c Character) String() string {
	switch c {
	case Player:
		return "character"
	case Slime:
		return "slime"
	default:
		return ""
	}
}

func (c Character) getInitialPlayerStatus() status {
	var playerStatus status = status{
		name:         "",
		lv:           1,
		hp:           100,
		maxHp:        100,
		nextExp:      10,
		speed:        100,
		tick:         0,
		blinkingTick: 0,
		turn:         false,
		move:         1,
		diff:         0,
		attack:       10,
		defense:      5,
		activeBar:    0,
		destroyed:    false,
		isAppearing:  false,
	}
	playerStatus.name = c.String()
	player.nextExp = calcNextLevelUpExp()

	return playerStatus
}

func (c Character) getInitialEnemyStatus() (maxHp float64, speed int, attack float64, defense float64) {
	switch c {
	case Slime:
		var maxHp float64 = 50
		var speed int = 160
		var attack float64 = 5
		var defense float64 = 1
		return maxHp, speed, attack, defense
	default:
		return 0, 0, 0, 0
	}
}

func (c Character) enemyFactory(lv float64) status {
	maxHp, speed, attack, defense := c.getInitialEnemyStatus()
	var enemy status = status{
		name:         "",
		lv:           1,
		hp:           maxHp,
		maxHp:        maxHp,
		nextExp:      0,
		speed:        speed,
		tick:         0,
		blinkingTick: 0,
		turn:         false,
		move:         -1,
		diff:         0,
		attack:       attack,
		defense:      defense,
		activeBar:    0,
		destroyed:    false,
		isAppearing:  false,
	}
	enemy.name = c.String()

	// レベルに応じてステータス変化
	hpRate, attackRate, defenseRate := c.statusGrowthRate()
	enemy.lv = int(lv)
	enemy.maxHp += hpRate * (lv - 1)
	enemy.hp = enemy.maxHp
	enemy.attack = attackRate * (lv - 1)
	enemy.defense = defenseRate * (lv - 1)

	return enemy
}

// レベルアップ時のステータスの伸び率（定数だけど）を返す
func (c Character) statusGrowthRate() (hpRate, attackRate, defenseRate float64) {
	switch c {
	case Player:
		return 30, 5, 3
	case Slime:
		return 10, 2, 1
	default:
		return 0, 0, 0
	}
}

// 次のレベルアップまでの経験値を計算する
func calcNextLevelUpExp() int {
	// 次の経験値の指数関数部分 base * (1.1 ^ (lv -1))
	NextExpExponential := int(math.Floor(float64(BaseLevelUpExp) * math.Pow(1.1, float64(player.lv-1))))
	// 次の経験値の比例関数部分 lv * 15
	NextExpProportional := player.lv * 15
	return (NextExpExponential + NextExpProportional) / 2
}

// レベルアップを管理する
func levelUp(exp int) (bool, int, int) {
	isLevelUp := false
	from := player.lv
	to := player.lv

	fmt.Printf("Get EXP: %d\n", exp)
	for exp > 0 {
		currentNextExp := player.nextExp
		exp -= currentNextExp
		if exp >= 0 {
			player.lv += 1
			player.maxHp += 20
			player.hp = player.maxHp
			player.attack += 5
			player.defense += 3
			to = player.lv
			isLevelUp = true
			player.nextExp = calcNextLevelUpExp()
			fmt.Printf("Level up: %d, PlayerNextExp: %d\n", player.lv, player.nextExp)
		} else {
			player.nextExp = exp * -1
			fmt.Printf("not Level up    next: %d\n", player.nextExp)
		}
	}
	return isLevelUp, from, to
}

func isShowBlinking() bool {
	if enemy.blinkingTick%(DestroyBlinkingInterval*2) < DestroyBlinkingInterval && enemy.blinkingTick <= DestroyBlinkingMaxTick {
		return true
	}
	return false
}

package minesweeper

import "image/color"

// メッセージ構造体
type message struct {
	value      string
	messageDiv MessageDiv
	x          float64
	y          float64
	tick       int
	crl        color.Color
}

// メッセージの時間を進める
func (m *message) update() error {
	mx, my := m.messageDiv.getMoveValue(m.tick)
	m.tick += 1
	m.x += mx
	m.y += my
	return nil
}

// メッセージの時間が存在可能時間内かチェックする
func (m message) isExist() bool {
	return m.tick <= (m.messageDiv.getExistTick())
}

// メッセージ区分
type MessageDiv int

const (
	PlayerDamage MessageDiv = iota
	EnemyDamage
	LevelUp
)

func (m MessageDiv) String() string {
	switch m {
	case PlayerDamage:
		return "PlayerDamage"
	case EnemyDamage:
		return "EnemyDamage"
	case LevelUp:
		return "LevelUp"
	default:
		return ""
	}
}

func (m MessageDiv) getExistTick() int {
	switch m {
	case PlayerDamage:
		return 40
	case EnemyDamage:
		return 40
	case LevelUp:
		return 70
	default:
		return 0
	}
}

func (m MessageDiv) getMoveValue(tick int) (x, y float64) {
	switch m {
	case PlayerDamage:
		g := float64(0.1)
		v := g * float64(m.getExistTick()) * -0.5
		diffY := v + g*float64(tick)
		return -1, diffY
	case EnemyDamage:
		g := float64(0.1)
		v := g * float64(m.getExistTick()) * -0.5
		diffY := v + g*float64(tick)
		return 1, diffY
	case LevelUp:
		g := float64(-0.1)
		v := g * float64(m.getExistTick()) * -0.5
		diffY := v + g*float64(tick)
		return 0, diffY
	default:
		return 0, 0
	}
}

func (m MessageDiv) getDefaultPosition() (x, y float64) {
	switch m {
	case PlayerDamage:
		return 100, 55
	case EnemyDamage:
		return 210, 55
	case LevelUp:
		return 100, 0
	default:
		return 0, 0
	}
}

package minesweeper

// メッセージ構造体
type message struct {
	value      string
	messageDiv MessageDiv
	x          float64
	y          float64
	tick       int
}

// メッセージの時間を進める
func (m *message) update() error {
	m.tick += 1
	mx, my := m.messageDiv.getMoveValue(m.tick)
	m.x += mx
	m.y += my
	return nil
}

// メッセージの時間が存在可能時間内かチェックする
func (m message) isExist() bool {
	return m.tick < m.messageDiv.getExistTick()
}

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
		return 40
	default:
		return 0
	}
}

func (m MessageDiv) getMoveValue(tick int) (x, y float64) {
	switch m {
	case PlayerDamage:
		g := float64(0.01)
		v := (g * -0.25) * float64(m.getExistTick())
		diffY := v*float64(tick) + g*0.5*float64(tick*tick)
		return -1, diffY
	case EnemyDamage:
		g := float64(0.01)
		v := (g * -0.25) * float64(m.getExistTick())
		diffY := v*float64(tick) + g*0.5*float64(tick*tick)
		return 1, diffY
	case LevelUp:
		return 0, 2
	default:
		return 0, 0
	}
}

func (m MessageDiv) getDefaultPosition() (x, y float64) {
	switch m {
	case PlayerDamage:
		return 100, 45
	case EnemyDamage:
		return 210, 45
	case LevelUp:
		return 0, 2
	default:
		return 0, 0
	}
}

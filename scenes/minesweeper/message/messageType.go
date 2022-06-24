package message

type MessageType int

const (
	PlayerDamage MessageType = iota
	EnemyDamage
	LevelUp
)

package minesweeper

import (
	"github.com/krile136/mineSweeper/scenes/minesweeper/message"
	"github.com/krile136/mineSweeper/scenes/minesweeper/message/messages"
)

var MessageMap = map[message.MessageType]messages.MessageInterface{
	message.PlayerDamage: messages.PlayerDamage{},
	message.EnemyDamage:  messages.EnemyDamage{},
	message.LevelUp:      messages.LevelUp{},
}

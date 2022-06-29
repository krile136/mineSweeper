package minesweeper

import (
	"github.com/krile136/mineSweeper/scenes/minesweeper/character"
	"github.com/krile136/mineSweeper/scenes/minesweeper/character/characterStatus"
)

var CharacterStatusMap = map[character.CharacterType]characterStatus.CharacterStatusInterface{
	character.Player: characterStatus.Player{},
	character.Slime:  characterStatus.Slime{},
}

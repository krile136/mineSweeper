package minesweeper

import (
	"fmt"

	"github.com/krile136/mineSweeper/scenes/minesweeper/character"
	"github.com/krile136/mineSweeper/scenes/minesweeper/character/characterDraw"
	"github.com/krile136/mineSweeper/scenes/minesweeper/character/characterStatus"
)

var playerStatusSlice []characterStatus.CharacterStatusInterface
var playerDrawSlice []characterDraw.CharacterDrawInterface
var enemyStatusSlice []characterStatus.CharacterStatusInterface
var enemyDrawSlice []characterDraw.CharacterDrawInterface

// ゲーム開始直後のキャラクターとエネミーをセットする
func setInitialCharacter() {
	player = playerStatusSlice[0]
	playerStatusSlice = playerStatusSlice[:1]
	playerDraw = playerDrawSlice[0]
	playerDrawSlice = playerDrawSlice[:1]

	enemy = enemyStatusSlice[0]
	enemyStatusSlice = enemyStatusSlice[1:]
	enemyDraw = enemyDrawSlice[0]
	enemyDrawSlice = enemyDrawSlice[1:]
}

// 次のエネミーをセットする
func setNextEnemy() {
	enemy = enemyStatusSlice[0]
	enemyStatusSlice = enemyStatusSlice[1:]
	enemyDraw = enemyDrawSlice[0]
	enemyDrawSlice = enemyDrawSlice[1:]
	fmt.Printf("new enemy name: %s\n", enemy.Name())
	fmt.Printf("new enemy Lv: %d\n", enemy.Lv())
}

// キャラクターの配列を初期化する
func initCharacterSlice() {
	// プレイヤーに関する初期値をセット
	setPlayerSlice(character.Player, 1)

	// モンスターに関する初期値をセット
	setInitialEnemySlice(character.Slime, 1)
	setEnemySlice(character.Slime, 3)
}

// プレイヤーに関するデータを詰める
func setPlayerSlice(t character.CharacterType, lv int) {
	status := characterStatusMap[t].New(lv)
	draw := characterDrawMap[t].New()

	playerStatusSlice = append(playerStatusSlice, status)
	playerDrawSlice = append(playerDrawSlice, draw)
}

// 敵に関するデータを詰める
func setEnemySlice(t character.CharacterType, lv int) {
	status := characterStatusMap[t].New(lv)
	draw := characterDrawMap[t].New()

	status = status.SetInitialStatus()
	draw = draw.SetInitialDraw()

	enemyStatusSlice = append(enemyStatusSlice, status)
	enemyDrawSlice = append(enemyDrawSlice, draw)
}

// 一番最初の敵のステータスを詰める
func setInitialEnemySlice(t character.CharacterType, lv int) {
	status := characterStatusMap[t].New(lv)
	draw := characterDrawMap[t].New()

	enemyStatusSlice = append(enemyStatusSlice, status)
	enemyDrawSlice = append(enemyDrawSlice, draw)
}

var characterStatusMap = map[character.CharacterType]characterStatus.CharacterStatusInterface{
	character.Player: characterStatus.Player{},
	character.Slime:  characterStatus.Slime{},
}

var characterDrawMap = map[character.CharacterType]characterDraw.CharacterDrawInterface{
	character.Player: characterDraw.Player{},
	character.Slime:  characterDraw.Slime{},
}

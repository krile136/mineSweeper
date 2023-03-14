package minesweeper

import (
	"image/color"
	"math"
	"math/rand"
	"strconv"

	"github.com/krile136/mineSweeper/scenes/minesweeper/explode"
	"github.com/krile136/mineSweeper/scenes/minesweeper/message"
	"github.com/krile136/mineSweeper/scenes/minesweeper/message/messages"
	"github.com/krile136/mineSweeper/store"
)

func (m *MineSweeper) placeBombs() error {
	// フィールドを初期化
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			m.field[i][j] = close
		}
	}

	// 重複のないランダムな数字をrows * columns の中からBombsNumberの分だけ作り
	// 配列に保存する
	m.bombsPosition = make([]int, m.bombsNumber)
	count := 0
	for count < m.bombsNumber {
		position := rand.Intn(m.rows * m.columns)
		if !inArray(m.bombsPosition, position) {
			m.bombsPosition[count] = position
			count++
		}
	}

	return nil
}

func (m *MineSweeper) placeFlag(x, y int) error {

	m.field[y][x] = flag

	return nil
}

// 周りの爆弾の数をカウントし、その数に応じて引数に渡されたマスのフィールドを決定する
func (m *MineSweeper) searchAround(x, y int) (exp, combo int) {
	var bombs int = 0
	var next []int
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if inBetween(0, i, m.rows-1) && inBetween(0, j, m.columns-1) {
				position := i*m.rows + j
				if inArray(m.bombsPosition, position) {
					bombs += 1
				} else {
					if m.field[i][j] == close {
						next = append(next, position)
					}
				}
			}
		}
	}
	if bombs == 0 {
		nextCheck = append(nextCheck, next...)
	}
	exp = 0
	combo = 0
	// フラグがおいてあるマスはフィールド情報を更新しない
	if m.field[y][x] != flag && m.field[y][x] == close {
		m.field[y][x] = nums[bombs]
		exp = 1
		combo = 1
	}
	if len(nextCheck) > 0 {
		nextCheck = nextCheck[1:]
	}
	return
}

// 周りの爆弾の数が表示されているフィールドにて、周囲のフィールドを走査する
func (m *MineSweeper) searchAroundOnNumberField(x, y int) (isExsistsBomb bool) {
	var bombs int = 0
	var bomb_interval int = 5
	var next []int
	isExsistsBomb = false
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if inBetween(0, i, m.rows-1) && inBetween(0, j, m.columns-1) {
				position := i*m.rows + j
				if inArray(m.bombsPosition, position) {
					// 周りのマスを開いたときに爆弾があった場合
					if m.field[i][j] != flag && m.field[i][j] != bomb {
						// フラグおいてないので爆発ダメージ
						m.field[i][j] = bomb
						addExplodes(bombs * bomb_interval)
						bombs += 1
						isExsistsBomb = true
					}
				} else {
					if m.field[i][j] == close {
						next = append(next, position)
					}
				}
			}
		}
	}
	nextCheck = append(nextCheck, next...)

	return
}

func (m *MineSweeper) executePlayerTurn() {
	// タイマーが止まるときの専用処理
	if enemy.StopTimer() {
		if playerDraw.IsReturningToBase() {
			playerDraw = playerDraw.ExecuteMoving()
		} else {
			player = player.FinishTurn()
			playerDraw = playerDraw.FinishTurn()
		}
	} else {
		playerDraw = playerDraw.ExecuteMoving()
	}

	if playerDraw.CanExecuteInvertAtTop() {
		playerDraw = playerDraw.InvertDirection()
		damage := player.GetDamageAmount(enemy)
		enemy = player.AttackTo(enemy)
		messageStruct := MessageMap[message.EnemyDamage]
		var value string = strconv.FormatFloat(damage, 'f', 0, 64)
		displayMessages = append(displayMessages, messageStruct.New(value))
	}
	if playerDraw.CanExecuteInvertAtBase() {
		player = player.FinishTurn()
		playerDraw = playerDraw.FinishTurn()
		if enemy.StopTimer() {
			// タイマーストップ状態ではプレイヤーのターン継続する（敵が現れたらターン終了）
			player = player.InvertTurn()
		}
	}

	if enemy.Dead() {
		enemyDraw = enemyDraw.UpdateBlinking()
		if enemyDraw.IsFinishDeadBlinking() {
			setNextEnemy()
		}
	}
	if enemy.Appearing() {
		enemyDraw.ExecuteMoving()
		if enemyDraw.CanFinishAppearing() {
			enemy = enemy.ResetCondition()
			enemy = enemy.FinishTurn()
			player = player.InvertTurn()
		}
	}
}

func (m *MineSweeper) executeEnemyTurn() error {
	if enemy.Turn() == false {
		// 敵のターンでなければなんの処理もしない
		return nil
	}

	enemyDraw = enemyDraw.ExecuteMoving()
	if enemyDraw.CanExecuteInvertAtTop() {
		enemyDraw = enemyDraw.InvertDirection()
		damage := enemy.GetDamageAmount(player)
		player = enemy.AttackTo(player)
		messageStruct := MessageMap[message.PlayerDamage]
		var value string = strconv.FormatFloat(damage, 'f', 0, 64)
		displayMessages = append(displayMessages, messageStruct.New(value))
	}
	if enemyDraw.CanExecuteInvertAtBase() {
		enemyDraw = enemyDraw.FinishTurn()
		enemy = enemy.FinishTurn()
	}

	return nil
}

func addExplodes(delay int) {
	for i := 0; i < 5; i++ {
		rdmX := rand.Float64()*20 - 10
		rdmY := rand.Float64()*20 - 10
		var baseX float64 = float64(store.Data.Layout.OutsideWidth)/2 - 45
		var baseY float64 = float64(60)
		explodes = explodes.Add(explode.Orange, baseX+rdmX, baseY+rdmY, -10*i, delay)
	}
}

func updateMessage() {
	tempMessages := []messages.MessageInterface{}
	for _, message := range displayMessages {
		newMessage := message.Update()
		if newMessage.IsExist() {
			tempMessages = append(tempMessages, newMessage)
		}
	}
	displayMessages = tempMessages
}

func updateExplodes() {
	var finishedExplodes int = 0
	explodes, finishedExplodes = explodes.Update()
	for i := 0; i < finishedExplodes; i++ {
		explodeDamage := float64(player.MaxHp()) * 0.05
		messageStruct := MessageMap[message.PlayerDamage]
		var value string = strconv.FormatFloat(explodeDamage, 'f', 0, 64)
		displayMessages = append(displayMessages, messageStruct.New(value))
		player = player.ReduceHp(explodeDamage)
	}
}

// 虹色表示の管理インデックスを更新する
func updateRainbowIndex() {
	rainbowIndex += 1
	if rainbowIndex > 6 {
		rainbowIndex = 0
	}
}

// コンボ維持時間を更新する
// 0になったらコンボが切れる
func updateCurrentComboTick() {
	currentComboTick = int(math.Max(0, float64(currentComboTick)-1))
	if currentComboTick == 0 && currentCombo > 0 {
		currentCombo = 0
	}
}

func getRainbow() (r color.Color) {
	r = store.Data.Color.Red
	switch rainbowIndex {
	case 1:
		r = store.Data.Color.Orange
	case 2:
		r = store.Data.Color.Yellow
	case 3:
		r = store.Data.Color.Green
	case 4:
		r = store.Data.Color.Cyan
	case 5:
		r = store.Data.Color.Blue
	case 6:
		r = store.Data.Color.Purple
	}
	return
}

// int型の配列の中に特定のint型の値が含まれるかチェックする
func inArray(array []int, needle int) bool {
	for _, val := range array {
		if needle == val {
			return true
		}
	}
	return false
}

// int型の値が最小と最大の間にあるかチェックする
func inBetween(min, val, max int) bool {
	if (val >= min) && (val <= max) {
		return true
	} else {
		return false
	}
}

// float64型の値が最大値と最小値の間に収まるようにする
func setBetween(min, val, max float64) float64 {
	return math.Min(max, math.Max(min, val))
}

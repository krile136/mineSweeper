package minesweeper

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/krile136/mineSweeper/enum/route"
	"github.com/krile136/mineSweeper/scenes/minesweeper/message"
	"github.com/krile136/mineSweeper/scenes/scene"
	"github.com/krile136/mineSweeper/store"
)

func (m *MineSweeper) Update() error {
	if scene.Is_just_changed {
		m.init()
	} else {
		// if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		// 	m.placeBombs()
		// }

		// スクロールしたときの処理
		wheelX, wheelY := ebiten.Wheel()
		scrollCorrectionValue := store.Data.Env.ScrollCorrectionValue
		scrollX = setBetween(-float64(maxScrollX), scrollX+wheelX*float64(scrollCorrectionValue), 0)
		scrollY = setBetween(-float64(maxScrollY), scrollY+wheelY*float64(scrollCorrectionValue), 0)

		// スクロールされている間だけスクロールバーのサイズと位置を計算する
		if wheelX != 0 || wheelY != 0 {
			isBarDisplay = true
			BarDisplayFrame = 30
			barLengthY = math.Max(0.5, float64(store.Data.Layout.OutsideHeight-store.Data.Layout.BattleField)/(float64(store.Data.Layout.OutsideHeight-store.Data.Layout.BattleField)+maxScrollY)) * float64(store.Data.Layout.OutsideHeight-store.Data.Layout.BattleField)
			barSlideY = float64(store.Data.Layout.BattleField) + ((float64(store.Data.Layout.OutsideHeight-store.Data.Layout.BattleField)-barLengthY)/maxScrollY)*math.Abs(scrollY) + barLengthY/2
			barLengthX = math.Max(0.5, float64(store.Data.Layout.OutsideWidth)/(float64(store.Data.Layout.OutsideWidth)+maxScrollX)) * float64(store.Data.Layout.OutsideWidth)
			barSlideX = ((float64(store.Data.Layout.OutsideWidth)-barLengthX)/maxScrollX)*math.Abs(scrollX) + barLengthX/2
		} else {
			if isBarDisplay {
				BarDisplayFrame -= 1
				if BarDisplayFrame <= 0 {
					isBarDisplay = false
				}
			}
		}

		// マウスの座標をスクロールとbattleFieldの分だけ補正する
		mouse_x, mouse_y := ebiten.CursorPosition()
		y := (mouse_y - int(scrollY) - store.Data.Layout.BattleField) / blockWidth
		x := (mouse_x - int(scrollX)) / blockWidth

		// クリック処理の前提条件として、マウスY座標がbattleFieldより下であること
		if mouse_y > store.Data.Layout.BattleField {
			// 左クリックしたときの処理
			m.leftClick(x, y)

			// 右クリックしたときの処理
			m.rightClick(x, y)
		}

		// バトル周りの処理

		// 攻撃タイミング処理
		// タイマーを進めて、Speedと同じになった攻撃ターン
		// どちらかの攻撃ターンになったときは、攻撃が完了するまで相手のタイマーは止まる
		if !player.Turn() && !enemy.Turn() {
			player = player.Update()
			enemy = enemy.Update()
		}

		if enemy.Turn() {
			m.executeEnemyTurn()
		}

		if player.Turn() {
			m.executePlayerTurn()
		}

		if player.CanTurnOn() {
			player = player.SetTurn(true)
		}
		if enemy.CanTurnOn() && !player.Turn() {
			enemy = enemy.SetTurn(true)
		}

		// メッセージの更新処理
		updateMessage()

		// 爆発の更新処理
		updateExplodes()

		// 虹色表示の管理インデックス更新
		updateRainbowIndex()

		// 現在のコンボ表示の更新
		updateCurrentComboTick()

		// クリア文字の時間の更新をする
		updateClearTick()

		// すべてマスを開けたときの文字の表示を更新する
		updateAllOpenTick()

		if player.Dead() {
			playerDraw = playerDraw.UpdateBlinking()
			if playerDraw.IsFinishDeadBlinking() {
				// プレイヤーがやられたときの点滅終了時
				// 現在のスコアを保存してゲームオーバー画面へ
				store.Data.CurrentScore = score
				scene.RouteType = route.GameOver
			}
		}
	}

	return nil
}

// 左クリックしたときの処理
func (m *MineSweeper) leftClick(x, y int) error {
	if player.Dead() || isClear == true {
		return nil
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {

		// クリックしたマスがフラグが立っていれば何もしない
		if m.field[y][x] == flag || m.field[y][x] == bomb {
			return nil
		}

		position := y*m.rows + x
		if inArray(m.bombsPosition, position) {
			// 爆弾があるのでダメージ
			m.field[y][x] = bomb
			currentCombo = 0
			addExplodes(0)
		} else {
			var getExp, getCombo int = m.searchAround(x, y)
			for len(nextCheck) > 0 {
				search_y := nextCheck[0] / m.rows
				search_x := nextCheck[0] % m.rows
				var exp, combo int = m.searchAround(search_x, search_y)
				getExp += exp
				getCombo += combo
			}
			isLevelUp := false
			score += getExp
			isLevelUp, player = player.LevelUp(getExp)

			// 得られるコンボ数はまとめて開いても、一つ開いても１しか増えない
			getCombo = int(math.Min(1, float64(getCombo)))
			if getCombo > 0 {
				currentComboTick = maxComboTick
			}
			currentCombo += getCombo
			if isLevelUp {
				messageStruct := MessageMap[message.LevelUp]
				displayMessages = append(displayMessages, messageStruct.New(messageStruct.String()))
			}
		}
		if m.checkAllOpen() {
			m.relocationBombs()
		}
	}

	return nil
}

// 右クリックしたときの処理
func (m *MineSweeper) rightClick(x, y int) error {
	if player.Dead() || isClear == true {
		return nil
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonRight) {
		switch m.field[y][x] {
		case close:
			m.placeFlag(x, y)
		case flag:
			m.field[y][x] = close
		case one, two, three, four, five, six, seven, eight:
			var isExsistsBombs bool = m.searchAroundOnNumberField(x, y)
			var getExp, getCombo int = 0, 0
			for len(nextCheck) > 0 {
				search_y := nextCheck[0] / m.rows
				search_x := nextCheck[0] % m.rows
				var exp, combo int = m.searchAround(search_x, search_y)
				getExp += exp
				getCombo += combo
			}
			isLevelUp := false
			score += getExp
			isLevelUp, player = player.LevelUp(getExp)

			if isExsistsBombs {
				// 爆弾を踏んだらコンボはリセットする
				currentCombo = 0
			} else {
				// 得られるコンボ数はまとめて開いても、一つ開いても１しか増えない
				getCombo = int(math.Min(1, float64(getCombo)))
				if getCombo > 0 {
					currentComboTick = maxComboTick
				}
				currentCombo += getCombo
			}
			if isLevelUp {
				messageStruct := MessageMap[message.LevelUp]
				displayMessages = append(displayMessages, messageStruct.New(messageStruct.String()))
			}
			if m.checkAllOpen() {
				m.relocationBombs()
			}
		default:
			// 何もしない
		}
	}

	return nil
}

package minesweeper

import (
	"log"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/krile136/mineSweeper/scenes/minesweeper/message"
	"github.com/krile136/mineSweeper/scenes/minesweeper/message/messages"
	"github.com/krile136/mineSweeper/scenes/scene"
	"github.com/krile136/mineSweeper/store"
)

func (m *MineSweeper) Update() error {
	// シーン切替時にstoreから行と列のデータを持ってくる
	if scene.Is_just_changed {
		m.rows = store.Data.MineSweeper.Rows
		m.columns = store.Data.MineSweeper.Columns
		m.bombsNumber = store.Data.MineSweeper.BombsNumber

		// rowsとcolumnsからフィールドを作成
		m.field = make([][]int, m.rows)
		for i := 0; i < m.columns; i++ {
			m.field[i] = make([]int, m.columns)
		}

		// 爆弾を配置する
		m.placeBombs()

		isBarDisplay = false

		// スクロール可能値を計算する
		// ブロックの大きさは、setWIndowの幅 / Layoutの幅に拡大される
		maxScrollX = math.Max(0, (32*float64(store.Data.MineSweeper.Columns))-float64(store.Data.Layout.OutsideWidth))
		maxScrollY = math.Max(0, (32*float64(store.Data.MineSweeper.Rows)+float64(store.Data.Layout.BattleField))-float64(store.Data.Layout.OutsideHeight))

		// ゲームに関するデータを初期化する
		displayMessages = []messages.MessageInterface{}

		// 各キャラクターの初期ステータスなどが入った配列を初期化する
		initCharacterSlice()

		// playerとenemyに初期値をセットする
		setInitialCharacter()
	} else {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			m.placeBombs()
		}
	}

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
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			// クリックしたマスがフラグが立っていれば何もしない
			if m.field[y][x] != flag {
				position := y*m.rows + x
				if inArray(m.bombsPosition, position) {
					// 爆弾があるのでゲームオーバー
					log.Print("game over! (left click)")
					m.field[y][x] = bomb
				} else {
					GetExp = 0
					m.searchAround(x, y)
					for len(nextCheck) > 0 {
						search_y := nextCheck[0] / m.rows
						search_x := nextCheck[0] % m.rows
						m.searchAround(search_x, search_y)
					}
					isLevelUp := false
					isLevelUp, ply = ply.LevelUp(GetExp)
					if isLevelUp {
						messageStruct := MessageMap[message.LevelUp]
						displayMessages = append(displayMessages, messageStruct.New(messageStruct.String()))
					}
				}
			}
		}

		// 右クリックしたときの処理
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonRight) {
			switch m.field[y][x] {
			case close:
				m.placeFlag(x, y)
			case flag:
				m.field[y][x] = close
			case one, two, three, four, five, six, seven, eight:
				GetExp = 0
				m.searchAroundOnNumberField(x, y)
				for len(nextCheck) > 0 {
					search_y := nextCheck[0] / m.rows
					search_x := nextCheck[0] % m.rows

					m.searchAround(search_x, search_y)
				}
				isLevelUp := false
				isLevelUp, ply = ply.LevelUp(GetExp)
				if isLevelUp {
					messageStruct := MessageMap[message.LevelUp]
					displayMessages = append(displayMessages, messageStruct.New(messageStruct.String()))
				}

			default:
				// 何もしない
			}
		}
	}

	// バトル周りの処理

	// 攻撃タイミング処理
	// タイマーを進めて、Speedと同じになった攻撃ターン
	// どちらかの攻撃ターンになったときは、攻撃が完了するまで相手のタイマーは止まる
	if !ply.Turn() && !enmy.Turn() {
		ply = ply.Update()
		enmy = enmy.Update()
	}

	if enmy.Turn() {
		enemyDraw = enemyDraw.ExecuteMoving()
		if enemyDraw.CanExecuteInvertAtTop() {
			enemyDraw = enemyDraw.InvertDirection()
			damage := enmy.GetDamageAmount(ply)
			ply = enmy.AttackTo(ply)
			messageStruct := MessageMap[message.PlayerDamage]
			var value string = strconv.FormatFloat(damage, 'f', 0, 64)
			displayMessages = append(displayMessages, messageStruct.New(value))
		}
		if enemyDraw.CanExecuteInvertAtBase() {
			enemyDraw = enemyDraw.FinishTurn()
			enmy = enmy.FinishTurn()
		}

	}

	if ply.Turn() {
		// タイマーが止まるときの専用処理
		if enmy.StopTimer() {
			if playerDraw.IsReturningToBase() {
				playerDraw = playerDraw.ExecuteMoving()
			} else {
				ply = ply.FinishTurn()
				playerDraw = playerDraw.FinishTurn()
			}
		} else {
			playerDraw = playerDraw.ExecuteMoving()
		}

		if playerDraw.CanExecuteInvertAtTop() {
			playerDraw = playerDraw.InvertDirection()
			damage := ply.GetDamageAmount(enmy)
			enmy = ply.AttackTo(enmy)
			messageStruct := MessageMap[message.EnemyDamage]
			var value string = strconv.FormatFloat(damage, 'f', 0, 64)
			displayMessages = append(displayMessages, messageStruct.New(value))
		}
		if playerDraw.CanExecuteInvertAtBase() {
			ply = ply.FinishTurn()
			playerDraw = playerDraw.FinishTurn()
			if enmy.StopTimer() {
				// タイマーストップ状態ではプレイヤーのターン継続する（敵が現れたらターン終了）
				ply = ply.InvertTurn()
			}
		}

		if enmy.Dead() {
			enemyDraw = enemyDraw.UpdateBlinking()
			if enemyDraw.IsFinishDeadBlinking() {
				setNextEnemy()
			}
		}
		if enmy.Appearing() {
			enemyDraw.ExecuteMoving()
			if enemyDraw.CanFinishAppearing() {
				enmy = enmy.ResetCondition()
				enmy = enmy.FinishTurn()
				ply = ply.InvertTurn()
			}
		}
	}

	if ply.CanTurnOn() {
		ply = ply.SetTurn(true)
	}
	if enmy.CanTurnOn() && !ply.Turn() {
		enmy = enmy.SetTurn(true)
	}

	// メッセージの更新処理
	tempMessages := []messages.MessageInterface{}
	for _, message := range displayMessages {
		newMessage := message.Update()
		if newMessage.IsExist() {
			tempMessages = append(tempMessages, newMessage)
		}
	}
	displayMessages = tempMessages

	return nil
}

package minesweeper

import (
	"fmt"
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
		log.Println(fmt.Sprintf("maxScrollX: %g", maxScrollX))

		// ゲームに関するデータを初期化する
		displayMessages = []messages.MessageInterface{}
		player = Player.getInitialPlayerStatus()
		enemy = Slime.enemyFactory(1)

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
					isLevelUp, _, _ := player.levelUp(GetExp)
					_, ply = ply.LevelUp(GetExp)
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
				isLevelUp, _, _ := player.levelUp(GetExp)
				_, ply = ply.LevelUp(GetExp)
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
	if !player.turn && !enemy.turn {
		player.tick += 1
		enemy.tick += 1
	}
	if !ply.Turn() && !enmy.Turn() {
		ply = ply.Update()
		enmy = enmy.Update()
	}

	if enmy.Turn() {
		enemyDraw.ExecuteMoving()
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

	if enemy.turn {
		enemy.executeMoving()
		if math.Abs(float64(enemy.diff)) >= 50 {
			enemy.invertMove()
			// damage := enemy.attackTo(&player)
			// messageStruct := MessageMap[message.PlayerDamage]
			// var value string = strconv.FormatFloat(damage, 'f', 0, 64)
			// displayMessages = append(displayMessages, messageStruct.New(value))

		}
		if enemy.diff > 0 {
			enemy.invertMove()
			enemy.reset()
			enemy.turn = false
		}
	}

	if player.turn {
		if !enemy.destroyed && !enemy.isAppearing {
			player.executeMoving()
		} else {
			if player.move < 0 {
				player.executeMoving()
			} else {
				player.reset()
			}
		}
		if math.Abs(float64(player.diff)) > 50 {
			player.invertMove()
			damage := player.attackTo(&enemy)
			messageStruct := MessageMap[message.EnemyDamage]
			var value string = strconv.FormatFloat(damage, 'f', 0, 64)
			displayMessages = append(displayMessages, messageStruct.New(value))
		}
		if player.diff <= 0 {
			player.invertMove()
			player.reset()
			if !enemy.destroyed && !enemy.isAppearing {
				player.turn = false
			}
		}
		if enemy.destroyed {
			enemy.blinkingTick += 1
			if enemy.blinkingTick >= DestroyBlinkingMaxTick {
				// 新しいモンスターを出現させる
				lv := enemy.lv
				enemy = Slime.enemyFactory(float64(lv + 1))
				enemy.diff = 150
				enemy.move = -1
				enemy.isAppearing = true
			}
		}
		if enemy.isAppearing {
			enemy.executeMoving()
			if enemy.diff < 0 {
				enemy.diff = 0
				enemy.isAppearing = false
				player.turn = false
			}
		}
	}

	if player.tick >= player.speed {
		player.turn = true
	}
	if enemy.tick >= enemy.speed && !player.turn {
		enemy.turn = true
	}

	player.updateActiveBar()
	enemy.updateActiveBar()

	// log.Println(fmt.Sprintf("PlayerTick: %d,    EnemyTick: %d", PlayerTick, EnemyTick))
	// log.Println(fmt.Sprintf("PlayerActiveBar: %g", PlayerActiveBar))

	// メッセージの更新処理
	tempMessages := []messages.MessageInterface{}
	for _, message := range displayMessages {
		newMessage := message.Update()
		if newMessage.IsExist() {
			tempMessages = append(tempMessages, newMessage)
		}
		// fmt.Printf("messageDiv: %s, value: %s,  exist: %d \n", message.messageDiv.String(), message.value, message.messageDiv.getExistTick()-message.tick)
		// fmt.Printf("mx: %d,  my; %d \n", int(message.x), int(message.y))
	}
	displayMessages = tempMessages

	return nil
}

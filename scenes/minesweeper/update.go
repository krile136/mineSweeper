package minesweeper

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
		messages = []message{}
		PlayerLv = 1
		PlayerHp = 100
		PlayerMaxHp = 100
		PlayerNextExp = calcNextLevelUpExp()
		PlayerSpeed = 100
		PlayerTick = 0
		PlayerTurn = false
		PlayerMove = 1
		PlayerDiff = 0
		PlayerAttack = 10
		PlayerDefense = 5
		PlayerActiveBar = 0
		EnemyLv = 1
		EnemyHp = 100
		EnemyMaxHp = 100
		EnemySpeed = 160
		EnemyTick = 0
		EnemyTurn = false
		EnemyDiff = 0
		EnemyMove = -1
		EnemyAttack = 5
		EnemyDefense = 1
		EnemyActiveBar = 0
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
					isLevelUp, _, _ := levelUp(GetExp)
					if isLevelUp {
						levelUpDefaultX, levelUpDefaultY := LevelUp.getDefaultPosition()
						newMessage := message{
							value:      "Level UP !!",
							messageDiv: LevelUp,
							x:          levelUpDefaultX,
							y:          levelUpDefaultY,
							tick:       0,
						}
						messages = append(messages, newMessage)
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
				isLevelUp, _, _ := levelUp(GetExp)
				if isLevelUp {
					levelUpDefaultX, levelUpDefaultY := LevelUp.getDefaultPosition()
					newMessage := message{
						value:      "Level UP !!",
						messageDiv: LevelUp,
						x:          levelUpDefaultX,
						y:          levelUpDefaultY,
						tick:       0,
					}
					messages = append(messages, newMessage)
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
	if !PlayerTurn && !EnemyTurn {
		PlayerTick += 1
		EnemyTick += 1
	}

	if EnemyTurn {
		EnemyDiff += EnemyMove * 5
		if math.Abs(float64(EnemyDiff)) >= 50 {
			EnemyMove = EnemyMove * -1
			damage := math.Max(1, math.Floor(float64(EnemyAttack)*0.5-float64(PlayerDefense)*0.25))
			defaultPositionX, defaultPositionY := PlayerDamage.getDefaultPosition()
			newMessage := message{
				value:      strconv.FormatFloat(damage, 'f', 0, 64),
				messageDiv: PlayerDamage,
				x:          defaultPositionX,
				y:          defaultPositionY,
				tick:       0,
			}
			messages = append(messages, newMessage)
			PlayerHp -= damage
		}
		if EnemyDiff > 0 {
			EnemyMove = EnemyMove * -1
			EnemyDiff = 0
			EnemyTick = 0
			EnemyTurn = false
		}
	}

	if PlayerTurn {
		PlayerDiff += PlayerMove * 5
		if math.Abs(float64(PlayerDiff)) > 50 {
			PlayerMove = PlayerMove * -1
			damage := math.Max(1, math.Floor(float64(PlayerAttack)*0.5-float64(EnemyDefense)*0.25))
			defaultPositionX, defaultPositionY := EnemyDamage.getDefaultPosition()
			newMessage := message{
				value:      strconv.FormatFloat(damage, 'f', 0, 64),
				messageDiv: EnemyDamage,
				x:          defaultPositionX,
				y:          defaultPositionY,
				tick:       0,
			}
			messages = append(messages, newMessage)
			EnemyHp -= damage
		}
		if PlayerDiff <= 0 {
			PlayerMove = PlayerMove * -1
			PlayerDiff = 0
			PlayerTick = 0
			PlayerTurn = false
		}
	}

	if PlayerTick >= PlayerSpeed {
		PlayerTurn = true
	}
	if EnemyTick >= EnemySpeed && !PlayerTurn {
		EnemyTurn = true
	}

	PlayerActiveBar = float64(PlayerTick) / float64(PlayerSpeed)
	EnemyActiveBar = float64(EnemyTick) / float64(EnemySpeed)

	// log.Println(fmt.Sprintf("PlayerTick: %d,    EnemyTick: %d", PlayerTick, EnemyTick))
	// log.Println(fmt.Sprintf("PlayerActiveBar: %g", PlayerActiveBar))

	// メッセージの更新処理
	tempMessages := []message{}
	for _, message := range messages {
		message.update()
		if message.isExist() {
			tempMessages = append(tempMessages, message)
		}
		// fmt.Printf("messageDiv: %s, value: %s,  exist: %d \n", message.messageDiv.String(), message.value, message.messageDiv.getExistTick()-message.tick)
		// fmt.Printf("mx: %d,  my; %d \n", int(message.x), int(message.y))
	}
	messages = tempMessages

	return nil
}

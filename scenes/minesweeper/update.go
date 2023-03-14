package minesweeper

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/krile136/mineSweeper/scenes/minesweeper/explode"
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

		explodes = explode.Create()

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

	return nil
}

// 左クリックしたときの処理
func (m *MineSweeper) leftClick(x, y int) error {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {

		// クリックしたマスがフラグが立っていれば何もしない
		if m.field[y][x] == flag || m.field[y][x] == bomb {
			return nil
		}

		position := y*m.rows + x
		fmt.Printf("position: %d\n", position)
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
	}

	return nil
}

// 右クリックしたときの処理
func (m *MineSweeper) rightClick(x, y int) error {
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

		default:
			// 何もしない
		}
	}

	return nil
}

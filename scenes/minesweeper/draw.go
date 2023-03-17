package minesweeper

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/krile136/mineSweeper/internal/draw"
	"github.com/krile136/mineSweeper/internal/text"
	"github.com/krile136/mineSweeper/store"
)

func (m *MineSweeper) Draw(screen *ebiten.Image) {

	// 各ブロックを表示
	c := 1.0
	p := 32
	bf := store.Data.Layout.BattleField
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			x := (float64(j)+0.5)*float64(p)*c + scrollX
			y := (float64(i)+0.5)*float64(p)*c + scrollY + float64(bf)
			switch m.field[i][j] {
			case close:
				draw.Draw(screen, "minesweeper", c, c, x, y, 0, 0, 0, p, p)
			case zero:
				draw.Draw(screen, "minesweeper", c, c, x, y, 0, p, 0, p, p)
			case flag:
				draw.Draw(screen, "minesweeper", c, c, x, y, 0, p*2, 0, p, p)
			case bomb:
				draw.Draw(screen, "minesweeper", c, c, x, y, 0, p*3, 0, p, p)
			case one:
				draw.Draw(screen, "minesweeper", c, c, x, y, 0, 0, p, p, p)
			case two:
				draw.Draw(screen, "minesweeper", c, c, x, y, 0, p, p, p, p)
			case three:
				draw.Draw(screen, "minesweeper", c, c, x, y, 0, p*2, p, p, p)
			case four:
				draw.Draw(screen, "minesweeper", c, c, x, y, 0, p*3, p, p, p)
			case five:
				draw.Draw(screen, "minesweeper", c, c, x, y, 0, p*4, p, p, p)
			case six:
				draw.Draw(screen, "minesweeper", c, c, x, y, 0, p*5, p, p, p)
			case seven:
				draw.Draw(screen, "minesweeper", c, c, x, y, 0, p*6, p, p, p)
			case eight:
				draw.Draw(screen, "minesweeper", c, c, x, y, 0, p*7, p, p, p)
			}
		}
	}

	// スクロールバーを表示
	if isBarDisplay {
		draw.Draw(screen, "minesweeper", barLengthX/float64(p), 0.5, barSlideX, float64(store.Data.Layout.OutsideWidth), 0, p*5, 0, p, p)
		draw.Draw(screen, "minesweeper", 0.5, barLengthY/float64(p), float64(store.Data.Layout.OutsideWidth), barSlideY, 0, p*5, 0, p, p)
	}

	// BattleFiledを表示
	bf_x := store.Data.Layout.OutsideWidth
	sliceOffsetX := (bf_x - store.Data.Layout.OutsideWidth) / 2
	sliceOffsetY := 0
	draw.Draw(
		screen,
		"pipo-battlebg001",
		1,
		0.5,
		float64(store.Data.Layout.OutsideWidth)/2,
		float64(store.Data.Layout.BattleField)/2,
		0,
		sliceOffsetX,
		sliceOffsetY,
		store.Data.Layout.OutsideWidth,
		store.Data.Layout.BattleField*2)

	// キャラクターを描画
	if player.Dead() {
		if playerDraw.IsShowBlinking() {
			draw.DrawWithoutRect(screen, player.Name(), 1, playerDraw.CurrentPosition(), playerDraw.PositionY(), 0)
		}
	} else {
		draw.DrawWithoutRect(screen, player.Name(), 1, playerDraw.CurrentPosition(), playerDraw.PositionY(), 0)
	}

	if enemy.Dead() {
		if enemyDraw.IsShowBlinking() {
			draw.DrawWithoutRect(screen, enemy.Name(), 1, enemyDraw.CurrentPosition(), enemyDraw.PositionY(), 0)
		}
	} else {
		draw.DrawWithoutRect(screen, enemy.Name(), 1, enemyDraw.CurrentPosition(), enemyDraw.PositionY(), 0)
	}

	// アクティブバーを描画
	var center float64 = float64(store.Data.Layout.OutsideWidth) / 2
	var barFromCenter float64 = 50
	var barYAxis float64 = 92

	draw.Draw(screen, "minesweeper", 1.2, 0.2, center-barFromCenter, barYAxis, 0, p*5, 0, p, p)
	draw.Draw(screen, "minesweeper", 1.2*player.ActiveBar(), 0.2, center-barFromCenter-(1-player.ActiveBar())*float64(p)*0.6, barYAxis, 0, p*7, 0, p, p)
	draw.Draw(screen, "minesweeper", 1.2, 0.2, center+barFromCenter, barYAxis, 0, p*5, 0, p, p)
	draw.Draw(screen, "minesweeper", 1.2*enemy.ActiveBar(), 0.2, center+barFromCenter-(1-enemy.ActiveBar())*float64(p)*0.6, barYAxis, 0, p*7, 0, p, p)

	// 文字を描画
	var playerTextFromCenter int = 170
	var HpString string = fmt.Sprintf(" %d/%d", player.Hp(), player.MaxHp())
	text.DrawText(screen, fmt.Sprintf("Lv %d", player.Lv()), int(center)-60, 20, "S", store.Data.Color.Black)
	text.DrawText(screen, "HP", int(center)-playerTextFromCenter, 30, "M", store.Data.Color.Black)
	text.DrawText(screen, HpString, int(center)-playerTextFromCenter, 45, "M", store.Data.Color.Black)
	text.DrawText(screen, "NEXT", int(center)-playerTextFromCenter, 70, "M", store.Data.Color.Black)
	text.DrawText(screen, fmt.Sprintf(" %d", player.NextExp()), int(center)-playerTextFromCenter, 85, "M", store.Data.Color.Black)

	var enemyTextFromCenter int = 100
	text.DrawText(screen, fmt.Sprintf("Lv %d", enemy.Lv()), int(center)+40, 20, "S", store.Data.Color.Black)
	text.DrawText(screen, "HP", int(center)+enemyTextFromCenter, 30, "M", store.Data.Color.Black)
	text.DrawText(screen, fmt.Sprintf(" %d/%d", enemy.Hp(), enemy.MaxHp()), int(center)+enemyTextFromCenter, 45, "M", store.Data.Color.Black)

	// コンボ数を描画
	if currentCombo > 1 {
		var comboNumber string = fmt.Sprintf("%d", currentCombo)
		text.DrawText(screen, comboNumber, 10, 70, "M", getRainbow())
		text.DrawText(screen, "combo!!", text.Length(comboNumber, "M")+20, 70, "S", getRainbow())
	}

	// 爆発を描画
	explodes.Draw(screen)

	// メッセージを描画
	for _, v := range displayMessages {
		value, x, y, crl := v.GetFieldForDraw()
		text.DrawText(screen, value, x, y, "M", crl)
	}

	// マスをすべて開けた時のメッセージを表示する
	if(allOpenTick > 0){
		var center_y int = store.Data.Layout.OutsideHeight / 2
		text.DrawTextAtCenter(screen, "ALL OPEN BONUS", int(center), center_y-20, "L", getRainbow())
		text.DrawTextAtCenter(screen, "FULL RECOVERY", int(center), center_y+20, "L", getRainbow())
	}
	// クリアしたときのメッセージを表示する
	if isClear {
		var center_y int = store.Data.Layout.OutsideHeight / 2
		text.DrawTextAtCenter(screen, "C O N G R A T U L A T I O N !!", int(center), center_y, "L", getRainbow())
	}

	// FPSを表示する
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %d", int(ebiten.CurrentFPS())))
}

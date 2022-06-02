package store

import (
	"log"
	"os"
	"strconv"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// シーン間共通変数を定義する
// ここに入れる変数は、「シーンをまたいで変更されたり使用されたりする変数」のみとする

// シーン間共通変数を定義する
var Data Store

type Store struct {
	Layout      Layout
	Env         Env
	MineSweeper MineSweeper
}

// シーン間共通変数の初期化
func (s *Store) Init() error {
	Data.Layout.OutsideWidth = 320
	Data.Layout.OutsideHeight = 320
	Data.Layout.BattleField = 75

	tempScrollCorrectiveValue, err := strconv.Atoi(os.Getenv("SCROLL_CORRECTION_VALUE"))
	if err != nil {
		return err
	}
	Data.Env.ScrollCorrectionValue = tempScrollCorrectiveValue

	Data.MineSweeper.Rows = 20
	Data.MineSweeper.Columns = 20
	Data.MineSweeper.BombsNumber = 50

	Data.Env.Font, err = loadFont()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

type Layout struct {
	OutsideWidth  int
	OutsideHeight int
	BattleField   int
}

type Env struct {
	ScrollCorrectionValue int
	Font                  font.Face
}
type MineSweeper struct {
	Rows        int
	Columns     int
	BombsNumber int
}

func loadFont() (font.Face, error) {
	// fontデータを開くディレクトリパス
	// 現状、main.goからの相対パス
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	fontDir := currentDir + "/internal/assets/font/"
	const fontName = "PixelMplus10-Regular.ttf"

	ftBinary, err := os.ReadFile(fontDir + fontName)
	if err != nil {
		return nil, err
	}

	tt, err := opentype.Parse(ftBinary)
	if err != nil {
		return nil, err
	}
	const dpi = 72

	font, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, err
	}

	return font, nil

}

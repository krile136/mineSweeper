package store

import (
	"image/color"
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
	Font        Font
	Color       Color
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

	Data.Font.Large, Data.Font.Middle, Data.Font.Small, err = loadFont()
	if err != nil {
		log.Fatal(err)
	}

	Data.Color.Red = color.RGBA{255, 0, 0, 255}
	Data.Color.Blue = color.RGBA{0, 0, 255, 255}
	Data.Color.Green = color.RGBA{0, 255, 0, 255}
	Data.Color.Black = color.RGBA{0, 0, 0, 255}
	Data.Color.White = color.RGBA{255, 255, 255, 255}
	Data.Color.Gray = color.RGBA{128, 128, 128, 255}
	Data.Color.Orange = color.RGBA{255, 102, 0, 255}
	Data.Color.Yellow = color.RGBA{255, 255, 0, 255}
	Data.Color.Cyan = color.RGBA{0, 255, 255, 255}
	Data.Color.Magenta = color.RGBA{255, 0, 255, 255}

	return nil
}

type Layout struct {
	OutsideWidth  int
	OutsideHeight int
	BattleField   int
}

type Env struct {
	ScrollCorrectionValue int
}
type MineSweeper struct {
	Rows        int
	Columns     int
	BombsNumber int
}

type Font struct {
	Large  font.Face
	Middle font.Face
	Small  font.Face
}

type Color struct {
	Red     color.Color
	Blue    color.Color
	Green   color.Color
	Black   color.Color
	White   color.Color
	Gray    color.Color
	Orange  color.Color
	Yellow  color.Color
	Cyan    color.Color
	Magenta color.Color
}

func loadFont() (font.Face, font.Face, font.Face, error) {
	// fontデータを開くディレクトリパス
	// 現状、main.goからの相対パス
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, nil, nil, err
	}
	fontDir := currentDir + "/internal/assets/font/"
	const fontName = "PixelMplus10-Regular.ttf"

	ftBinary, err := os.ReadFile(fontDir + fontName)
	if err != nil {
		return nil, nil, nil, err
	}

	tt, err := opentype.Parse(ftBinary)
	if err != nil {
		return nil, nil, nil, err
	}
	const dpi = 72

	fontLarge, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	fontMiddle, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     54,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	fontSmall, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     36,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	return fontLarge, fontMiddle, fontSmall, nil

}

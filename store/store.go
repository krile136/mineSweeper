package store

import (
	"image/color"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	fonts "github.com/krile136/mineSweeper/internal/assets/fonts"
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
	CurrentScore int
}

// シーン間共通変数の初期化
func (s *Store) Init() error {
	Data.Layout.OutsideWidth = 640
	Data.Layout.OutsideHeight = 480
	Data.Layout.BattleField = 100

	Data.CurrentScore = 0

	tempScrollCorrectiveValue := 1
	Data.Env.ScrollCorrectionValue = tempScrollCorrectiveValue

	Data.MineSweeper.Rows = 20
	Data.MineSweeper.Columns = 20
	Data.MineSweeper.BombsNumber = 50

	Data.Font.Large, Data.Font.Middle, Data.Font.Small, _ = loadFont()

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
	Data.Color.Purple = color.RGBA{128, 0, 128, 255}

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
	Purple  color.Color
}

func loadFont() (font.Face, font.Face, font.Face, error) {
	// バイナリデータを直接渡す
	fontData := fonts.CompressedTTF
	tt, err := opentype.Parse(fontData)
	if err != nil {
		log.Printf("failed to load font Data: %s", err.Error())
		return nil, nil, nil, err
	}
	fontLarge, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Printf("failed to open large font face : %s", err.Error())
		return nil, nil, nil, err
	}

	fontMiddle, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     54,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Printf("failed to open middle font face : %s", err.Error())
		return nil, nil, nil, err
	}

	fontSmall, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     36,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Printf("failed to open small font face : %s", err.Error())
		return nil, nil, nil, err
	}

	return fontLarge, fontMiddle, fontSmall, nil

}

package store

import (
	"os"
	"strconv"
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

	tempScrollCorrectiveValue, err := strconv.Atoi(os.Getenv("SCROLL_CORRECTION_VALUE"))
	if err != nil {
		return err
	}
	Data.Env.ScrollCorrectionValue = tempScrollCorrectiveValue

	Data.MineSweeper.Rows = 20
	Data.MineSweeper.Columns = 20
	Data.MineSweeper.BombsNumber = 50

	return nil
}

type Layout struct {
	OutsideWidth  int
	OutsideHeight int
}

type Env struct {
	ScrollCorrectionValue int
}
type MineSweeper struct {
	Rows        int
	Columns     int
	BombsNumber int
}

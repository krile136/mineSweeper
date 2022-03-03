package store

// シーン間共通変数を定義する
// ここに入れる変数は、「シーンをまたいで変更されたり使用されたりする変数」のみとする

// シーン間共通変数を定義する
var Data Store

type Store struct {
	MineField MineField
}

// シーン間共通変数の初期化
func (s *Store) Init() error {
	Data.MineField.Rows = 25
	Data.MineField.Columns = 25

	return nil
}

type MineField struct {
	Rows    int
	Columns int
}

package minesweeper

import (
	"log"
	"math"
	"math/rand"
)

func (m *MineSweeper) placeBombs() error {
	// フィールドを初期化
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			m.field[i][j] = close
		}
	}

	// 重複のないランダムな数字をrows * columns の中からBombsNumberの分だけ作り
	// 配列に保存する
	m.bombsPosition = make([]int, m.bombsNumber)
	count := 0
	for count < m.bombsNumber {
		position := rand.Intn(m.rows * m.columns)
		if !inArray(m.bombsPosition, position) {
			m.bombsPosition[count] = position
			count++
		}
	}

	return nil
}

func (m *MineSweeper) placeFlag(x, y int) error {

	m.field[y][x] = flag

	return nil
}

// 周りの爆弾の数をカウントし、その数に応じて引数に渡されたマスのフィールドを決定する
func (m *MineSweeper) searchAround(x, y int) {
	var bombs int = 0
	var next []int
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if inBetween(0, i, m.rows-1) && inBetween(0, j, m.columns-1) {
				position := i*m.rows + j
				if inArray(m.bombsPosition, position) {
					bombs += 1
				} else {
					if m.field[i][j] == close {
						next = append(next, position)
					}
				}
			}
		}
	}
	if bombs == 0 {
		nextCheck = append(nextCheck, next...)
	}
	// フラグがおいてあるマスはフィールド情報を更新しない
	if m.field[y][x] != flag && m.field[y][x] == close {
		m.field[y][x] = nums[bombs]
		GetExp += 1
	}
	if len(nextCheck) > 0 {
		nextCheck = nextCheck[1:]
	}
}

// 周りの爆弾の数が表示されているフィールドにて、周囲のフィールドを走査する
func (m *MineSweeper) searchAroundOnNumberField(x, y int) {
	var next []int
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if inBetween(0, i, m.rows-1) && inBetween(0, j, m.columns-1) {
				position := i*m.rows + j
				if inArray(m.bombsPosition, position) {
					// 周りのマスを開いたときに爆弾があった場合
					if m.field[i][j] != flag {
						// フラグおいてないのでゲームーオーバー
						m.field[i][j] = bomb
						log.Print("game over! (right click)")
					}
				} else {
					if m.field[i][j] == close {
						next = append(next, position)
					}
				}
			}
		}
	}
	nextCheck = append(nextCheck, next...)
}

// int型の配列の中に特定のint型の値が含まれるかチェックする
func inArray(array []int, needle int) bool {
	for _, val := range array {
		if needle == val {
			return true
		}
	}
	return false
}

// int型の値が最小と最大の間にあるかチェックする
func inBetween(min, val, max int) bool {
	if (val >= min) && (val <= max) {
		return true
	} else {
		return false
	}
}

// float64型の値が最大値と最小値の間に収まるようにする
func setBetween(min, val, max float64) float64 {
	return math.Min(max, math.Max(min, val))
}

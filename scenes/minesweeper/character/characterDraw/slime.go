package characterDraw

import "github.com/krile136/mineSweeper/store"

type Slime struct {
	*abstractCharacterDraw
}

// コンストラクタ
func (s Slime) New() (new CharacterDrawInterface) {
	acd := s.makeAbstractCharacterDraw(s.defaultField())
	new = Slime{acd}
	return
}

// 表示位置を変化させる
func (s Slime) ExecuteMoving() CharacterDrawInterface {
	new := s
	new.addDirectionToDifference()
	return new
}

func (s Slime) CanExecuteInvertAtBase() bool {
	return s.difference >= 0
}

func (s Slime) CanFinishAppearing() bool {
	return s.difference <= 0
}

func (s Slime) IsReturningToBase() bool {
	return s.direction > 0
}

func (s Slime) UpdateBlinking() CharacterDrawInterface {
	new := s
	new.blinkingTick += 1

	return new
}

// 移動方向を反転させる
func (s Slime) InvertDirection() CharacterDrawInterface {
	new := s
	s.invertDirection()
	return new
}

// ターンを終了させる
func (s Slime) FinishTurn() CharacterDrawInterface {
	new := s
	new.positionX, new.positionY, new.direction, new.difference = s.defaultField()

	return new
}

func (s Slime) SetInitialDraw() CharacterDrawInterface {
	new := s
	new.difference = float64(store.Data.Layout.OutsideWidth) / 2
	new.direction = -1
	return new
}

// デフォルトのフィールド値を取得する
func (s Slime) defaultField() (positionX, positionY, direction, difference float64) {
	var center float64 = float64(store.Data.Layout.OutsideWidth) / 2
	const diffFromCenterToCharacter = 30
	positionX = center + s.getDiffFromCenterToCharacter()
	positionY = s.getCharacterYPosition()
	direction = -1
	difference = 0
	return
}

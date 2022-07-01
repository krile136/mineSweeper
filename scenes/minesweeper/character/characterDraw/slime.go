package characterDraw

type Slime struct {
	*abstractCharacterDraw
}

// コンストラクタ
func (s Slime) New() (new CharacterDrawInterface) {
	acd := s.makeAbstractCharacterDraw(s.defaultField())
	new = Player{acd}
	return
}

// 表示位置を変化させる
func (s Slime) ExecuteMoving() CharacterDrawInterface {
	new := s
	new.addDirectionToDifference()
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
	new.direction, new.difference = s.defaultField()

	return new
}

// デフォルトのフィールド値を取得する
func (s Slime) defaultField() (direction, difference float64) {
	direction = -1
	difference = 0
	return
}

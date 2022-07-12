package characterDraw

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
	new.difference = 150
	new.direction = -1
	return new
}

// デフォルトのフィールド値を取得する
func (s Slime) defaultField() (positionX, positionY, direction, difference float64) {
	positionX = 195
	positionY = 40
	direction = -1
	difference = 0
	return
}

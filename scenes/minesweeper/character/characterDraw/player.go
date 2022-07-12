package characterDraw

type Player struct {
	*abstractCharacterDraw
}

// コンストラクタ
func (p Player) New() (new CharacterDrawInterface) {
	acd := p.makeAbstractCharacterDraw(p.defaultField())
	new = Player{acd}
	return
}

// 表示位置を変化させる
func (p Player) ExecuteMoving() CharacterDrawInterface {
	new := p
	new.addDirectionToDifference()
	return new
}

func (p Player) CanExecuteInvertAtBase() bool {
	return p.difference <= 0
}

func (p Player) UpdateBlinking() CharacterDrawInterface {
	new := p
	new.blinkingTick += 1

	return new
}

// 移動方向を反転させる
func (p Player) InvertDirection() CharacterDrawInterface {
	new := p
	p.invertDirection()
	return new
}

// ターンを終了させる
func (p Player) FinishTurn() CharacterDrawInterface {
	new := p
	new.positionX, new.positionY, new.direction, new.difference = p.defaultField()

	return new
}

func (p Player) SetInitialDraw() CharacterDrawInterface {
	new := p
	new.difference = -150
	new.direction = 1
	return new
}

// デフォルトのフィールド値を取得する
func (p Player) defaultField() (positionX, positionY, direction, difference float64) {
	positionX = 110
	positionY = 40
	direction = 1
	difference = 0
	return
}

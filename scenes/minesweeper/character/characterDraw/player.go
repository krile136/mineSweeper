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

// 移動方向を反転させる
func (p Player) InvertDirection() CharacterDrawInterface {
	new := p
	p.invertDirection()
	return new
}

// ターンを終了させる
func (p Player) FinishTurn() CharacterDrawInterface {
	new := p
	new.direction, new.difference = p.defaultField()

	return new
}

// デフォルトのフィールド値を取得する
func (p Player) defaultField() (direction, difference float64) {
	direction = 1
	difference = 0
	return
}

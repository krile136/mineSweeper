package view

type Orange struct {
	*abstractExplodeView
}

func (o Orange) New(x, y float64, tick int, delay int) (new ExplodeViewInterface) {
	ae := o.makeAbstractExplodeView(x, y, 32*5, tick, delay)
	new = Orange{ae}
	return
}

func (o Orange) Update() ExplodeViewInterface {
	new := o
	if new.delay <= 0 {
		new.tick += 1
	} else {
		new.delay -= 1
	}
	return new
}

package view

type Orange struct {
	*abstractExplodeView
}

func (o Orange) New(x, y float64, tick int) (new ExplodeViewInterface) {
	ae := o.makeAbstractExplodeView(x, y, 32*5, tick)
	new = Orange{ae}
	return
}

func (o Orange) Update() ExplodeViewInterface {
	new := o
	new.tick += 1
	return new
}

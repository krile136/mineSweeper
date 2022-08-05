package view

type Orange struct {
	*abstractExplodeView
}

func (o Orange) New(x, y float64) (new ExplodeViewInterface) {
	ae := o.makeAbstractExplodeView(x, y, 32*5)
	new = Orange{ae}
	return
}

func (o Orange) Update() ExplodeViewInterface {
	new := o
	new.tick += 1
	return new
}

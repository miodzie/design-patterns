package visitor

type Picture struct {
	graphics []Graphic
}

func (p *Picture) Accept(visitor GraphicVisitor) {
	for _, g := range p.graphics {
		g.Accept(visitor)
	}
}

func (p *Picture) Draw() {
	for _, g := range p.graphics {
		g.Draw()
	}
}

func (p *Picture) Add(g Graphic) {
	p.graphics = append(p.graphics, g)
}

func (p *Picture) Remove(graphic Graphic) {
	for i, g := range p.graphics {
		if g == graphic {
			p.graphics = append(p.graphics[:i], p.graphics[i+1:]...)
			return
		}
	}
}

func (p *Picture) GetChild(i int) {
	//TODO implement me
	panic("implement me")
}

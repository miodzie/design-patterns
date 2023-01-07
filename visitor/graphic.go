package visitor

type Graphic interface {
	Accept(GraphicVisitor)
	Draw()
	Add(Graphic)
	Remove(Graphic)
	GetChild(int)
}

type Primitive struct{}

func (p Primitive) Accept(GraphicVisitor) {}
func (p Primitive) Add(Graphic)           {}
func (p Primitive) Remove(Graphic)        {}
func (p Primitive) GetChild(int)          {}

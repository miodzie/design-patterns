package composite

type Graphic interface {
	Draw()
	Add(Graphic)
	Remove(Graphic)
	GetChild(int)
}

type Primitive struct{}

func (p Primitive) Add(Graphic)    {}
func (p Primitive) Remove(Graphic) {}
func (p Primitive) GetChild(int)   {}

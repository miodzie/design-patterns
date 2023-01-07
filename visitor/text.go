package visitor

import "fmt"

type Text struct {
	Text string
	Primitive
}

func (t *Text) Accept(visitor GraphicVisitor) {
	visitor.VisitText(t)
}

func (t *Text) Draw() {
	fmt.Print(t.Text)
}

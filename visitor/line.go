package visitor

import "fmt"

type Line struct {
	Text string
	Primitive
}

func NewLine() *Line {
	return &Line{Text: "__________________\n"}
}

func (l *Line) Draw() {
	fmt.Print(l.Text)
}

func (l *Line) Accept(visitor GraphicVisitor) {
	visitor.VisitLine(l)
}

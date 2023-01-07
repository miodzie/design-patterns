package composite

import "fmt"

type Text struct {
	Text string
	Primitive
}

func (t *Text) Draw() {
	fmt.Print(t.Text)
}

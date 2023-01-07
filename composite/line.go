package composite

import "fmt"

type Line struct {
	Primitive
}

func (l Line) Draw() {
	fmt.Println("__________________")
}

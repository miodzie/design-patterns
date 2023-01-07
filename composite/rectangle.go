package composite

import "fmt"

type Rectangle struct {
	Primitive
}

func (d *Rectangle) Draw() {
	fmt.Print(`
	_____________________
        |                    |
	|                    |
	|                    |
	|____________________|
`)
}

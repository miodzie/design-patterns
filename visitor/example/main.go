package main

import (
	. "patterns/visitor"
)

func main() {
	picture := &Picture{}
	picture.Add(&Line{})
	picture.Add(&Text{Text: "howdy doody\n"})
	picture.Add(&Line{})
	picture.Draw()

	picture.Accept(&SpookySkeletonVisitor{})
	picture.Draw()
}
